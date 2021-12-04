#!/bin/bash

searchContents () {
  # Function to search files in a dir hierarchy
  # & display results that match a pattern in egrep.
  # -----------------------------------------------
  #  var $1 is the file type criteria
  #  var $2 is the hash representing space
  #  var $3 is the hash representing tab
  #  var $4 is the search input
  #  var $5 is the exclude params, acquired from the config
  
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
}

searchCaseInsensitive () {
  fileType=${1//:/*}
  hash1=$2
  hash2=$3
  orig=$4
  mod=${orig//$hash1/ }    # replace space hash(es) with space(s)
  final=${mod//$hash2/	}  # replace tab hash(es) with tab(s)
  if [ "$#" -ne 5 ]; then
    find . -iname "$fileType" -print0 | xargs -0 egrep -nsi "$final"
  else
    find . -iname "$fileType" -print0 | xargs -0 egrep -nsi "$final" | egrep -v $5
  fi
  exit 0
}

searchForFileExt () {
  find . -iname "*.$1"
}

searchFileNames () {
  find . -iname "*" -not -path "*/vendor/*" | grep $1
}
