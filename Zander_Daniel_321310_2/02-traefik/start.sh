#!/bin/bash
#daniel: skript zum startet aller anwendungen
rm countervar.txt
touch countervar.txt
chmod u+x eventDrivenScaler.sh dockerScaler.sh dockerDeScaler.sh test.sh
(( sh ./eventDrivenScaler.sh & ) & )
(( sh ./dockerDeScaler.sh & ) & )
