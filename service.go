package sprintsd

import "context"

type Service interface {
	Enroll(context.Context, *Registration) error
	Locate(context.Context, string) (*Registration, error)
	Forget(context.Context, string) error
}
