package requests

type RegisterRequest struct {
	Name            string `validate:"required"`
	Email           string `validate:"required,email,max=128"`
	Password        string `validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `validate:"required"`
}
