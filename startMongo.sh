#!/bin/bash
echo $(PWD)/db
docker run -p 27017:27017 -v $(PWD)/db:/data/db mongo
