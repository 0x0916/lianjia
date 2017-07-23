#!/bin/bash
while read line
do
      ../0004-get-hushu $line
done < url.txt
