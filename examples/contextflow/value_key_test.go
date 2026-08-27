package contextflow

import "testing"

func TestContextWithValueで文字列キーを共有すると値が衝突する(t *testing.T) {
	if got := StringKeyCollision(); got != "middleware" {
		t.Fatalf("StringKeyCollision() = %q, want middleware", got)
	}
	application, middleware := PrivateKeyDoesNotCollide()
	if application != "application" || middleware != "middleware" {
		t.Fatalf("PrivateKeyDoesNotCollide() = %q, %q", application, middleware)
	}
}
