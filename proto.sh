#!/bin/bash

ROOT=$GOPATH
echo $ROOT

for element in  `ls ${ROOT}`
do
    if [ -d "$ROOT/$element" ]; then echo $element
    fi

done

    
  