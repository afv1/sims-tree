package repository

import (
    "context"
    "github.com/afv1/sims-tree/ent"
    "github.com/afv1/sims-tree/pkg/dto"
)

type ORM interface {
    CreateUser(ctx context.Context, data dto.AuthUser) (*ent.User, error)
    AuthUser(ctx context.Context, data dto.AuthUser) (*ent.User, error)
    QueryUser(ctx context.Context, data dto.AuthUser) (*ent.User, error)
    EditUser(ctx context.Context, data dto.AuthUser) (*ent.User, error)
    DeleteUser(ctx context.Context, data dto.AuthUser) error
}
