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
import csv, codecs

class WordCount:
    def __init__(self):
        self.m = {}
    
    def add(self, word, count):
        if self.m.get(word):
            self.m[word] += count
            return
        self.m[word] = 1
    def export(self, path):
        # Storing and saving as CSV
        rows = [[word, count] for word, count in self.m.items()]
        with open(path, 'w') as f:
            writer = csv.writer(f)
            writer.writerows(rows)
            f.close()

def get_files(path:str):
    onlyfiles = [f for f in listdir(path) if isfile(join(path, f))]
    return onlyfiles

def execute():
    
    wc = WordCount()
    files = get_files(path)
    for file in files:
        with codecs.open(path, encoding='utf-8') as f: # Fix encoding issues
            for line in f:
                words = line.lower().split()
                wc.add(word.lower(), 1)
            f.close()
    wc.export(output)

if __name__ == "__main__":    
    execute()