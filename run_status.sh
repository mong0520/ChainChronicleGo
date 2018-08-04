#!/bin/bash

CONFIG=$1
while true
do
	#go run main.go -c $CONFIG -a status -r 1 | grep 精靈石
	go run main.go -c $CONFIG -a status -r 1  | grep 轉蛋幣
    sleep 1
done
