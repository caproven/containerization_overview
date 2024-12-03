#!/usr/bin/env bash

[[ -z "$URL" ]] && echo "no URL env set" && exit 1

while true; do
  curl -iks -o /dev/null "$URL"
  result=$?
  echo "hitting url ${URL} gave result ${result}"
  sleep 3
done

exit 0
