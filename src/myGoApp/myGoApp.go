package myGoApp

type User struct {
	Email string
	Password string
}


type Loginer interface {
	Login()
}




