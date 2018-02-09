#!/bin/bash

CONFIG=$1
while true
do
	go run main.go -c $CONFIG
    sleep 3
done
