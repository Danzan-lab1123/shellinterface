# Traefik - Server mit integration der Bash-Skripte
Verzeichnis mit Dockerfile und Bash-Skripte das ein Traefik - Server mit integration der Bash-Skripte für die Inbetriebnahme der entwickelte Inbetriebnahme.

Anmerkung:

Anzumerken ist das dies nur ein Teilschritt der Inbetriebnahme ist, somit ist die Middleware nach dem Erstellen des Images basierend auf dem Dockerfiles dieses Verzeichnisses
nicht funktionsfähig ist. Das Images dient für die Inbetriebnahme im Verzeichniss "03-traefik-middleware"
Bediengsanleitung für die gesamte [Inbetriebnahme](https://github.com/hochschule-pforzheim/bachelor-thesis-321310/blob/main/README.md)

Befehl zum Bauen des Traefik - Servers mit integrierte Middleware:
```bash
foo@bar:~$ docker build -f Dockerfile --tag traefik:predproduction .
```
