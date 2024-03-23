#!/bin/bash
#daniel:kontinuierlich reduzierung aller services, abhänging von .txt-datei
#daniel:deklaration der methode für die ausführung einer abwärts sklalierung

#daniel:prüfen ob bereits ob recheneinheit bereits die mindestanzahl erreicht ist
deScale() {
        calunit=$((calunit-1))
        servicetimercurrent=0
        requestcountercurrent=0
        dcommand=$servicename"="$calunit
        changevar=$servicemin"="$dcommand"="$servicetimercap"="$servicetimercurrent"="$requestcountercap"="$requestcountercurrent
        if [ $((calunit)) -lt $((servicemin)) ];then
                updateData
        else doDeScale
        fi
}

#daniel:sklalierung eines services, und anpassung der werte in der .txt-datei
doDeScale() {
        docker service scale $dcommand -d
        sed -i s/$line/$changevar/g countervar.txt
}

#daniel: update die daten auch ohne skalierung
updateData() {
        calunit=$((servicemin))
        dcommand=$servicename"="$calunit
        changevar=$servicemin"="$dcommand"="$servicetimercap"="$servicetimercurrent"="$requestcountercap"="$requestcountercurrent
        sed -i s/$line/$changevar/g countervar.txt
}

#daniel:kontinuierlich kontrolle aller services, hier jede sekunde
while true; do
        sleep 1

#daniel:kontroller aller service mittels .txt datei,wie auch entnahme informationen
        cat countervar.txt | while read line; do
        servicename=$(cat countervar.txt | grep $line | cut -d '=' -f 2)
        servicemin=$(cat countervar.txt | grep  $line | cut -d '=' -f 1)
        calunit=$(cat countervar.txt | grep $line | cut -d '=' -f 3)
        servicetimercap=$(cat countervar.txt | grep $line | cut -d '=' -f 4)
        servicetimercurrent=$(cat countervar.txt | grep $line | cut -d '=' -f 5)
        requestcountercap=$(cat countervar.txt | grep $line | cut -d '=' -f 6)
        requestcountercurrent=$(cat countervar.txt | grep $line | cut -d '=' -f 7)

#daniel: aktualisierung der referenzzeit
        servicetimercurrent=$((servicetimercurrent+1))
        changevar=$servicemin"="$servicename"="$calunit"="$servicetimercap"="$servicetimercurrent"="$requestcountercap"="$requestcountercurrent

#daniel: abhänging von wert der referenzzeit wird die sklalierung aufgerufen
        if [ $((servicetimercurrent)) -gt $((servicetimercap)) ];then
        deScale
        fi
        sed -i s/$line/$changevar/g countervar.txt
        done;
done;
