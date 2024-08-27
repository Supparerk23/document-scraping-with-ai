import PyPDF2
from constants.keyword import *

def process(inputFile,outputFile):
     # Open the PDF file
    pdf_file = open(inputFile, 'rb')
    # Read the PDF file
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    # Print the number of pages
    rawText = ""
    # print("# Start Extract File",inputFile)   
   
    for pageNumber, page in enumerate(pdf_reader.pages):
        # Extract the text from page
        text = page.extract_text()

        # if pageNumber == 0:
        #     # pureText = text.replace(' ', '')
        #     print(text)
        #     print("---------")
        if ISSUED_ON_DATE_KEYWORD in text:
            # print("have ISSUED_ON_DATE_KEYWORD on page",pageNumber)
            issued_on_date_index = text.index(ISSUED_ON_DATE_KEYWORD)
            last_issued_on_date_index = issued_on_date_index + 50
            # print("index",issued_on_date_index)
            # print(text[issued_on_date_index:last_issued_on_date_index])
            rawText = text[issued_on_date_index:last_issued_on_date_index]
        # Reduce content in token by find keyword 
        if FEE_KEYWORD not in text: 
            continue
        rawText = rawText+text
        # Print the text
        # print(text.encode("utf-8"))

    if len(rawText) == 0:
        # Close the PDF file
        pdf_file.close()
        print("Can't read file",inputFile)
    else: 
        # Write file
        with open(outputFile, 'w', encoding='utf-8') as f:
            # print("Write File",outputFile)
            f.write(rawText)
        # Close the PDF file
        pdf_file.close()