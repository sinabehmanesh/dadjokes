#!/bin/bash

#color codes for level output
RED='\033[0;31m'
NC='\033[0m'
Blue='\033[1;36m'

#check if notify-send exists on this machine!

if ! command -v notify-send  $> /dev/null
then
	echo -e "${RED}ERROR: notify-send package not present${NC}"
	echo -e "Install it using:
	${Blue} apt-get install notify-send${NC}"
	exit 1
fi

joke=$(curl -s https://icanhazdadjoke.com)
iconpath=$(pwd)
echo $iconpath
notify-send -a dadjokes -u critical -i "${iconpath}/1151214.png"  "Joke" "${joke}"