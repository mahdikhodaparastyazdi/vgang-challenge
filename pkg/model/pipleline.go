package model

import "context"

// Runner Represents a pipeline of stages that are run as a unit.
type Runner interface {
	Run(ctx context.Context) error
}
