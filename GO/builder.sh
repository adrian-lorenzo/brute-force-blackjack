#!/bin/bash
echo "Creating binaries..."
go build -o "blackjack"
if [ $? ]; then
    echo "Done! Execute the program with './blackjack <num-first-card> <num-to-take>/-intense'"
    exit 0
else 
    echo "We could not build the code. Check for issues, please!"
    exit 1
fi