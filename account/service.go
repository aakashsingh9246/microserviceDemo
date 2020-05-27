package account


type Service interface {
	CreateUser(id,phone,sal int, city,name string)(string, error)
	GetUser(id int)(string, error)
}