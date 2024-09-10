package model

type userDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) SetId(id string) {
	ud.Id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) GetId() string {
	return ud.Id
}
