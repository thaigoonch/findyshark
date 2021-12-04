#!/bin/bash

#PS4='$LINENO: '
# Get a random number between 1 and 100
my_number=$(( RANDOM % 100 + 1 )); #echo $my_number
stillSearching=1;
count=0
guess="n"
printf "        _________         .    .\
\n       (..       \_    ,  |\  /|\
\n        \       0  \  /|  \ \/ /\
\n         \______    \/ |   \  /\
\n            vvvv\    \ |   /  |\
\n            \^^^^  ==   \_/   |\
\n             \`\_   ===    \.  |\
\n             / /\_   \ /      |\
\n             |/   \_  \|     /\
\n                    \_______/   
\n"
until $stillSearching == 0; do
	read -p "Enter input: " guess
	if test $guess = "q"; then
		echo -e "goodbye.\n"
		$stillSearching=0;
	fi
	
done
echo -e "Correct! You got the answer ($my_number) in $count tries."
exit 0


