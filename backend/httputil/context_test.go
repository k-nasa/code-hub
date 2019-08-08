package httputil

import (
	"context"
	"testing"

	"github.com/voyagegroup/treasure-app/model"
)

func TestUserContextValue(t *testing.T) {
	ctx := context.Background()
	u := &model.User{
		FirebaseUID: "firebaseUID",
		DisplayName: "DisplayName",
		Email:       "Email",
		PhotoURL:    "PhotoURL",
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
