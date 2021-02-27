package keycloak

import (
	"context"
	"net/http"
	"testing"
)

// create a new client role
func createClientRole(t *testing.T, k *Keycloak, realm, clientID, roleName string) {
	t.Helper()

	role := &Role{
		Name:        String(roleName),
		Description: String(roleName + " description"),
	}

	if _, err := k.ClientRoles.Create(context.Background(), realm, clientID, role); err != nil {
		t.Errorf("ClientRoles.Create returned error: %v", err)
	}
}

func TestClientRolesService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	ctx := context.Background()

	role := &Role{
		Name:        String("role"),
		Description: String("description"),
	}

	res, err := k.ClientRoles.Create(ctx, realm, clientID, role)
	if err != nil {
		t.Errorf("ClientRoles.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestClientRolesService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	createClientRole(t, k, realm, clientID, "first")
	createClientRole(t, k, realm, clientID, "second")

	roles, res, err := k.ClientRoles.List(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("ClientRoles.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// includes "uma_protection"
	if len(roles) != 3 {
		t.Errorf("got: %d, want: %d", len(roles), 3)
	}
}
