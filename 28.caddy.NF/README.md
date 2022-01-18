# Caddy

```CMD
Caddy 2 is a powerful, enterprise-ready, open source web server with automatic HTTPS written in Go.

docker run -d -p 80:80 -p 443:443 \
    -v /site:/srv \
    -v caddy_data:/data \
    -v caddy_config:/config \
    caddy caddy file-server --domain dockerttt.ddns.net
```