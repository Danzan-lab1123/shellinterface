#!/bin/bash
#daniel:testet aktuelles verzeichniss jede sekunden auf einen die "scaleRequest.sh" erstellt durch die "shellinterface.go" -plugin
while true; do
        sleep 0.2
        if [ -e scaleRequest.sh ]
        then
                chmod u+x ./scaleRequest.sh
                sh ./scaleRequest.sh
                rm ./scaleRequest.sh
        fi
done;
