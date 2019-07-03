#!/bin/bash
APPNAME="article-server"

sudo docker stop $APPNAME
sudo docker rm $APPNAME

sudo docker run -d --name $APPNAME -p "8080:8080" article-server
