**Einfacher test zur konntrolle auf die funktionalität des traefik-servers, einsatz auf den unten dargestellten Anwendungsfall beschränkt:**


Ausführung des Test wenn ein Traefik - Server mit integrierteter Middleware bereits vorhanden:

1. Starten des Testservices folgenden Befehl:

```bash
foo@bar:~$ docker stack deploy -c docker-compose-traefik-proxy03.yml traefik-proxy03
```
2. Ausführen des Test:

```bash
foo@bar:~$ container_id=$(docker container ls | grep danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-produktiv | awk '{print $1}' ); docker exec -it $container_id sh ./test.sh
```
Ausführung des Test wenn kein Traefik - Server mit integrierteter Middleware vorhanden ist:
1. Deployment des folgenden Stack mit Befehl:

Docker Compose YML:
```YML
version: "3.8"
services:
  traefik-hochschule-pforzheim:
    image: danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-test
    deploy:
     replicas: 1
     placement:
       constraints: [node.hostname == clusterpi01] 
    ports:
      - "80:80"
      - "8080:8080"
      - "9091:9091"
    command:
      - --log.level=INFO
      - --api.insecure=true
      - --providers.docker=true
      - --providers.docker.swarmMode=true
      - --providers.docker.exposedbydefault=false
      - --providers.docker.network=traefik
      - --entrypoints.web.address=:80
      - --entrypoints.whaomi.address=:9091
      - --experimental.localPlugins.shellinterface.moduleName=github.com/hochschule-pforzheim/bachelor-thesis-321310/shellinterface
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - traefik
  whaomi:
   image: ghcr.io/traefik/whoami:v1.10.1
   deploy:
    labels:
      - "traefik.enable=true"
      - "traefik.docker.lbswarm=true"
      - "traefik.http.routers.whoami.rule=Path(`/`)"
      - "traefik.http.routers.whoami.entrypoints=whaomi"
      - "traefik.http.routers.whoami.middlewares=shell-whoami"
      - "traefik.http.middlewares.shell-whoami.plugin.shellinterface.headers.serviceName=stack-test_whaomi"
      - "traefik.http.middlewares.shell-whoami.plugin.shellinterface.headers.serviceMax=15"
      - "traefik.http.middlewares.shell-whoami.plugin.shellinterface.headers.serviceMin=5"
      - "traefik.http.middlewares.shell-whoami.plugin.shellinterface.headers.requestCounterCap=5"
      - "traefik.http.middlewares.shell-whoami.plugin.shellinterface.headers.serviceTimerCap=150"
      - "traefik.http.services.whoami.loadbalancer.server.port=80"
   networks:
     - default
     - traefik
networks:
  traefik:
    driver: overlay
    attachable: true
    name: traefik
```

```bash
foo@bar:~$ docker stack deploy -c docker-compose.yml stack-test
```
2. Ausführen des Test:
```bash
foo@bar:~$ container_id=$(docker container ls | grep danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-test | awk '{print $1}' ); docker exec -it $container_id sh ./test.sh
```

