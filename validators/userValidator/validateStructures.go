package userValidators

type Address struct {
	Prefecture string `json:"prefecture" validate:"required"`
	City       string `json:"city" validate:"required"`
	District   string `json:"district" validate:"required"`
	Street     string `json:"street" validate:"required"`
	Additional string `json:"additional"`
}

type CreateUser struct {
	Username  string     `json:"username" validate:"required"`
	Password  string     `json:"password" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Role      string     `json:"role"`
	Addresses []*Address `json:"addresses" validate:"required,dive,required"`
}

type UpdateUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email" validate:"email"`
	Role      string `json:"role"`
	Addresses []*Address
}

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
