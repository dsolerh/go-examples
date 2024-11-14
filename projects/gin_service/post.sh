#!/bin/zsh

curl -X POST 'http://localhost:8080/recipes' \
 --header 'Content-Type: application/json' \
 --data-raw "$(cat $1)"