#!/bin/sh

# If addrs.json doesn't exist, then run data.
if [[ ! -e "addrs.json" ]]
then echo "addrs.json does not exist." && exec $GOPATH/src/rhodopsin-micros/data/data
# If addrs.json does exits, then don't do anything
elif [[ -e "addrs.json" ]]
then echo "addrs.json exists. not pulling new data."
fi


