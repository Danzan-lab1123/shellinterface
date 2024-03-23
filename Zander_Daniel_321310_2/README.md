# bachelor-thesis-321310
Repository der Bachelorarbeit von Daniel Zander (Matrikelnummer 321310) im Wintersemester 2023/24
Inhalt: Integration automatisiertes Bereitstellverfahren, aufbauend auf der [Bachelorarbeit Andreas Ralph Schneider](https://github.com/hochschule-pforzheim/bachelor-thesis-320469.git) 

Bearbeitungszeitraum:  04.12.2023 bis einschließlich 04.03.2024

---
1. Konzeption des Traefik Middelware in Go
2. Traefik Middelware
4. Konfigurationen am Traefik - Proxy
5. Dockerfile und Dockercompose Dateien
6. Bash - Skripte

Einen Nummerierung der Dateien wird auf Ordnerebene umgesetzt.

---
**Inbetriebnahme:**

Für die Inbetriebnahme des Ergebnisses der Auftragsstellung kann der Traefik mit integrierter Middleware als Dockerimage aus dem Docker-Hub Repositorium:
[https://hub.docker.com/repository/docker/danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel/general](https://hub.docker.com/r/danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel) heruntergeladen und ausgeführt werden.

```bash
foo@bar:~$ docker pull danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-v1.0
foo@bar:~$ docker run -d -p 8080:8080 -p 80:80 -v /var/run/docker.sock:/var/run/docker.sock danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-v1.0\
  --api.insecure=true \
  --providers.docker \ 
  --log.level=DEBUG \
  --experimental.localPlugins.shellinterface.moduleName=github.com/hochschule-pforzheim/bachelor-thesis-321310/shellinterface
```
      
Alternativ kann die Inbetriebnahme durch ein erneutes bauen des Traefik - Server erreicht werden.
Hier dargestellt am Beispiel der ersten Version:
1. Herunterladen des Repositorium
```bash
foo@bar:~$ git clone https://${GitHub-Token}@github.com/hochschule-pforzheim/bachelor-thesis-321310.git
```
2. In das Verzeichnis ["02-traefik"](https://github.com/hochschule-pforzheim/bachelor-thesis-321310/tree/main/02-traefik) navigieren.
```bash
foo@bar:~$ cd bachelot-thesis-321310/02-traefik/. 
```
3. Ausführen des Befehls zum erstellen eines Images des Traefik - Servers mittels Docker-Client.
```bash
foo@bar:~$ docker build -f Dockerfile --tag traefik:predproduction .
```
     
Das erstellte Image enthält die Bash-Skripte des Plugins, der Golang Programmcode wird in den darauffolgenden Schritten eingefügt.

4. In das Verzeichnis ["03-traefik-middleware"](https://github.com/hochschule-pforzheim/bachelor-thesis-321310/tree/main/03-traefik-middleware) navigieren.
```bash
foo@bar:~$ cd ..
foo@bar:~$ cd bachelot-thesis-321310/03-traefik-middleware/. 
```
5. Ausfürung der docker-compose.yml innerhalb des Verzeichnisses mittels Docker-Client.
```bash
foo@bar:~$ docker build -f Dockerfile --tag ghcr.io/hochschule-pforzheim/bachelor-thesis-321310/traefik:hochschule-pforzheimv1.0 .
```
       
Als zweite Alternative kann die Inbetriebnahme durch das Ausführen des Skriptes "build-traefik-skript.sh" im Repositorium erreicht werden.
```bash
foo@bar:~$ sh ./build-traefik-skript.sh
```
***
Beispielhafte Inbetriebnahme des Traefik - Servers mit Middleware - Als Service mit Docker Compose:
```YAML
version: "3.8"

services:
  traefik-hochschule-pforzheim:
    image: danzandoc1123/winter_2023-the4999-abschlussarbeit-bachelor_of_science-zander_daniel:traefik-hochschule-pforzheim-produktive
    deploy:
     replicas: 1
     placement:
       constraints: [node.role == manager]
    ports:
      - "80:80"
      - "8080:8080"
    command:
      - --log.level=INFO
      - --api.insecure=true
      - --providers.docker=true
      - --providers.docker.swarmMode=true
      - --providers.docker.exposedbydefault=false
      - --providers.docker.network=traefik
      - --entrypoints.web.address=:80
      - --experimental.localPlugins.shellinterface.moduleName=github.com/hochschule-pforzheim/bachelor-thesis-321310/shellinterface
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - traefik

networks:
  traefik:
    driver: overlay
    attachable: true
    name: traefik
```
***
Beispielhafte Inbetriebnahme des Services mit Docker Compose:
```YAML
version: "3.8"

services:
 nginx:
   image: nginx:latest
   deploy:
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.nginx.rule=Host(`nginx.docker.localhost`)"
      - "traefik.http.routers.nginx.entrypoints=web"
      - "traefik.http.routers.nginx.middlewares=shellinterfacemiddleware"
      - "traefik.http.middlewares.shellinterfacemiddleware.plugin.shellinterface.headers.serviceName=stack02_nginx"
      - "traefik.http.middlewares.shellinterfacemiddleware.plugin.shellinterface.headers.serviceMax=25"
      - "traefik.http.middlewares.shellinterfacemiddleware.plugin.shellinterface.headers.serviceMin=5"
      - "traefik.http.middlewares.shellinterfacemiddleware.plugin.shellinterface.headers.requestCounterCap=45"
      - "traefik.http.middlewares.shellinterfacemiddleware.plugin.shellinterface.headers.serviceTimerCap=300"
      - "traefik.http.services.nginx.loadbalancer.server.port=80"
   networks:
     - default
     - traefik

networks:
  traefik:
    external: true
```
