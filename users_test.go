package keycloak

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

// create a new user.
func createUser(t *testing.T, k *Keycloak, realm string, username string) string {
	t.Helper()

	user := &User{
		Enabled:   Bool(true),
		Username:  String(username),
		Email:     String(username + "@email.com"),
		FirstName: String("first"),
		LastName:  String("last"),
	}

	res, err := k.Users.Create(context.Background(), realm, user)
	if err != nil {
		t.Errorf("Users.Create returned error: %v", err)
	}

	parts := strings.Split(res.Header.Get("Location"), "/")
	userID := parts[len(parts)-1]
	return userID
}

func TestUsersService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	ctx := context.Background()

	user := &User{
		Enabled:   Bool(true),
		Username:  String("username"),
		Email:     String("user@email.com"),
		FirstName: String("first"),
		LastName:  String("last"),
	}

	res, err := k.Users.Create(ctx, realm, user)
	if err != nil {
		t.Errorf("Users.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestUsersService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createUser(t, k, realm, "john")
	createUser(t, k, realm, "mark")

	users, res, err := k.Users.List(context.Background(), realm)
	if err != nil {
		t.Errorf("Users.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the master realm
	if len(users) != 2 {
		t.Errorf("got: %d, want: %d", len(users), 2)
	}
}

func TestUsersService_ResetPassword(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	userID := createUser(t, k, realm, "john")

	ctx := context.Background()

	credential := &Credential{
		Type:      String("password"),
		Value:     String("mypassword"),
		Temporary: Bool(false),
	}
	res, err := k.Users.ResetPassword(ctx, realm, userID, credential)
	if err != nil {
		t.Errorf("Users.ResetPassword returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_JoinGroup(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	userID := createUser(t, k, realm, "john")
	groupID := createGroup(t, k, realm, "group")

	ctx := context.Background()

	res, err := k.Users.JoinGroup(ctx, realm, userID, groupID)
	if err != nil {
		t.Errorf("Users.JoinGroup returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_LeaveGroup(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	userID := createUser(t, k, realm, "john")
	groupID := createGroup(t, k, realm, "group")

	ctx := context.Background()

	res, err := k.Users.JoinGroup(ctx, realm, userID, groupID)
	if err != nil {
		t.Errorf("Users.JoinGroup returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}

	res, err = k.Users.LeaveGroup(ctx, realm, userID, groupID)
	if err != nil {
		t.Errorf("Users.LeaveGroup returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_AddRealmRoles(t *testing.T) {
	k := client(t)

	realm := "first"
	roleName := "role"

	createRealm(t, k, realm)
	createRealmRole(t, k, realm, roleName)
	userID := createUser(t, k, realm, "user")

	ctx := context.Background()

	// get role in order to assign it
	role, _, err := k.RealmRoles.GetByName(ctx, realm, roleName)
	if err != nil {
		t.Errorf("RealmRoles.GetByName returned error: %v", err)
	}

	roles := []*Role{role}

	res, err := k.Users.AddRealmRoles(ctx, realm, userID, roles)
	if err != nil {
		t.Errorf("Users.AddRealmRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_RemoveRealmRoles(t *testing.T) {
	k := client(t)

	realm := "first"
	roleName := "role"

	createRealm(t, k, realm)
	createRealmRole(t, k, realm, roleName)
	userID := createUser(t, k, realm, "user")

	ctx := context.Background()

	// get role in order to assign it
	role, _, err := k.RealmRoles.GetByName(ctx, realm, roleName)
	if err != nil {
		t.Errorf("RealmRoles.GetByName returned error: %v", err)
	}

	roles := []*Role{role}

	if _, err := k.Users.AddRealmRoles(ctx, realm, userID, roles); err != nil {
		t.Errorf("Users.AddRealmRoles returned error: %v", err)
	}

	res, err := k.Users.RemoveRealmRoles(ctx, realm, userID, roles)
	if err != nil {
		t.Errorf("Users.RemoveRealmRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_AddClientRoles(t *testing.T) {

}

func TestUsersService_RemoveClientRoles(t *testing.T) {

}
