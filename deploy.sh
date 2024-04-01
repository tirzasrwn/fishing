#!/bin/bash

set -x

cp .env.example .env
sudo apt update
cd ..
sudo apt autoremove docker* -y
cd ./fishing
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo groupadd docker
sudo usermod -aG docker $USER
sudo docker compose up --build -d
sudo docker ps -a
