//go:generate go run github.com/99designs/gqlgen -v

package graph

import (
	"github.com/mojeico/gqlgen-golang/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    service.UserService
	MeetupsService service.MeetupsService
}
