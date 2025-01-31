# Caddyfile-validator
On Caddy start validate imported files, skip importing files that have invalid configuration


why?

because i does not male sense to fail to start a webserver just because one domain file is invalid 💁

## Usage

1. Build caddy with the plugin:
   ```
   go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

   xcaddy build --with github.com/stefanpejcic/Caddyfile-validator

   ```

2. add inCaddyfile before the import:
   ```
   # Domain file validator
   domain_file_validator /etc/openpanel/caddy/domains
   ```

3. start caddy
   ```
   ./caddy run
   ```
that's it.
