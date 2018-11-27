#!/bin/sh

# Save the original log
./mock.sh 30 hello | tee -a log.txt | pbar
