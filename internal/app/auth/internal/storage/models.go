package storage

type App struct {
	ID     int
	Name   string
	Secret string
}

type User struct {
	ID       int64
	Email    string
	PassHash []byte
}
