import PyPDF2
# Open the PDF file
pdf_file = open('./assets/M0009_2553.pdf', 'rb')
# Read the PDF file
pdf_reader = PyPDF2.PdfReader(pdf_file)
# Print the number of pages
print(len(pdf_reader.pages))
# Get the first page
page = pdf_reader.pages[0]
# Extract the text from the first page
text = page.extract_text()
# Print the text
# print(text.encode("utf-8"))
# Write file
with open('./assets/raw/test.txt', 'w', encoding='utf-8') as f:
    f.write(text)

# Close the PDF file
pdf_file.close()