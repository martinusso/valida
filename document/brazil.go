package document

import (
	docCNPJ "github.com/padmoney/go-docs/cnpj"
	docCPF "github.com/padmoney/go-docs/cpf"
)

func CNPJ(v string) bool {
	return docCNPJ.Valid(v)
}

func CPF(v string) bool {
	return docCPF.Valid(v)
}
