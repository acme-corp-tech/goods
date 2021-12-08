package service

import (
	"github.com/acme-corp-tech/goods/internal/domain/greeting"
)

// GreetingMakerProvider is a service provider.
type GreetingMakerProvider interface {
	GreetingMaker() greeting.Maker
}
