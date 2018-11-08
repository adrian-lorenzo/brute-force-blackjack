#!/bin/bash
echo "Creating binaries..."
go build -o "brute-force"
if [ $? ]; then
    echo "Done! Execute the program with './brute-force <card number> <number of cards>/-intense'"
    exit 0
else 
    echo "We could not build the code. Check for issues, please!"
    exit 1
fi