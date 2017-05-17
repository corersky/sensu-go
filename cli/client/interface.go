package client

import "github.com/sensu/sensu-go/types"

// APIClient client methods across the Sensu API
type APIClient interface {
	EventAPIClient
	CheckAPIClient
	HandlerAPIClient
	EntityAPIClient
	UserAPIClient
}

// EventAPIClient client methods for events
type EventAPIClient interface {
	ListEvents() ([]types.Event, error)
}

// CheckAPIClient client methods for checks
type CheckAPIClient interface {
	ListChecks() ([]types.Check, error)
	CreateCheck(*types.Check) error
	DeleteCheck(*types.Check) error
}

// HandlerAPIClient client methods for checks
type HandlerAPIClient interface {
	ListHandlers() ([]types.Handler, error)
	CreateHandler(*types.Handler) error
	DeleteHandler(*types.Handler) error
}

// EntityAPIClient client methods for checks
type EntityAPIClient interface {
	ListEntities() ([]types.Entity, error)
	FetchEntity(ID string) (types.Entity, error)
}

// CheckAPIClient client methods for checks
type UserAPIClient interface {
	CreateUser(*types.User) error
}
