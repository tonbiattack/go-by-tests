package contextflow

import "context"

type userIDKey struct{}

func StringKeyCollision() string {
	ctx := context.WithValue(context.Background(), "user-id", "application")
	ctx = context.WithValue(ctx, "user-id", "middleware")
	return ctx.Value("user-id").(string)
}

func PrivateKeyDoesNotCollide() (string, string) {
	ctx := context.WithValue(context.Background(), "user-id", "application")
	ctx = context.WithValue(ctx, userIDKey{}, "middleware")
	return ctx.Value("user-id").(string), ctx.Value(userIDKey{}).(string)
}
