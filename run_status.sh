#!/bin/bash

CONFIG=$1
while true
do
	go run main.go -c $CONFIG -a status -r 999999 -t 3
    sleep 3
done
