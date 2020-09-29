#!/bin/bash
sudo apt-get update
cd /usr/src/personal-site
sudo git pull origin master
sudo make rrefresh
