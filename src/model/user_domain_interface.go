package model

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int8
	GetId() string

	SetId(ID string)

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
