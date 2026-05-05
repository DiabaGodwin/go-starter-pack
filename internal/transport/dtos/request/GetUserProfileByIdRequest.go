package request

import "github.com/google/uuid"

type GetUserProfileByIdRequest struct {
	Id uuid.UUID
}
