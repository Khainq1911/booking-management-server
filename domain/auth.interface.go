package domain

import (
	"context"
	"manage-system-server/model"
)

type AuthStorage interface {
	CheckLogin(ctx context.Context, username string) (model.Employee, error)
}
