package models

type SignUp struct {
	FirstName string `json: "firstName" validate:"required"`
	LastName  string `json: "lastName" validate:"required"`
	Password  string `json: "password" validate:"required"`
	Email     string `json: "email" validate:"required,email"`
}

// func (p *SignUp) Validate() error {
// 	validate := validator.New() //creating validator
// 	validate.RegisterValidation("email", validateSKU)
// 	return validate.Struct(p) // passing struct to validate
// }
