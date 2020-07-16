package valida

import "github.com/martinusso/valida/document"

type DocumentType string

const (
	BrazilCNPJ DocumentType = "brazil/cnpj"
	BrazilCPF  DocumentType = "brazil/cpf"
)

func Document(docType DocumentType, value string) bool {
	if Empty(value) {
		return false
	}
	switch docType {
	case BrazilCNPJ:
		return document.CNPJ(value)
	case BrazilCPF:
		return document.CPF(value)
	}
	return false
}
