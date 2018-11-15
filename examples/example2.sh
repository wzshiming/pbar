#!/bin/sh

# Choose a different style
./mock.sh 100 normal | pbar
./mock.sh 100 pad | pbar -s pad
