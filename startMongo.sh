#!/bin/bash
echo $(PWD)/volume
docker run --rm -p 27017:27017 -v $(PWD)/volume:/data/db mongo
