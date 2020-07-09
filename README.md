# valida

[![Build Status](https://travis-ci.org/martinusso/zx.svg?branch=master)](https://travis-ci.org/martinusso/zx)
[![Coverage Status](https://coveralls.io/repos/github/martinusso/zx/badge.svg?branch=master)](https://coveralls.io/github/martinusso/zx?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinusso/zx)](https://goreportcard.com/report/github.com/martinusso/zx)

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
