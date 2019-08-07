package httputil

import (
	"context"
	"github.com/voyagegroup/treasure-app/domain/model"
	"testing"
)

func TestUserContextValue(t *testing.T) {
	ctx := context.Background()
	u := &model.User{
		"firebaseUID",
		"DisplayName",
		"Email",
		"PhotoURL",
	}
	ctx = SetUserToContext(ctx, u)
	getu, err := GetUserFromContext(ctx)
	if err != nil {
		t.Fatalf("%d: invalid Show User ContextValue", err)
	}
	if u != getu {
		t.Fatalf("%d: invalid Show User ContextValue ,expected %v, got %v", err, u, getu)
	}
}
