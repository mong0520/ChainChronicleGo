#!/bin/bash
echo $(PWD)/volume
docker run --rm -d -p 27017:27017 -v $(PWD)/volume:/data/db mongo
docker run --rm -p 6379:6379 redis