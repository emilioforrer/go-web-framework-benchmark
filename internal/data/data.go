package data

import _ "embed"

//go:embed data.json
var fileData []byte

func Get() []byte {
	return fileData
}
