name="Sgt. 'Sina' Behmanesh"
message="Status Check"
day=$(cat /home/sina.behmanesh/activedays)
date=$(date +%r)
final_date="Day ${day} - ${date}"
company="US second devision"
team=""
location="Munich, Germany"
status="Agent Status:"
starter="....."

GREEN='\033[3;92m'
STATE_GREEN='\e[1;32m'
BLINK='\033[5;32m'
NS='\033[0m'

for (( i=0; i<${#starter}; i++)); do
        echo -n -e "${BLINK}${GREEN}${starter:$i:1}${NS}"
        sleep 0.08
done

printf "\n\n"

sleep 0.1

for (( i=0; i<${#message}; i++)); do
	echo -n -e "${GREEN}${message:$i:1}${NS}"
	sleep 0.04
done

sleep 0.1

printf "\n"

for (( i=0; i<${#final_date}; i++)); do
        echo -n -e "${GREEN}${final_date:$i:1}"
        sleep 0.03
done

sleep 0.1
printf "\n"

for (( i=0; i<${#name}; i++)); do
        echo -n -e "${GREEN}${name:$i:1}"
        sleep 0.03
done

sleep 0.1
printf "\n"

for (( i=0; i<${#company}; i++)); do
        echo -n -e "${GREEN}${company:$i:1}"
        sleep 0.03
done

sleep 0.1
printf "\n"

for (( i=0; i<${#team}; i++)); do
        echo -n -e "${GREEN}${team:$i:1}"
        sleep 0.03
done

sleep 0.1
printf "\n"

for (( i=0; i<${#location}; i++)); do
        echo -n -e "${GREEN}${location:$i:1}${NS}"
        sleep 0.03
done


sleep 0.1
printf "\n"

for (( i=0; i<${#status}; i++)); do
        echo -n -e "${GREEN}${status:$i:1}${NS}"
        sleep 0.03
done



echo -n -e "${BLINK}${STATE_GREEN} ACTIVE${NS}"

sleep 0.1
printf "\n\n"

