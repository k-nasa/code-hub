package model

import (
	"context"
	"testing"
)

func TestUserContextValue(t *testing.T) {
	ctx := context.Background()
	u := &User{
		"firebaseUID",
		"DisplayName",
		"Email",
		"PhotoURL",
	}
	ctx = SetUserToContext(ctx, u)
	getu, err := GetUserFromContext(ctx)
	if err != nil {
		t.Fatalf("%d: invalid Get User ContextValue", err)
	}
	if u != getu {
		t.Fatalf("%d: invalid Get User ContextValue ,expected %v, got %v", err, u, getu)
	}
}
