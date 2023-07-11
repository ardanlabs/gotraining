#!/bin/bash
declare -i counter=0

for filename in ./*.go; do
    echo
    echo
    echo "Executing $filename"
    go1.21rc2 run $filename
    counter=`expr $counter + 1`
done

echo "Executed $counter functions from the slices API"