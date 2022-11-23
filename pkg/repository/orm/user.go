package orm

import (
    "context"
    "fmt"
    "github.com/afv1/sims-tree/ent"
    "github.com/afv1/sims-tree/ent/user"
    "github.com/afv1/sims-tree/pkg/dto"
    "github.com/afv1/sims-tree/pkg/repository"
    "strconv"
    "time"
)

func (o *ORM) CreateUser(ctx context.Context, data dto.AuthUser) (*ent.User, error) {
    client := repository.GetClient()
    defer repository.CloseClient(client)

    u, err := client.User.
        Create().
        SetLogin(data.Login).
        SetPassword(data.Password).
        SetComment("").
        SetCreatedAt(time.Now()).
        Save(ctx)

    // TODO: Replace error output with presenter data instead of Ent errors
    if err != nil {
        return nil, fmt.Errorf("could not create user %s: %s", data.Login, err)
    }

    return u, nil
}

func (o *ORM) QueryUser(ctx context.Context, data dto.AuthUser) (*ent.User, error) {
    client := repository.GetClient()
    defer repository.CloseClient(client)

    u, err := client.User.
        Query().
        Where(user.Login(data.Login)).
        Only(ctx)

    if err != nil {
        return nil, fmt.Errorf("could not query user %s:%s", data.Login, err)
    }

    return u, nil
}

func (o *ORM) EditUser(ctx context.Context, data dto.AuthUser) (*ent.User, error) {
    client := repository.GetClient()
    defer repository.CloseClient(client)

    userId, _ := strconv.Atoi("1")

    u, err := client.User.
        UpdateOneID(userId).
//        SetComment(data.Comment).
        SetPassword(data.Password).
        Save(ctx)

    if err != nil {
        return nil, fmt.Errorf("could not update user %s:%s", data.Login, err)
    }

    return u, nil
}

func (o *ORM) DeleteUser(ctx context.Context, data dto.AuthUser) error {
    client := repository.GetClient()
    defer repository.CloseClient(client)

    userId, _ := strconv.Atoi("1")

    err := client.User.
        DeleteOneID(userId).
        Exec(ctx)

    if err != nil {
        return fmt.Errorf("could not update user %s:%s", data.Login, err)
    }

    return nil
}
