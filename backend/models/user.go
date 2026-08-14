package models

type User struct {
	ID              uint
	Name            string
	Email           string
	Password_digest string
	Description     *string
	Icon_url        string
}
