#!/bin/bash
cd $HOME/personal-site
caddy reload
git pull origin master
make rrefresh
