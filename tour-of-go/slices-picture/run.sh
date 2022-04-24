#!/bin/sh

go run . > base64.txt

base64Str=`cat base64.txt`
base64Str=${base64Str:6} # Removes 'IMAGE:' from the start of the string

echo "$base64Str" | base64 -d > picture.jpg