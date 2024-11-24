package storage

type User struct {
	ID       int
	Name     string
	Email    string
	Phone    string
	PassHash []byte
	IsAdmin  bool
}
