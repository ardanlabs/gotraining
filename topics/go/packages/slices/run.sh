#!/bin/bash
counter=0

for filename in ./*.go; do
    echo
    echo
    echo "Executing $filename"
    $GO_COMMAND run $filename
    counter=`expr $counter + 1`
done

echo "Executed $counter functions from the slices API"