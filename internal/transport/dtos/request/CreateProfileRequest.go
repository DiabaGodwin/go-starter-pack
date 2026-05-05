package request

type CreateProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	AvatarUrl string `json:"avatar_url"`
	Phone     string `json:"phone"`
}
