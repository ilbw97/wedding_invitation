#!/bin/bash

# Check if two arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: ./run.sh <number> <language>"
    exit 1
fi

# Assign arguments to variables
NUM=$1
LANG=$2

# Call make with the provided arguments
make run NUM=$NUM LANG=$LANG
