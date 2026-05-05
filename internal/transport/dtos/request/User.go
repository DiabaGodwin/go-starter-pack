package request

type ListUserRequest struct {
	Limit  int32
	Offset int32
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4"`
	FirstName string `json:"firstname" validate:"required,max=20"`
	LastName  string `json:"lastname" validate:"required,max=20"`
	Role      string `json:"role" validate:"required"`
}
