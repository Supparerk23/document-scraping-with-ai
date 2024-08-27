from constants.keyword import *
from constants.tag import *

def scrapingLogic(text):
	rawText = ""
	if ISSUED_ON_DATE_KEYWORD in text:
		# print("have ISSUED_ON_DATE_KEYWORD on page")
		issued_on_date_index = text.index(ISSUED_ON_DATE_KEYWORD)
		last_issued_on_date_index = issued_on_date_index + 50
		rawText = ISSUED_ON_DATE_TAG.replace("{data}",text[issued_on_date_index:last_issued_on_date_index])
	# Reduce content in token by find keyword 
	if FEE_KEYWORD in text:
		rawText = rawText+RAW_TAG.replace("{data}",text)

	return rawText