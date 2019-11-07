#!/bin/bash
echo $(pwd)/volume
docker run --rm -d -p 27017:27017 -v $(pwd)/volume:/data/db mongo
docker run --rm -d -p 6379:6379 redis
