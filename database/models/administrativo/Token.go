package administrativo

import "github.com/google/uuid"

type Token struct {
	Token    string
	AccessID uuid.UUID
}
