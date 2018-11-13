#fileToArray - Parses from a text file, reading line by line, giving back an array of integers

def fileToArray(name):
    with open(name) as line:
        lines = line.readlines()
    int_array = [int(numeric_string) for numeric_string in lines]
    return int_array
    