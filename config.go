package psdb

import "net/url"

// Config is a PlanetScale connection configuration
type Config struct {
	Host     string
	User     string
	Password string
}

func ConfigFromURL(rawURL string) Config {
	var cfg Config
	if u, err := url.Parse(rawURL); err == nil {
		cfg.Host = u.Host
		cfg.User = u.User.Username()
		cfg.Password, _ = u.User.Password()
	}
	return cfg
}
