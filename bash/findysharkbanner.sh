#!/bin/bash
config=$1
ignores=$(cat $config | grep ignore)
ignores=${ignores//"ignore: "/}
GRN='\033[1;32m'
GRYI='\033[7;34m'
GRY='\033[1;30m'
NC='\033[0m' # No Color
printf "        ${GRY}_________         .    .${NC}\
\n       ${GRY}(..       \_    ,  |\  /|${NC}    ${GRN}+-+-+-+-+-+-+            dooo${NC}\
\n        ${GRY}\       ${NC}0${GRY}  \  /|  \ \/ /${NC}    ${GRN}|findy|shark|                    doo doo${NC}\
\n         ${GRY}\______    \/ |   \  /${NC}     ${GRN}+-+-+-+-+-+-+      doo                        da-doo${NC}\
\n           vvvvv${GRY}       |  /   | ${NC}\
\n            ${NC}^^^^${GRY}       \_/    ) ${NC}    Using config file: ${GRYI}$config${NC}\
\n             ${GRY}\`\_   )))       /${NC}\
\n             ${GRY}/ /\_   \ /    /${NC}       Ignoring: ${GRYI}$ignores${NC}\
\n             ${GRY}|/   \___\|___/${NC}\
\n                    ${GRY}   v${NC}   
\n"
exit 0
