# valida

[![Build Status](https://travis-ci.org/martinusso/valida.svg?branch=master)](https://travis-ci.org/martinusso/valida)
[![Coverage Status](https://coveralls.io/repos/github/martinusso/valida/badge.svg?branch=master)](https://coveralls.io/github/martinusso/valida?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinusso/valida)](https://goreportcard.com/report/github.com/martinusso/valida)

## Getting started

import valida

```
import "github.com/martinusso/valida"
```

Get the first error:

```
	err := Forge("valid@email.com").
		AllowEmpty().
		Email(expected).
		Assert()
```

Get all errors:

```
	err := Forge("valid@email.com").
		AllowEmpty().
		Email(expected).
		Errors()
```
