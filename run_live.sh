#!/bin/sh

git pull

docker-compose -f docker-compose.live.yml up --build | tee log
