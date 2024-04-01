#!/bin/bash
#
cp .env.example .env
sudo apt update
sudo apt autoremove docker*
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo groupadd docker
sudo usermod -aG docker $USER
sudo docker compose up --build -d
sudo docker ps -a
