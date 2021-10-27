# Traefic

```CMD
Start Traefik:

docker run -d -p 8080:8080 -p 80:80 \
-v $PWD/traefik.yml:/etc/traefik/traefik.yml \
-v /var/run/docker.sock:/var/run/docker.sock \
traefik:v2.5


Start a backend server, named test:

docker run -d --name test traefik/whoami

And finally, you can access to your whoami server throught Traefik, on the domain name test.docker.localhost:

# $ curl --header 'Host:test.docker.localhost' 'http://localhost:80/'

curl --header 'Host:test.docker.localhost' 'http://192.168.1.106:80/'


$ curl test.docker.localhost
Hostname: 390a880bdfab
IP: 127.0.0.1
IP: 172.17.0.3
GET / HTTP/1.1
Host: test.docker.localhost
User-Agent: curl/7.65.3
Accept: */*
Accept-Encoding: gzip
X-Forwarded-For: 172.17.0.1
X-Forwarded-Host: test.docker.localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 7e073cb54211
X-Real-Ip: 172.17.0.1

The web UI http://localhost:8080 will give you an overview of the routers, services, and middlewares.


Grab a sample configuration file and rename it to traefik.toml. Enable docker provider and web UI:

## traefik.toml

# API and dashboard configuration
[api]

# Docker configuration backend
[docker]
  domain = "docker.localhost"
Start Traefik:

docker run -d -p 8080:8080 -p 80:80 \
-v $PWD/traefik.toml:/etc/traefik/traefik.toml \
-v /var/run/docker.sock:/var/run/docker.sock \
traefik:v1.7
Start a backend server, named test:

docker run -d --name test traefik/whoami
And finally, you can access to your whoami server throught Traefik, on the domain name {containerName}.{configuredDomain} (test.docker.localhost):

# $ curl --header 'Host:test.docker.localhost' 'http://localhost:80/'
$ curl 'http://test.docker.localhost'
Hostname: 117c5530934d
IP: 127.0.0.1
IP: ::1
IP: 172.17.0.3
IP: fe80::42:acff:fe11:3
GET / HTTP/1.1
Host: test.docker.localhost
User-Agent: curl/7.35.0
Accept: */*
Accept-Encoding: gzip
X-Forwarded-For: 172.17.0.1
X-Forwarded-Host: 172.17.0.3:80
X-Forwarded-Proto: http
X-Forwarded-Server: f2e05c433120
The web UI http://localhost:8080 will give you an overview of the frontends/backends and also a health dashboard.
```