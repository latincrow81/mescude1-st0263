#!/bin/bash

sudo mount 10.128.0.6:/var/nfs/general /mnt/wordpress
sudo docker-compose up -d
