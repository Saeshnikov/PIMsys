package storage

type User struct {
	ID       int
	Email    string
	PassHash []byte
	IsAdmin  bool
}
