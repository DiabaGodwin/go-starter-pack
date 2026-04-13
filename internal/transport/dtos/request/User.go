package request

type ListUserRequest struct {
	Limit  int32
	Offset int32
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
