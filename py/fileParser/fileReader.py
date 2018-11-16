# fileToArray - Returns a file parsed into an Array of ints
def fileToArray(name):
    with open(name) as line:
        lines = line.readlines()
    int_array = [int(numeric_string) for numeric_string in lines]
    return int_array