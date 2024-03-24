package service

import (
	"document-scraping-with-ai/business/pdf"
	"document-scraping-with-ai/model"
	"os/exec"
)

type service struct {
	pdfConfig model.PdfConfig
	command *exec.Cmd
}


func NewPDFService(pdfConfig model.PdfConfig,command *exec.Cmd) pdf.Service {
	return &service{
		pdfConfig: pdfConfig,
		command: command,
	}
}