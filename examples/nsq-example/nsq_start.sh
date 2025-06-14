#!/bin/bash

ip=$(ifconfig | grep -w 'inet' | awk 'NR==2 {print $2}')

echo "host ip is: $ip"

sed -i '' "s/localhost/${ip}/g" docker-compose.yml

docker compose up -d