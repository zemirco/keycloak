package keycloak

import (
	"context"
	"net/http"
	"testing"
)

// create a new realm role.
func createRealmRole(t *testing.T, k *Keycloak, realm string, name string) {
	t.Helper()

	role := &Role{
		Name:        String(name),
		Description: String(name + " description"),
	}

	if _, err := k.RealmRoles.Create(context.Background(), realm, role); err != nil {
		t.Errorf("RealmRoles.Create returned error: %v", err)
	}
}

func TestRealmRolesService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	ctx := context.Background()

	role := &Role{
		Name:        String("my name"),
		Description: String("my description"),
	}

	res, err := k.RealmRoles.Create(ctx, realm, role)
	if err != nil {
		t.Errorf("RealmRoles.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestRealmRolesService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createRealmRole(t, k, realm, "first")
	createRealmRole(t, k, realm, "second")

	roles, res, err := k.RealmRoles.List(context.Background(), realm, nil)
	if err != nil {
		t.Errorf("RealmRoles.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the "offline_access" and "uma_authorization" roles
	if len(roles) != 4 {
		t.Errorf("got: %d, want: %d", len(roles), 4)
	}
}

func TestRealmRolesService_GetByName(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createRealmRole(t, k, realm, "first")

	role, res, err := k.RealmRoles.GetByName(context.Background(), realm, "first")
	if err != nil {
		t.Errorf("RealmRoles.GetByName returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *role.Name != "first" {
		t.Errorf("got: %s, want: %s", *role.Name, "first")
	}
}

func TestRealmRolesService_GetByID(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	createRealmRole(t, k, realm, "first")

	// get by name first to get id
	role, res, err := k.RealmRoles.GetByName(context.Background(), realm, "first")
	if err != nil {
		t.Errorf("RealmRoles.GetByName returned error: %v", err)
	}

	// now get by id
	role, res, err = k.RealmRoles.GetByID(context.Background(), realm, *role.ID)
	if err != nil {
		t.Errorf("RealmRoles.GetByID returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *role.Name != "first" {
		t.Errorf("got: %s, want: %s", *role.Name, "first")
	}
}
