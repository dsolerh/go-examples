#!/bin/zsh

curl -X PUT "http://localhost:8080/recipes/$1" \
 --header 'Content-Type: application/json' \
 --data-raw "$(cat $2)"