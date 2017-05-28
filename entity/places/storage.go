package places

type Storage interface {
	Open()
	Close()
	Append(str string)
}
