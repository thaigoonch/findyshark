#!/bin/bash

fileType=${1//:/*}
hash1=$2
hash2=$3
orig=$4
mod=${orig//$hash1/ }    # replace space hash(es) with space(s)
final=${mod//$hash2/	}  # replace tab hash(es) with tab(s)
if [ "$#" -ne 5 ]; then
  find . -iname "$fileType" -print0 | xargs -0 egrep -ns "$final"
else
  find . -iname "$fileType" -print0 | xargs -0 egrep -ns "$final" | egrep -v $5
fi
exit 0