# Tocadns module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records for toca.

## Caddy module name

```
dns.providers.tocadns
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "tocadns",
				"api_token": "YOUR_TOCA_API_TOKEN" //
				"api_host": "API_HOST_HERE"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns godaddy {env.GODADDY_TOKEN}
}
```

You can replace `{env.GODADDY_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/godaddy) for important information about credentials.