package docs

import (
	_ "embed"
)

//go:embed swagger.json
var SwaggerByte []byte
