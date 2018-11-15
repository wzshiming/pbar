#!/bin/sh

key="${1:-"hello"}"
count="${2:-1000}"

for ((i = 0; i <= $count; i += $(($RANDOM % 10)))); do
	echo "$key $i/$count"
	sleep 0.0$(($RANDOM % 100))
done

echo "$key $count/$count"
