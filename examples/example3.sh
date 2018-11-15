#!/bin/sh

# Multiple progress bars
./mock.sh 1000 $(crun -r -l 10 "(Hello|Hi)_[a-z]{2,3}" | xargs) | pbar
