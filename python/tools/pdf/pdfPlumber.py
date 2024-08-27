import pdfplumber
from utils.func import *

def process(inputFile):
	rawText = ""
	with pdfplumber.open(inputFile) as pdf:

		for page in pdf.pages:
			context = page.extract_text()
			# print(context)
			text = scrapingLogic(context)
			rawText = rawText+text

	return rawText