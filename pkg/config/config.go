package config

import (
	"fmt"
	"strings"
)

// Config is the config format for the main application.
type Config struct {
	Web Web `json:"web"`
}

//Validate the configuration
func (c Config) Validate() error {
	// Fast checks. Perform these first for a more responsive CLI.
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{c.Web.HTTP == "" && c.Web.HTTPS == "", "must supply a HTTP/HTTPS  address to listen on"},
		{c.Web.HTTPS != "" && c.Web.TLSCert == "", "no cert specified for HTTPS"},
		{c.Web.HTTPS != "" && c.Web.TLSKey == "", "no private key specified for HTTPS"},
	}

	var checkErrors []string

	for _, check := range checks {
		if check.bad {
			checkErrors = append(checkErrors, check.errMsg)
		}
	}
	if len(checkErrors) != 0 {
		return fmt.Errorf("invalid Config:\n\t-\t%s", strings.Join(checkErrors, "\n\t-\t"))
	}
	return nil
}

// Web is the config format for the HTTP server.
type Web struct {
	HTTP           string   `json:"http"`
	HTTPS          string   `json:"https"`
	TLSCert        string   `json:"tlsCert"`
	TLSKey         string   `json:"tlsKey"`
	AllowedOrigins []string `json:"allowedOrigins"`
}
