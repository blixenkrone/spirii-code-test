// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package example

import (
	"github.com/google/uuid"
)

type Foo struct {
	ID    uuid.UUID
	Value string
}
