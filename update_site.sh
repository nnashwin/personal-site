#!/bin/bash
cd $HOME/personal-site
git pull origin master
caddy reload
make rrefresh
