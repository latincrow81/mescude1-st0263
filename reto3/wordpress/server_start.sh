#!/bin/bash

sudo mount 10.128.0.6:/var/nfs/general /mnt/wordpress
sudo docker-compose -f ~/mescude1-st0263/reto3/wordpress/docker_compose.sh up -d