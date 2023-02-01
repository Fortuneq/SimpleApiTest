package domain

import "context"

type UserService interface {
	Create(ctx context.Context, p dto.CreateUser) error
}
