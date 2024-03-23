#!/bin/bash
while true; do
	if [ -e dockerscaler.sh ]
	then
		sh ./dockerscaler.sh
	else
		sleep 10

	fi
done;