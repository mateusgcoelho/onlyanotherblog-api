package utils

import (
	"github.com/oklog/ulid/v2"
)

func GenerateUniqueId() string {
	id := ulid.Make()
	return id.String()
}
