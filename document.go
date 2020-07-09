package validate

import "github.com/martinusso/valida/document"

type DocumentType string

const (
	brazilCNPJ DocumentType = "brazil/cnpj"
	BrazilCPF  DocumentType = "brazil/cpf"
)

func Document(docType DocumentType, value string) bool {
	if Empty(value) {
		return false
	}
	switch docType {
	case brazilCNPJ:
		return document.CNPJ(value)
	case BrazilCPF:
		return document.CPF(value)
	}
	return false
}
