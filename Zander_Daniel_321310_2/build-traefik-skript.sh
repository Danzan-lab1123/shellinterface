#!bin/sh
Traefik1() {
cd 02-traefik
rm -r 00-erweiterung README.md
docker build -f Dockerfile --tag traefik:predproduction .
cd .. && cd 03-traefik-middleware && rm README.md LICENSE
docker build -f Dockerfile --tag ghcr.io/hochschule-pforzheim/bachelor-thesis-321310/traefik:hochschule-pforzheimv1.0 .
image_id=$(docker image ls --all | grep predproduction | awk '{print $3}')
docker image rm ${image_id?}
alrbuilt=1
}

Traefik2() {
cd 02-traefik
rm dockerScaler.sh dockerDeScaler.sh
cd 00-erweiterung && pfad00=$(dirname $PWD)
cp dockerDeScaler.sh dockerScaler.sh $pfad00
cd ..
rm -r 00-erweiterung README.md
docker build -f Dockerfile --tag traefik:predproduction .
cd .. && cd 03-traefik-middleware && rm README.md LICENSE shellinterface.go shellinterface_test.go
cd 00-erweiterung  && pfad01=$(dirname $PWD)
cp shellinterface.go shellinterface_test.go $pfad01 && cd ..
docker build -f Dockerfile --tag ghcr.io/hochschule-pforzheim/bachelor-thesis-321310/traefik:hochschule-pforzheimv2.0 .
image_id=$(docker image ls --all | grep predproduction | awk '{print $3}')
docker image rm ${image_id?}
alrbuilt=1
}

close() {
printf "Danke für die Nutzung des Clients\n"
exit=1
}

exit=0
alrbuilt=0
while [ $((exit)) = 0 ]; do
printf "Bauen des Traefik-Servers mittels Bash-Skript\n"
printf "Versionswahl:\ntraefik:hochschulepforzheimv1.0 = 1\ntraefik:hochschule-pforzheimv2.0 = 2\nVerlassen der Versionswahl = 0\n"
read -p "Eingabe Auswahl: " argument
case $argument in
  1)
  Traefik1
  close
  ;;
  2)
  Traefik2
  close
  ;;
  0)
  close
  ;;
  *) printf "Diese Eingabe wird nicht unterstützt! Bitte wählen Sie von den gegebenen Möglichkeiten\n";;
esac
done;

 if [ $((alrbuilt)) = 1 ];then
    cd $(dirname $PWD)
    rm -- "$0"
 fi
