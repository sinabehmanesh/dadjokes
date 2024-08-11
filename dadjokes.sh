#!/bin/bash

#color codes for level output
RED='\033[0;31m'
NC='\033[0m'
Blue='\033[1;36m'

#Get os name
osType=$(uname -o)

function dadjoke_linux() {
	echo "Linux dadjoke!"
}
function dadjoke_macos() {
	title="Dad Joke"	
	joke=$(curl -s https://icanhazdadjoke.com)
	iconfile="dad.png"
	iconpath="/Users/sina/Documents/linux-tools/dad.png"
	osascript <<EOF
display notification "$joke" with title "$title" sound name "default"
EOF
}

if [ "$osType" = "GNU/Linux" ]; then
	if command -v notify-send > /dev/null; then
		dadjoke_linux
	else
		echo -e "${RED}ERROR: notify-send package not present${NC}"
		echo -e "Install it using:
		${BLUE}apt-get install notify-send${NC}"
		exit 1
	fi
elif [ "$osType" = "Darwin" ]; then
	if command -v osascript > /dev/null; then
		dadjoke_macos
	else
		echo -e "${RED}ERROR: osascript not found!${NC}"
		exit 1
	fi
else
	echo "OS not supported!"
fi
