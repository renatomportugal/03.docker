# Caddy

```CMD
docker run -d -p 80:80 -p 443:443 \
    -v /site:/srv \
    -v caddy_data:/data \
    -v caddy_config:/config \
    caddy caddy file-server --domain tecnocrata.ddns.net
```