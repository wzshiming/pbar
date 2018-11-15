#!/bin/sh

bars=${@:2}
base=$1

for i in $bars; do
	./mock_log.sh $i $(($RANDOM % $base + $base)) &
done
