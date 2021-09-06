package keycloak

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

// create a new group.
func createGroup(t *testing.T, k *Keycloak, realm string, groupName string) string {
	t.Helper()

	group := &Group{
		Name: String(groupName),
	}

	res, err := k.Groups.Create(context.Background(), realm, group)
	if err != nil {
		t.Errorf("Groups.Create returned error: %v", err)
	}

	parts := strings.Split(res.Header.Get("Location"), "/")
	groupID := parts[len(parts)-1]
	return groupID
}

func TestGroupsService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	ctx := context.Background()

	group := &Group{
		Name: String("mygroup"),
	}

	res, err := k.Groups.Create(ctx, realm, group)
	if err != nil {
		t.Errorf("Groups.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestGroupsService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createGroup(t, k, realm, "group_a")
	createGroup(t, k, realm, "group_b")

	groups, res, err := k.Groups.List(context.Background(), realm)
	if err != nil {
		t.Errorf("Groups.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if len(groups) != 2 {
		t.Errorf("got: %d, want: %d", len(groups), 2)
	}
}

func TestGroupsService_Get(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	groupID := createGroup(t, k, realm, "group")

	group, res, err := k.Groups.Get(context.Background(), realm, groupID)
	if err != nil {
		t.Errorf("Groups.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *group.Name != "group" {
		t.Errorf("got: %s, want: %s", *group.Name, "group")
	}
}

func TestGroupsService_Delete(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	groupID := createGroup(t, k, realm, "group")

	res, err := k.Groups.Delete(context.Background(), realm, groupID)
	if err != nil {
		t.Errorf("Groups.Delete returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestGroupsService_AddRealmRoles(t *testing.T) {
	k := client(t)

	realm := "first"
	roleNames := [2]string{"role1", "role2"}

	createRealm(t, k, realm)
	for _, r := range roleNames {
		createRealmRole(t, k, realm, r)
	}
	groupID := createGroup(t, k, realm, "group")

	ctx := context.Background()

	// get role in order to assign it
	var roles []*Role
	for _, r := range roleNames {
		role, _, err := k.RealmRoles.GetByName(ctx, realm, r)
		if err != nil {
			t.Errorf("RealmRoles.GetByName returned error: %v", err)
		}
		roles = append(roles, role)
	}

	res, err := k.Groups.AddRealmRoles(ctx, realm, groupID, roles)
	if err != nil {
		t.Errorf("Groups.AddRealmRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestGroupsService_RemoveRealmRoles(t *testing.T) {
	k := client(t)

	realm := "first"
	roleNames := [2]string{"role1", "role2"}

	createRealm(t, k, realm)
	for _, r := range roleNames {
		createRealmRole(t, k, realm, r)
	}
	groupID := createGroup(t, k, realm, "group")

	ctx := context.Background()

	// get role in order to assign it
	var roles []*Role
	for _, r := range roleNames {
		role, _, err := k.RealmRoles.GetByName(ctx, realm, r)
		if err != nil {
			t.Errorf("RealmRoles.GetByName returned error: %v", err)
		}
		roles = append(roles, role)
	}

	_, err := k.Groups.AddRealmRoles(ctx, realm, groupID, roles)
	if err != nil {
		t.Errorf("Groups.AddRealmRoles returned error: %v", err)
	}

	res, err := k.Groups.RemoveRealmRoles(ctx, realm, groupID, roles)
	if err != nil {
		t.Errorf("Groups.RemoveRealmRoles returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}
