package oparation

type Service interface {
	ProcessPDF() error
	WriteResult(path string, res string) error
}