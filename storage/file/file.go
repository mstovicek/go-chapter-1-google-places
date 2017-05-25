package file

import (
	"os"
)

type File struct {
	descriptor *os.File
}

func Open(filename string) *File {
	descriptor, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return &File{descriptor}
}

func (f File) Close() {
	err := f.descriptor.Close()
	if err != nil {
		panic(err)
	}
}

func (f File) Append(str string) {
	f.descriptor.WriteString(str)
}
