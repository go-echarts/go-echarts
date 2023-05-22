package core

import (
	"os"
)

type DefaultWriter struct {
}

func (dw *DefaultWriter) Write(data []byte, dest string) {
	writer, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}
