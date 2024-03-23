#!/bin/bash
#daniel: skript zum startet aller anwendungen
chmod u+x eventDrivenScaler.sh dockerScaler.sh dockerDeScaler.sh
touch countervar.txt
(( sh ./eventDrivenScaler.sh & ) & )
(( sh ./dockerDeScaler.sh & ) & )
