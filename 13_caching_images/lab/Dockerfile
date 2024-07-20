FROM alpine:latest
ENV DAEMON_VERSION 1.0.0

# Добавили в Dockerfile дополнительные команды, выполнение которых требует времени:
RUN apk update && apk add build-base

COPY daemon /usr/local/bin/daemon
COPY config.yml /etc/daemon/config.yml
