package file

import (
	"fmt"
	"github.com/mstovicek/go-chapter-1-google-places/entity"
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

func (f File) Append(readableEntity entity.Readable) {
	f.descriptor.WriteString(fmt.Sprintln(readableEntity.ToString()))
}
