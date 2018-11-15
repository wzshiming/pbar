#!/bin/sh

# Save the original log
./mock.sh 10 hello | tee -a log.txt | pbar
cat log.txt
