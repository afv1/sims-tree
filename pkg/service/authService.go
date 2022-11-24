package service

import (
    "context"
    "github.com/afv1/sims-tree/ent"
    "github.com/afv1/sims-tree/pkg/dto"
    "github.com/afv1/sims-tree/pkg/repository"
    "github.com/afv1/sims-tree/pkg/repository/orm"
)

type AuthService struct {
    ORM repository.ORM
}

func newAuthService() *AuthService {
    return &AuthService{ORM: orm.NewORM()}
}

func (as *AuthService) Register(data dto.AuthUser) (*ent.User, error) {
    data.Password = generatePasswordHash(data.Password)

    user, err := as.ORM.CreateUser(context.Background(), data)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (as *AuthService) Login(data dto.AuthUser) (*ent.User, error) {
    data.Password = generatePasswordHash(data.Password)

    user, err := as.ORM.AuthUser(context.Background(), data)

    if err != nil {
        return nil, err
    }

    return user, err
}
