FROM --platform=linux/amd64 portainer/portainer-ce:latest as portainer
FROM i386/alpine as alpine

COPY --from=portainer / /

VOLUME /data
WORKDIR /

EXPOSE 9000
EXPOSE 8000

# amd64 emulation
RUN apk add qemu-x86_64

# 32bit docker bin
RUN apk add docker && rm /docker && mv /usr/bin/docker /docker && apk del docker

#RUN apk add file util-linux

#ENTRYPOINT ["sh"]
CMD qemu-x86_64 /portainer