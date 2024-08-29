import PyPDF2
from constants.keyword import *
from utils.func import *

def process(inputFile):
    # Open the PDF file
    pdf_file = open(inputFile, 'rb')
    # Read the PDF file
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    # Print the number of pages
    rawText = ""
    # print("# Start Extract File",inputFile)   
   
    for pageNumber, page in enumerate(pdf_reader.pages):
        # Extract the text from page
        context = page.extract_text()

        text = scrapingLogic(context)
        rawText = rawText+text
    
    pdf_file.close()
    return rawText