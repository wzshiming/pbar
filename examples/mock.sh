#!/bin/sh

for i in {"hello","world"}; do
	./mock_log.sh $i $(($RANDOM % 1000 + 100)) &
done
