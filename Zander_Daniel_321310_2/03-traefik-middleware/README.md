# Traefik - Server mit Programmcode für die Middleware
Verzeichnis mit docker-compose.yml und Programmcode für die Inbetriebnahme der entwickelte Middleware.


Anmerkung: 

Anzumerken ist das dies nur ein Teilschritt der Inbetriebnahme ist, zuvor ist ein Erstellen des Image aus dem Verzeichnis ["02-traefik"](https://github.com/hochschule-pforzheim/bachelor-thesis-321310/tree/main/02-traefik) notwendig.
Bediengsanleitung für die gesamte [Inbetriebnahme](https://github.com/hochschule-pforzheim/bachelor-thesis-321310/blob/main/README.md)

Befehl zur Inbetriebnahme des Traefik - Servers mit integrierte Middleware:
```bash
foo@bar:~$ docker build -f Dockerfile --tag ghcr.io/hochschule-pforzheim/bachelor-thesis-321310/traefik:hochschule-pforzheimv1.0 .
```
