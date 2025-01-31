package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(DomainFileValidator{})
}

// DomainFileValidator is a Caddy module that validates domain files.
type DomainFileValidator struct {
	DomainsDir string `json:"domains_dir,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (DomainFileValidator) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.domain_file_validator",
		New: func() caddy.Module { return new(DomainFileValidator) },
	}
}

// Provision sets up the module.
func (dfv *DomainFileValidator) Provision(ctx caddy.Context) error {
	if dfv.DomainsDir == "" {
		dfv.DomainsDir = "/etc/openpanel/caddy/domains" // Default directory
	}

	// Scan the directory for files
	files, err := filepath.Glob(filepath.Join(dfv.DomainsDir, "*"))
	if err != nil {
		return fmt.Errorf("failed to scan domains directory: %v", err)
	}

	// Validate each file
	for _, file := range files {
		_, err := caddyfile.Load(file, nil)
		if err != nil {
			fmt.Printf("Invalid syntax in file: %s. Skipping...\n", file)
			continue
		}
		fmt.Printf("File is valid: %s\n", file)
	}

	return nil
}

// Validate ensures the module's configuration is valid.
func (dfv *DomainFileValidator) Validate() error {
	return nil
}

// ServeHTTP implements the http.Handler interface.
func (dfv DomainFileValidator) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile sets up the module from a Caddyfile.
func (dfv *DomainFileValidator) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&dfv.DomainsDir) {
			return d.ArgErr()
		}
	}
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*DomainFileValidator)(nil)
	_ caddy.Validator             = (*DomainFileValidator)(nil)
	_ caddyhttp.MiddlewareHandler = (*DomainFileValidator)(nil)
	_ caddyfile.Unmarshaler       = (*DomainFileValidator)(nil)
)
