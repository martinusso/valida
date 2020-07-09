package valida

import (
	"testing"
)

func TestAllowEmpty(t *testing.T) {
	expected := "invalid"
	err := Forge("").
		AllowEmpty().
		Email(expected).
		Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	got := Forge("").
		Empty(expected).
		Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestErrors(t *testing.T) {
	err := Forge("daenerys.stormborn#house-targaryen").
		AllowEmpty().
		Email("email/invalid").
		MaxLength(20, "email/max-length").
		Errors()
	if len(err) != 2 {
		t.Errorf("Expected 2 errors, got '%d'", len(err))
	}
}

func TestContains(t *testing.T) {
	os := []string{"linux", "windows", "macos"}

	expected := "contains/string-invalid"
	got := Forge("android").AllowEmpty().Contains(os, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge("linux").AllowEmpty().Contains(os, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge("").AllowEmpty().Contains(os, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestDocumentsBrazil(t *testing.T) {
	expected := "cnpj/invalid"
	got := Forge("11111111111111").AllowEmpty().Document(brazilCNPJ, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	err := Forge("24761136000182").AllowEmpty().Document(brazilCNPJ, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	err = Forge("").AllowEmpty().Document(brazilCNPJ, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	expected = "cpf/invalid"
	got = Forge("11111111111").AllowEmpty().Document(BrazilCPF, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	err = Forge("61434160521").AllowEmpty().Document(BrazilCPF, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	err = Forge("").AllowEmpty().Document(BrazilCPF, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestEmpty(t *testing.T) {
	expected := "empty/invalid"
	got := Forge("").Empty(expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	got = Forge(nil).Email(expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	var s string
	got = Forge(&s).Email(expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	s = "hello"
	got = Forge(&s).Email(expected).Assert().Error()
}

func TestEmail(t *testing.T) {
	expected := "email/invalid"
	got := Forge("1").AllowEmpty().Email(expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge("").AllowEmpty().Email(expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	var null *string
	err = Forge(null).AllowEmpty().Email(expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestMax(t *testing.T) {
	expected := "max/error"
	got := Forge(42).Max(41, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge(42).Max(42, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge(42).Max(43, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestMaxLength(t *testing.T) {
	expected := "maxLength/error"
	got := Forge("valida").AllowEmpty().MaxLength(5, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge("valida").AllowEmpty().MaxLength(6, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge("valida").AllowEmpty().MaxLength(7, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	err = Forge("").AllowEmpty().MaxLength(7, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestMin(t *testing.T) {
	expected := "min/error"
	got := Forge(42).Min(43, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge(42).Min(42, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge(42).Min(41, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestMinLength(t *testing.T) {
	expected := "minLength/error"
	got := Forge("valida").AllowEmpty().MinLength(7, expected).Assert().Error()
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	err := Forge("valida").AllowEmpty().MinLength(6, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge("valida").AllowEmpty().MinLength(5, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	err = Forge("").AllowEmpty().MinLength(5, expected).Assert()
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
}

func TestPassword(t *testing.T) {
	expected := "password/invalid"

	err := Forge("").AllowEmpty().Password(expected).Assert()
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	err = Forge("1234567890").AllowEmpty().Password(expected).Assert()
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
	err = Forge("Lc9!").AllowEmpty().Password(expected).Assert()
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
	err = Forge("123456Lc").AllowEmpty().Password(expected).Assert()
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
	err = Forge("abcdefLc!").AllowEmpty().Password(expected).Assert()
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}

	err = Forge("luk31@").AllowEmpty().Password(expected).Assert()
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
}
