#!/bin/bash
#daniel:skalierung eines spezifischen service, abhänging von argumenten
servicename=$1
servicemax=$2
servicemin=$3
serviceup="false"

#daniel:kontroller ob service bereits angelegt mittels .txt datei
cat countervar.txt | grep -o $servicename.* >/dev/null && serviceup="true"
    case $serviceup in
    "false")

#daniel:erst anlage des service wenn nicht bereits vorhanden innerhalb der .txt datei
    echo $servicemin"="$servicename"="$servicemin >> countervar.txt
    docker service scale $servicename"="$servicemin -d
    ;;
    "true")

#daniel:entnahme informationen aus .txt datei
    changevar=$(cat countervar.txt | grep -o $servicename.*)
    calunit=$(cat countervar.txt | grep -o $servicename.* | cut -f2- -d=)

#daniel:anpassung der recheneinheit "calunit"
    calunit=$((calunit+1))
	
#daniel:prüfen ob bereits ob recheneinheit bereits die maximalanzahl erreicht ist
    if [ $calunit -gt $servicemax ];then
    calunit=$servicemax
    fi
	
#daniel:ausführen des dockerbefehls	
    dcommand=$servicename"="$calunit
    docker service scale $dcommand -d

#daniel:anpassungen der werte in der .txt datei 
    sed -i s/$changevar/$dcommand/g countervar.txt
    ;;
    esac
