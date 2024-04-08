package service

import (
	"document-scraping-with-ai/business/operation"
	"document-scraping-with-ai/model"
	"os/exec"
)

type service struct {
	pdfConfig model.PdfConfig
	command *exec.Cmd
}


func NewPDFService(pdfConfig model.PdfConfig,command *exec.Cmd) oparation.Service {
	return &service{
		pdfConfig: pdfConfig,
		command: command,
	}
}