FROM ghcr.io/hochschule-pforzheim/bachelor-thesis-321310/traefik-enc-entrypoint:latest
RUN apk update && apk upgrade
RUN set -ex && apk --no-cache add sudo
RUN apk add docker
COPY . /
RUN chmod u+x ./start.sh
