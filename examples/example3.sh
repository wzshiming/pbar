#!/bin/sh

# Multiple progress bars
./mock.sh 1000 $(seq 20 | xargs) | pbar -s pad
