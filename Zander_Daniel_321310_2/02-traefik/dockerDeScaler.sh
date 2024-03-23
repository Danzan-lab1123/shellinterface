#!/bin/bash
#daniel:kontinuierlich reduzierung aller services, hier alle 5min/300sec

while true; do
#daniel:timer alle 5min/300sec
        sleep 300
		
#daniel:kontroller aller service mittels .txt datei,wie auch entnahme informationen
        cat countervar.txt | while read line; do
        servicename=$(cat countervar.txt | grep $line | cut -d '=' -f 2)
        servicemin=$(cat countervar.txt | grep  $line | cut -d '=' -f 1)
        calunit=$(cat countervar.txt | grep $line | cut -d '=' -f 3)

#daniel:anpassung der recheneinheit "calunit"
        calunit=$((calunit-1))

#daniel:prüfen ob bereits ob recheneinheit bereits die mindestanzahl erreicht ist
        if [ $calunit -lt $servicemin ];then
        calunit=$servicemin
        fi
		
#daniel:ausführen des dockerbefehls
        dcommand=$servicename"="$calunit
        changevar=$servicemin"="$dcommand
        docker service scale $dcommand -d
		
#daniel:anpassungen der werte in der .txt datei 
        sed -i s/$line/$changevar/g countervar.txt
                done;
done;
