package file

import (
	"fmt"
	"github.com/mstovicek/go-chapter-1-google-places/entity"
	"os"
)

type file struct {
	descriptor *os.File
}

func Open(filename string) *file {
	descriptor, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return &file{descriptor}
}

func (f file) Close() {
	err := f.descriptor.Close()
	if err != nil {
		panic(err)
	}
}

func (f file) Append(readableEntity entity.Readable) {
	f.descriptor.WriteString(fmt.Sprintln(readableEntity.ToString()))
}
