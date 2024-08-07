import PyPDF2
from os import walk

# Open directory

files = []
issued_on_date_keyword = "ข้อมูล ณ วันที่"
fee_keyword = "ค่าธรรมเนียม"

for (dirpath, dirnames, filenames) in walk("../assets/"):
    files.extend(filenames)
    break

for file in files:
    if ".pdf" not in file: 
        continue
    readFileName = dirpath+file
    writeFileName = dirpath+"/"+file.replace(".pdf", "_raw.txt")
   
    # Open the PDF file
    pdf_file = open(readFileName, 'rb')
    # Read the PDF file
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    # Print the number of pages
    rawText = ""
    # print("# Start Extract File",readFileName)
    for pageNumber, page in enumerate(pdf_reader.pages):
        # Extract the text from page
        text = page.extract_text()

        # if pageNumber == 0:
        #     # pureText = text.replace(' ', '')
        #     print(text)
        #     print("---------")
        if issued_on_date_keyword in text:
            # print("have issued_on_date_keyword on page",pageNumber)
            issued_on_date_index = text.index(issued_on_date_keyword)
            last_issued_on_date_index = issued_on_date_index + 50
            # print("index",issued_on_date_index)
            # print(text[issued_on_date_index:last_issued_on_date_index])
            rawText = text[issued_on_date_index:last_issued_on_date_index]
        # Reduce content in token by find keyword 
        if fee_keyword not in text: 
            continue
        rawText = rawText+text
        # Print the text
        # print(text.encode("utf-8"))

    if len(rawText) == 0:
        # Close the PDF file
        pdf_file.close()
        print("Can't read file",readFileName)
    else: 
        # Write file
        with open(writeFileName, 'w', encoding='utf-8') as f:
            # print("Write File",writeFileName)
            f.write(rawText)
        # Close the PDF file
        pdf_file.close()