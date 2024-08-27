from os import walk
from tools.pdf.pyPdf2 import process as pypdf2
from tools.pdf.pdfPlumber import process as pdfplumber

# Open directory

files = []

for (dirpath, dirnames, filenames) in walk("../assets/"):
    files.extend(filenames)
    break

processFile = 0
doneFile = 0

for file in files:
    if ".pdf" not in file: 
        continue

    readFileName = dirpath+file
    writeFileName = dirpath+"/"+file.replace(".pdf", "_raw.txt")

    # pypdf2(readFileName, writeFileName)
    rawText = pdfplumber(readFileName)

    processFile = processFile + 1

    if len(rawText) > 0:
        # Write file
        with open(writeFileName, 'w', encoding='utf-8') as f:
            print("Write File : ",writeFileName)
            f.write(rawText)
            doneFile = doneFile + 1

print("done",doneFile,"/",processFile)