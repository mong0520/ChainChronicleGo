#!/bin/bash

CONFIG=$1
while true
do
	go run main.go -c $CONFIG -a status -r 1 | grep 精靈石
    sleep 3
done
