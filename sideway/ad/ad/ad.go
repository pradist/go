package ad

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"time"
)

type Client struct {
	name   string
	config *Config
}

func NewClient(name string, cfg *Config) func() (*Client, func(), error) {
	return func() (*Client, func(), error) {
		client := &Client{
			name:   name,
			config: cfg,
		}

		cleanup := func() {
			if err := client.Close(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Ad Closeed [%s]", client.name)
			}
		}

		return client, cleanup, nil
	}
}

func (a *Client) connect() (*ldap.Conn, error) {
	ldap.DefaultTimeout = 10 * time.Second
	switch a.config.Security {
	case SecurityNone:
		conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", a.config.Server, a.config.Port))
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		return conn, nil
	case SecurityTLS:
		conn, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%s", a.config.Server, a.config.Port),
			&tls.Config{
				ServerName:         a.config.Server,
				InsecureSkipVerify: true,
			})
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		return conn, nil
	case SecurityStartTLS:
		conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", a.config.Server, a.config.Port))
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		err = conn.StartTLS(&tls.Config{ServerName: a.config.Server})
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		return conn, nil
	case SecurityInsecureTLS:
		conn, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%s", a.config.Server, a.config.Port), &tls.Config{ServerName: a.config.Server, InsecureSkipVerify: true})
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		return conn, nil
	case SecurityInsecureStartTLS:
		conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", a.config.Server, a.config.Port))
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		err = conn.StartTLS(&tls.Config{ServerName: a.config.Server, InsecureSkipVerify: true})
		if err != nil {
			return nil, fmt.Errorf("connection error: %v", err)
		}
		return conn, nil
	default:
		return nil, errors.New("configuration error: invalid SecurityType")
	}
}

func (a *Client) Close() error {
	return nil
}

func (a *Client) AuthenticateWithOutDomain(username, password string) (bool, int, error) {
	domain, _ := a.config.Domain()
	usernameWithDomain := username + "@" + domain
	return a.Authenticate(usernameWithDomain, password)
}

func (a *Client) Authenticate(username, password string) (bool, int, error) {
	if password == "" {
		return false, 1, nil
	}
	lcon, err := a.connect()
	if err != nil {
		return false, 2, err
	}
	defer lcon.Close()
	err = lcon.Bind(username, password)

	if err != nil {
		if e, ok := err.(*ldap.Error); ok {
			if e.ResultCode == ldap.LDAPResultInvalidCredentials {
				return false, 3, fmt.Errorf("user or password invalid (%s): %v", username, err)
			}
		}
		return false, 4, fmt.Errorf("bind error (%s): %v", username, err)
	}
	return true, 0, nil
}

func (a *Client) FindInfo(filtername string, filter string) (map[string]string, error) {
	if a.config.User == "" || a.config.Password == "" || a.config.BaseDN == "" {
		return nil, fmt.Errorf("user or password or baseDN is empty")
	}

	lcon, err := a.connect()
	if err != nil {
		return nil, err
	}

	defer lcon.Close()
	err = lcon.Bind(a.config.User, a.config.Password)
	if err != nil {
		if e, ok := err.(*ldap.Error); ok {
			if e.ResultCode == ldap.LDAPResultInvalidCredentials {
				return nil, fmt.Errorf("user or password invalid (%s): %v", a.config.User, err)
			}
		}
		return nil, fmt.Errorf("bind error (%s): %v", a.config.User, err)
	}

	searchRequest := ldap.NewSearchRequest(
		a.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(%s=%s))", filtername, filter),
		[]string{},
		nil,
	)
	sr, err := lcon.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	entry := sr.Entries[0]
	data := make(map[string]string)

	for _, attr := range entry.Attributes {
		data[attr.Name] = attr.Values[0]
	}

	return data, nil

}
