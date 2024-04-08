import PyPDF2
from os import walk

# Open directory

files = []
for (dirpath, dirnames, filenames) in walk("../assets/"):
    files.extend(filenames)
    break

for file in files:
    if ".pdf" not in file: 
        continue
    readFileName = dirpath+file
    writeFileName = dirpath+"raw/"+file.replace(".pdf", ".txt")
   
    # Open the PDF file
    pdf_file = open(readFileName, 'rb')
    # Read the PDF file
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    # Print the number of pages
    rawText = ""
    print("Start Extract File",readFileName)
    for pageNumber, page in enumerate(pdf_reader.pages):
        # Extract the text from page
        text = page.extract_text()
        # Reduce content in token by find keyword 
        if "ค่าธรรมเนียม" not in text: 
            continue
        rawText = rawText+text
        # Print the text
        # print(text.encode("utf-8"))

    # Write file
    with open(writeFileName, 'w', encoding='utf-8') as f:
        print("Write File",writeFileName)
        f.write(rawText)
    # Close the PDF file
    pdf_file.close()