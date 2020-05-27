package account


type User struct{
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	City 	string 	`json:city`
	Phone  	int 	`json:phone`
	Sal 	int 	`json:sal`
}

type Repository interface{
	CreateUSer(user User)error
	GetUser(id)(string, error)
}