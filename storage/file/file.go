package file

import (
	"os"
)

type File struct {
	filename   string
	descriptor *os.File
}

func NewFile(f string) *File {
	return &File{filename: f}
}

func (f *File) Open() {
	d, err := os.Create(f.filename)
	if err != nil {
		panic(err)
	}
	f.descriptor = d
}

func (f *File) Close() {
	err := f.descriptor.Close()
	if err != nil {
		panic(err)
	}
}

func (f *File) Append(str string) {
	f.descriptor.WriteString(str)
}
