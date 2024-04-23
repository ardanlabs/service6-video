// Package mid provides app level middleware support.
package mid

import "context"

// Handler represents the handler function that needs to be called.
type Handler func(context.Context) error
