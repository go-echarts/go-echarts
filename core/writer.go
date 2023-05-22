package core

import (
	"os"
)

type DefaultFileWriter struct {
	File string
}

func (dw *DefaultFileWriter) Write(data []byte) {
	writer, err := os.Create(dw.File)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}
