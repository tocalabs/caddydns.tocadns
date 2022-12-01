package tocadns

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	godaddy "github.com/tocalabs/tlsdns.godaddy"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *godaddy.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.tocadns",
		New: func() caddy.Module { return &Provider{new(godaddy.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIToken = caddy.NewReplacer().ReplaceAll(p.Provider.APIToken, "")
	p.Provider.APIHost = caddy.NewReplacer().ReplaceAll(p.Provider.APIHost, "")
	return nil
}

// TODO: This is just an example. Update accordingly.
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// tocadns {
//     api_token <api_token>
// .   api_host <api_host>
// }
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API Token already exists")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIToken = d.Val()
			case "api_host":
				if p.Provider.APIHost != "" {
					return d.Err("API host already exists")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIHost = d.Val()
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	if p.Provider.APIHost == "" {
		return d.Err("missing API host")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
