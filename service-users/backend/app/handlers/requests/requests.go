package requests

type RegisterRequest struct {
	Name             string `validate:"required,max=128,alphanumunicode"`
	Surname          string `validate:"required,max=128,alphanumunicode"`
	Age              int    `validate:"required,numeric,max=99,min=10"`
	Interests        string `validate:"required,max=768"`
	City             string `validate:"required"`
	Sex              int    `validate:"required,numeric"`
	Email            string `validate:"required,email,max=128"`
	Password         string `validate:"required,eqfield=Confirm_password"`
	Confirm_password string `validate:"required"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserRelationRequest struct {
	Friend_user_id int `validate:"required,numeric"`
}
