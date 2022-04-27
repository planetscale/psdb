package snappy

import (
	"io"
	"sync"

	"github.com/golang/snappy"
	"google.golang.org/grpc/encoding"
)

const Name = "snappy"

func init() {
	c := &compressor{}
	c.poolCompressor.New = func() any {
		return &writer{Writer: snappy.NewBufferedWriter(io.Discard), pool: &c.poolCompressor}
	}
	c.poolDecompressor.New = func() any {
		return &reader{Reader: snappy.NewReader(&discardReader{}), pool: &c.poolDecompressor}
	}
	encoding.RegisterCompressor(c)
}

type discardReader struct{}

func (discardReader) Read(p []byte) (int, error) { return 0, io.EOF }

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}

type writer struct {
	*snappy.Writer
	pool *sync.Pool
}

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	z := c.poolCompressor.Get().(*writer)
	z.Reset(w)
	return z, nil
}

func (z *writer) Close() error {
	defer z.pool.Put(z)
	return z.Writer.Close()
}

type reader struct {
	*snappy.Reader
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	z := c.poolDecompressor.Get().(*reader)
	z.Reset(r)
	return z, nil
}

func (z *reader) Read(p []byte) (n int, err error) {
	n, err = z.Reader.Read(p)
	if err == io.EOF {
		z.pool.Put(z)
	}
	return n, err
}

func (c *compressor) DecompressedSize(buf []byte) int {
	l, err := snappy.DecodedLen(buf)
	if err != nil {
		return -1
	}
	return l
}

func (c *compressor) Name() string {
	return Name
}

/*
TODO: figure out how/if we can use a snappy no-op compressor
      for the transparent proxy path. This proved to be difficult.

type noopCompressor struct{}
type noopWriteCloser struct {
	io.Writer
}

func (w noopWriteCloser) Close() error { return nil }

func (c *noopCompressor) Name() string {
	return Name
}

func (c *noopCompressor) Compress(w io.Writer) (io.WriteCloser, error) {
	return noopWriteCloser{w}, nil
}

func (c *noopCompressor) Decompress(r io.Reader) (io.Reader, error) {
	return r, nil
}

func Noop() *noopCompressor {
	return &noopCompressor{}
}

func MakeNoop() {
	encoding.RegisterCompressor(Noop())
}
*/
