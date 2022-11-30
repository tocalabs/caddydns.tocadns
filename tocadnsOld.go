// package tocadnsold

// import (
// 	"fmt"

// 	"github.com/caddyserver/caddy/v2"
// 	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
// 	godaddy "github.com/tocalabs/tlsdns.godaddy"
// )

// // Provider wraps the provider implementation as a Caddy module.
// type Provider struct{ *godaddy.Provider }

// func init() {
// 	caddy.RegisterModule(Provider{})
// }

// // CaddyModule returns the Caddy module information.
// func (Provider) CaddyModule() caddy.ModuleInfo {
// 	return caddy.ModuleInfo{
// 		ID:  "dns.providers.tocadns",
// 		New: func() caddy.Module { return &Provider{new(godaddy.Provider)} },
// 	}
// }

// // Before using the provider config, resolve placeholders in the API token.
// // Implements caddy.Provisioner.
// func (p *Provider) Provision(ctx caddy.Context) error {
// 	repl := caddy.NewReplacer()
// 	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
// 	p.Provider.APIHost = repl.ReplaceAll(p.Provider.APIHost, "")
// 	return nil
// }

// // UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
// //
// //	tocadns {
// //	    api_token <api_token>
// //		api_host <api_host>
// //	}
// func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
// 	for d.Next() {
		
// 		for nesting := d.Nesting(); d.NextBlock(nesting); {
// 			fmt.Printf("Nesting block")
// 			switch d.Val() {
// 				case "api_token":
// 					if p.Provider.APIToken != "" {
// 						return d.Err("API Token already set")
// 					}
// 					if d.NextArg() {
// 						fmt.Printf("API Token %s", d.Val())
// 						p.Provider.APIToken = d.Val()
// 					}
// 					if d.NextArg() {
// 						fmt.Printf("Got to next arg - api_token")
// 						return d.ArgErr()
// 					}
// 				case "api_host":
// 					if p.Provider.APIHost != "" {
// 						return d.Err("API Host already set")
// 					}
// 					if d.NextArg() {
// 						p.Provider.APIHost = d.Val()
// 					}
// 					if d.NextArg() {
// 						return d.ArgErr()
// 					}
// 				default:
// 					return d.Errf("unrecognized subdirective '%s'", d.Val())
// 			}
// 		}
// 	}
// 	if p.Provider.APIToken == "" {
// 		return d.Err("missing API token")
// 	}
// 	if p.Provider.APIHost == "" {
// 		return d.Err("missing API host")
// 	}
// 	return nil
// }

// // Interface guards
// var (
// 	_ caddyfile.Unmarshaler = (*Provider)(nil)
// 	_ caddy.Provisioner     = (*Provider)(nil)
// )
