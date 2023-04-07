package keycloak

import (
	"context"
	"net/http"
	"reflect"
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
		Attributes: &map[string][]string{
			"some_key": {"some_value"},
		},
	}

	res, err := k.Users.Create(ctx, realm, user)
	if err != nil {
		t.Errorf("Users.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestUsersService_GetByID(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	id := createUser(t, k, realm, "newuser")

	ctx := context.Background()
	user, res, err := k.Users.GetByID(ctx, realm, id)
	if err != nil {
		t.Errorf("Users.GetByID returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}
	if *user.Username != "newuser" {
		t.Errorf("got: %s, want: %s", *user.Username, "newuser")
	}
}

func TestUsersService_Update(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	id := createUser(t, k, realm, "newuser")

	ctx := context.Background()
	user, res, err := k.Users.GetByID(ctx, realm, id)
	if err != nil {
		t.Errorf("Users.GetByID returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	newAttributes := &map[string][]string{
		"some_key": {"other_value"},
		"new_key":  {"some_value"},
	}
	user.Attributes = newAttributes

	res, err = k.Users.Update(ctx, realm, user)
	if err != nil {
		t.Errorf("Users.Update returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}

	user, res, err = k.Users.GetByID(ctx, realm, id)
	if err != nil {
		t.Errorf("Users.GetByID returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(user.Attributes, newAttributes) {
		t.Errorf("Users.Update attributes do not match: %v != %v", user.Attributes, newAttributes)
	}
}

func TestUsersService_Delete(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	userID := createUser(t, k, realm, "john")

	res, err := k.Users.Delete(context.Background(), realm, userID)
	if err != nil {
		t.Errorf("Users.Delete returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
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

func TestUsersService_ListGroups(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	id := createUser(t, k, realm, "freddy")
	groupID := createGroup(t, k, realm, "group_of_freddy")

	res, err := k.Users.JoinGroup(context.Background(), realm, id, groupID)

	groups, res, err := k.Users.ListUserGroups(context.Background(), realm, id)
	if err != nil {
		t.Errorf("Users.ListUserGroups returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the master realm
	if len(groups) != 1 {
		t.Errorf("got: %d, want: %d", len(groups), 1)
	}

	if *groups[0].Name != "group_of_freddy" {
		t.Errorf("got: %v, want: %v", *groups[0].Name, "group_of_freddy")
	}
}

func TestUsersService_GetByUsername(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createUser(t, k, realm, "john")

	users, res, err := k.Users.GetByUsername(context.Background(), realm, "john")
	if err != nil {
		t.Errorf("Users.GetByUsername returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the master realm
	if len(users) != 1 {
		t.Errorf("got: %d, want: %d", len(users), 1)
	}

	if *users[0].Username != "john" {
		t.Errorf("got: %s, want: %s", *users[0].Username, "john")
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

func TestUsersService_ListRealmRoles(t *testing.T) {
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

	roles, res, err := k.Users.ListRealmRoles(ctx, realm, userID)
	if err != nil {
		t.Errorf("Users.ListRealmRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// "role" and "default-roles-first"
	if len(roles) != 2 {
		t.Errorf("got: %d, want: %d", len(roles), 2)
	}
}

func TestUsersService_AddClientRoles(t *testing.T) {
	k := client(t)

	realm := "first"

	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	createClientRole(t, k, realm, clientID, "role")
	userID := createUser(t, k, realm, "user")

	ctx := context.Background()

	// get role in order to assign it
	role, _, err := k.ClientRoles.Get(ctx, realm, clientID, "role")
	if err != nil {
		t.Errorf("ClientRoles.Get returned error: %v", err)
	}

	roles := []*Role{role}

	res, err := k.Users.AddClientRoles(ctx, realm, userID, clientID, roles)
	if err != nil {
		t.Errorf("Users.AddClientRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_RemoveClientRoles(t *testing.T) {

}

func TestUsersService_SendVerifyEmail(t *testing.T) {
	k := client(t)

	realm := "first"

	createRealm(t, k, realm)
	userID := createUser(t, k, realm, "user")

	res, err := k.Users.SendVerifyEmail(context.Background(), realm, userID, nil)
	if err != nil {
		t.Errorf("Users.SendVerifyEmail returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestUsersService_ExecuteActionsEmail(t *testing.T) {
	k := client(t)

	realm := "first"

	createRealm(t, k, realm)
	userID := createUser(t, k, realm, "user")

	opts := &ExecuteActionsEmailOptions{
		Lifespan: 1000,
	}

	res, err := k.Users.ExecuteActionsEmail(context.Background(), realm, userID, opts, []string{"UPDATE_PROFILE"})
	if err != nil {
		t.Errorf("Users.ExecuteActionsEmail returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}
