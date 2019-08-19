package httputil

import (
	"context"

	"github.com/k-nasa/code-hub/model"
	"github.com/pkg/errors"
)

type contextKey string

const UserContextKey contextKey = "user"

func SetUserToContext(ctx context.Context, u *model.User) context.Context {
	return context.WithValue(ctx, UserContextKey, u)
}

func GetUserFromContext(ctx context.Context) (*model.User, error) {
	v := ctx.Value(UserContextKey)
	user, ok := v.(*model.User)
	if !ok {
		return nil, errors.New("user not found from context value")
	}
	return user, nil
}
