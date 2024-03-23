#!/bin/bash
#daniel:skalierung eines spezifischen service, abhänging von argumenten
servicename=$1
servicemax=$2
servicemin=$3
servicetimercap=$4
requestcountercap=$5

#daniel:deklaration der methode für die ausführung einer aufwärts sklalierung
scale() {
        calunit=$((calunit+1))
        if [ $((calunit)) -lt $((servicemax)) ] || [ $((calunit)) = $((servicemax)) ];then
                doScale
        fi
}

#daniel:sklalierung eines services, und anpassung der werte in der .txt-datei
doScale() {
        requestcountercurrent=0
        dcommand=$servicename"="$calunit
        changevar=$servicemin"="$dcommand"="$servicetimercap"="$servicetimercurrent"="$requestcountercap"="$requestcountercurrent
        docker service scale $dcommand -d
        sed -i s/$line/$changevar/g countervar.txt
}

#daniel:update des swarm wenn services getrennt vom treafik redeployed wird
updateSwarm() {
        service_current=$(docker service ls | grep $servicename | awk '{print $4}' | awk -F'/' '{print $2}')
        if [ $((service_current)) -lt $((calunit)) ];then
                doScale
        fi
}

#daniel:kontroller ob service bereits angelegt mittels .txt datei
serviceup="false"
cat countervar.txt | grep -o $servicename.* >/dev/null && serviceup="true"

#daniel:erst anlage des service wenn nicht bereits vorhanden innerhalb der .txt datei
case $serviceup in
	"false")
		echo $servicemin"="$servicename"="$servicemin"="$servicetimercap"=0="$requestcountercap"=0" >> countervar.txt
		dinitcommand=$servicename"="$servicemin
		docker service scale $dinitcommand -d
	;;

#daniel:entnahme informationen aus .txt datei
	"true")
		line=$(cat countervar.txt | grep $servicename.*)
		servicetimercap=$(cat countervar.txt | grep $line | cut -d '=' -f 4)
		servicetimercurrent=$(cat countervar.txt | grep $line | cut -d '=' -f 5)

#daniel:anpassung der recheneinheit "calunit"
		calunit=$(cat countervar.txt | grep $line | cut -d '=' -f 3)
		requestcountercap=$(cat countervar.txt | grep $line | cut -d '=' -f 6)
		requestcountercurrent=$(cat countervar.txt | grep $line | cut -d '=' -f 7)
		requestcountercurrent=$((requestcountercurrent+1))
		changevar=$servicemin"="$servicename"="$calunit"="$servicetimercap"="$servicetimercurrent"="$requestcountercap"="$requestcountercurrent
			updateSwarm
#daniel:prüfen ob die recheneinheit bereits die maximalanzahl erreicht hat, wenn nein > führe skalierung aus
		if [ $((requestcountercurrent)) -gt $((requestcountercap)) ] || [ $((requestcountercurrent)) = $((requestcountercap)) ];then
			scale
		fi
		sed -i s/$line/$changevar/g countervar.txt
		;;
esac