package ad

import (
	"errors"
	"strings"
)

type SecurityType int

//Security will default to SecurityNone if not given.
const (
	SecurityNone SecurityType = iota
	SecurityTLS
	SecurityStartTLS
	SecurityInsecureTLS
	SecurityInsecureStartTLS
)

type Config struct {
	Server   string
	Port     string
	BaseDN   string
	Security SecurityType

	User     string
	Password string
}

func (c *Config) Domain() (string, error) {
	domain := ""
	for _, v := range strings.Split(strings.ToLower(c.BaseDN), ",") {
		if trimmed := strings.TrimSpace(v); strings.HasPrefix(trimmed, "dc=") {
			domain = domain + "." + trimmed[3:]
		}
	}
	if len(domain) <= 1 {
		return "", errors.New("configuration error: invalid BaseDN")
	}
	return domain[1:], nil
}
