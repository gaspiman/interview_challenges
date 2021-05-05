
"""
Given:
- Server: 16 Cores CPU, 32GB RAM, 2TB SSD
- 50 text files with avg size of 48GB
- Example file content: List of paragraphs with variable length 
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
            """ 
            # Old school approach
            line = f.readline()
            while line:
                words = line.lower().split()
            """
            for line in f:
                words = line.lower().split()
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