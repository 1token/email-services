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
		{c.Web.Protocol != "http" && c.Web.Protocol != "https", "must supply 'http' or 'https' Protocol"},
		{c.Web.Host == "", "must supply a Host to listen on"},
		{c.Web.Protocol == "https" && (c.Web.TLSCert == "" || c.Web.TLSKey == ""), "must specific both a TLS cert and key"},
		{len(c.Web.AllowedOrigins) == 0, "must specify at least one Allowed Origin"},
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
	Protocol       string   `json:"protocol"`
	Host           string   `json:"host"`
	Port           string   `json:"port"`
	TLSCert        string   `json:"tlsCert"`
	TLSKey         string   `json:"tlsKey"`
	AllowedOrigins []string `json:"allowedOrigins"`
}

func (w Web) Addr() string {
	addr := w.Host
	if w.Port != "" {
		addr += ":" + w.Port
	}
	return addr
}
