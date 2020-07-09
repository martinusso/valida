package valida

import (
	"errors"
	"reflect"
)

type validator struct {
	allowEmpty bool
	value      interface{}
	typeOf     string
	messages   []string
}

func Forge(value interface{}) validator {
	var typeOf string
	if value != nil {
		typeOf = reflect.TypeOf(value).String()
	}
	return validator{
		value:  value,
		typeOf: typeOf,
	}
}

func (v validator) AllowEmpty() validator {
	v.allowEmpty = true
	return v
}

func (v validator) Assert() error {
	if len(v.messages) == 0 {
		return nil
	}
	return errors.New(v.messages[0])
}

func (v validator) Errors() (all []error) {
	if len(v.messages) == 0 {
		return
	}
	for _, m := range v.messages {
		all = append(all, errors.New(m))
	}
	return
}

func (v validator) Contains(arr []string, m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !Contains(arr, s) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Document(doc DocumentType, m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !Document(doc, s) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Email(m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !Email(s) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Empty(m string) validator {
	s := v.getString()
	if Empty(s) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Max(l int, m string) validator {
	if !Max(v.value.(int), l) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) MaxLength(l int, m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !MaxLength(s, l) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Min(l int, m string) validator {
	if !Min(v.value.(int), l) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) MinLength(l int, m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !MinLength(s, l) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) Password(m string) validator {
	s := v.getString()
	if v.allowEmpty && Empty(s) {
		return v
	}
	if !Password(s) {
		v.messages = append(v.messages, m)
	}
	return v
}

func (v validator) getString() string {
	if v.value == nil {
		return ""
	}
	switch v.typeOf {
	case "string":
		return v.value.(string)
	case "*string":
		s := v.value.(*string)
		if s == nil {
			return ""
		}
		return *s
	}
	return ""
}
