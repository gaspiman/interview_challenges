"""
Given:
- Input folder containing text files
- Count the number of time each word is repeated
- Export the results in a CSV file (ex: word, 10)
"""
path = "/path/to/folder"
output = "/path/to/output/file.csv" 

from os import listdir
from os.path import isfile, join
import csv

def get_files(path:str):
    onlyfiles = [f for f in listdir(path) if isfile(join(path, f))]
    return onlyfiles

def execute():
    m = {} # Dict to store all word: count data
    files = get_files(path)
    for file in files:
        with open(file) as f:
            content = f.read() # Reads the entire content - some files will not fit in memory. Readlines method is better wuited
            words = content.lower().split()
            for word in words:
                if m.get(word):
                    m[word] += 1
                    continue
                m[word] = 1
            f.close()
    output = "/tmp/somefile.csv"
    # Storing and saving as CSV
    rows = [[word, count] for word, count in m.items()]
    with open(output, 'w') as f:
        writer = csv.writer(f)
        writer.writerows(rows)
        f.close()


if __name__ == "__main__":    
    execute()