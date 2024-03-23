#!bin/bash
#einfacher test zur konntrolle auf die funktionalität des traefik-servers, einsatz auf den unten dargestellten anwendungsfall beschränkt
apk add curl
servicetimercap=""
requestcountercurrent=""
servicecurrent=""
testrequest=""
expservicenumber=""

readService() {
        printf "Aktualisiere Service, warten Sie kurz"
        curl -s 'http://localhost:9091' > /dev/null
        sleep 25
        line=$(cat countervar.txt | grep traefik-proxy02_whoamit.* )
        servicemin=$(cat countervar.txt | grep  $line | cut -d '=' -f 1 )
        servicetimercap=$(cat countervar.txt | grep $line | cut -d '=' -f 4 )
        requestcountercap=$(cat countervar.txt | grep $line | cut -d '=' -f 6 )
        servicecurrent=$(docker service ls | grep traefik-proxy02_whoamit.* | awk '{print $4}' | awk -F'/' '{print $2}')
        testrequest=$((requestcountercap * 2)) && expservicenumber=$((servicecurrent+2))
        changevar=$servicemin"=traefik-proxy02_whoamit="$servicecurrent"="$servicetimercap"=0="$requestcountercap"=0"
        sed -i s/$line/$changevar/g countervar.txt
}

echo -e "Einfacher Test der automatisierten Skalierung des Traefik Servers mittels Bash-Skript"
echo -e "Test basierend auf docker-compose-test-skalierung.yml"
readService
echo -e "Servicename: traefik-proxy02_whoamit"
echo -e "Anzahl Anfragen für einen Aufwärtsskalierung: " $requestcountercap
echo -e "Vergangene Zeit in Sekunden für einen Abwärtsskalierung: " $servicetimercap

echo -e "Frage den Service " $testrequest "nach!"
i=0
j=$((testrequest+1))
while [ $((i)) -lt $((j)) ];do curl -s 'http://localhost:9091' > /dev/null && sleep 2 && echo $i && i=$((i+1)); done

sleep 10 && line=$(docker service ls | grep traefik-proxy02_whoamit.* | awk '{print $4}' | awk -F'/' '{print $2}')
echo -e "Vergleich Testwerte: \nAnzahl an Services zu Beginn: " $servicecurrent " Erwartet Anzahl an Services: " $expservicenumber " Anzahl an Service: " $line

echo -e "Abwartet für einen Abwärtsskalierung"
i=0
j=$((servicetimercap))
while [ $((j)) -gt $((0)) ];do sleep 1 && echo $j && j=$((j-1)); done
sleep 10
line=$(docker service ls | grep traefik-proxy02_whoamit.* | awk '{print $4}' | awk -F'/' '{print $2}')
echo -e "Vergleich Testwerte: \nErwartet Anzahl an Services: " $((expservicenumber-1)) " Anzahl an Service: " $line
echo -e "Danke für für die Nutzung des Clients"
apk del curl
