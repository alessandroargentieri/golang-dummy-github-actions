#! /bin/bash

# start echoserver
docker run --network host -it -d -e PORT=8080 --name echoserver cilium/echoserver:latest

# start ngrok
ngrok http 8080
