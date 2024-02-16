package option

import "testing"

type User struct {
	ID int
}

func getUser(id int) Option[User] {
	if id == 0 {
		return None()
	}
	return Some(User{id})
}

func TestOption(t *testing.T) {
	u := getUser(0)
	if u.IsNone() {
		t.Log("User is None")
	} else {
		t.Errorf("User is not None: %v", u)
	}

	u = getUser(1)
	if u.IsSome() {
		t.Logf("User is Some: %v", u)
	} else {
		t.Errorf("User is not Some")
	}
}

func TestUnwrap(t *testing.T) {
	u := getUser(1)
	user := Unwrap[User](u)
	if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", user.ID)
	}
}

func TestUnwrapOr(t *testing.T) {
	u := getUser(0)
	user := UnwrapOr[User](u, User{ID: 1})
	if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", user.ID)
	}

	u = getUser(2)
	user = UnwrapOr[User](u, User{ID: 1})
	if user.ID != 2 {
		t.Errorf("Expected user ID to be 2, got %d", user.ID)
	}
}
