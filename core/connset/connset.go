package connset

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"
)

// ErrConnSetClosed is returned for any operation when the connset was
// closed with the Close() method.
var ErrConnSetClosed = errors.New("connset is closed")

type ConnSet struct {
	dialer DialerFn

	mu    sync.Mutex
	conns map[string]*trackedConn

	closed  bool
	closeCh chan struct{}
}

type trackedConn struct {
	conn io.Closer
	c    int
	lu   time.Time
}

type DialerFn func(ctx context.Context, addr string) (io.Closer, error)

func New(dialer DialerFn, ttl time.Duration) *ConnSet {
	c := &ConnSet{
		conns:   make(map[string]*trackedConn),
		dialer:  dialer,
		closeCh: make(chan struct{}, 1),
	}
	if int64(ttl) >= 0 {
		go c.cleanup(ttl)
	}
	return c
}

// Conn returns the connection for the given addr.
func (c *ConnSet) Get(ctx context.Context, addr string) (io.Closer, error) {
	if c.closed {
		return nil, ErrConnSetClosed
	}

	c.mu.Lock()
	tconn, ok := c.conns[addr]
	if ok {
		// we already dialed, return the existing conn
		tconn.c++
		c.mu.Unlock()
		return tconn.conn, nil
	}
	c.mu.Unlock()

	conn, err := c.dialer(ctx, addr)
	if err != nil {
		return nil, err
	}

	// it's possible that two connections dialed at the same time, make sure to
	// only store one connection
	c.mu.Lock()
	defer c.mu.Unlock()

	tconn, ok = c.conns[addr]
	if ok {
		// we already dialed (the first dial won the race), return the existing
		// conn and close the second dialed connection
		go conn.Close()
		tconn.c++
		return tconn.conn, nil
	}

	// the first dial won the race, store the connection
	tconn = &trackedConn{conn, 1, time.Time{}}
	c.conns[addr] = tconn

	return tconn.conn, nil
}

// Releases the connection, effectively decrementing it's lease count
func (c *ConnSet) Release(addr string) error {
	if c.closed {
		return ErrConnSetClosed
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	tconn, ok := c.conns[addr]
	if !ok {
		return nil
	}
	tconn.c--
	tconn.lu = time.Now()

	return nil
}

// Close closes all underlying connections
func (c *ConnSet) Close() {
	c.closed = true

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, tconn := range c.conns {
		go tconn.conn.Close()
	}

	c.conns = make(map[string]*trackedConn)
	c.closeCh <- struct{}{}
}

func (c *ConnSet) cleanup(ttl time.Duration) {
	freq := time.Duration(ttl.Milliseconds()/3) * time.Millisecond

	for {
		select {
		case <-c.closeCh:
			return
		case <-time.After(freq):
			if c.closed {
				return
			}

			now := time.Now()
			threshold := now.Add(-ttl)

			c.mu.Lock()
			for addr, tconn := range c.conns {
				if tconn.c <= 0 && tconn.lu.Before(threshold) {
					delete(c.conns, addr)
					go tconn.conn.Close()
				}
			}
			c.mu.Unlock()
		}
	}
}
