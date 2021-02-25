package keycloak

import (
	"context"
	"net/http"
	"testing"
)

func createScope(t *testing.T, k *Keycloak, realm, clientID, name string) *Scope {
	t.Helper()

	scope := &Scope{
		Name: String(name),
	}

	scope, _, err := k.Scopes.Create(context.Background(), realm, clientID, scope)
	if err != nil {
		t.Errorf("Scopes.Create returned error: %v", err)
	}
	return scope
}

func TestScopesService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	scope := &Scope{
		Name: String("scope_read"),
	}

	scope, res, err := k.Scopes.Create(context.Background(), realm, clientID, scope)
	if err != nil {
		t.Errorf("Scopes.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *scope.Name != "scope_read" {
		t.Errorf("got: %s, want: %s", *scope.Name, "scope_read")
	}
}

func TestScopesService_List(t *testing.T) {
	k := client(t)
	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	createScope(t, k, realm, clientID, "read")
	createScope(t, k, realm, clientID, "write")

	scopes, res, err := k.Scopes.List(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Scopes.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if len(scopes) != 2 {
		t.Errorf("got: %d, want: %d", len(scopes), 2)
	}
}

func TestScopesService_Get(t *testing.T) {
	k := client(t)
	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	scope := createScope(t, k, realm, clientID, "read")

	scope, res, err := k.Scopes.Get(context.Background(), realm, clientID, *scope.ID)
	if err != nil {
		t.Errorf("Scopes.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *scope.Name != "read" {
		t.Errorf("got: %s, want: %s", *scope.Name, "read")
	}
}

func TestScopesService_Delete(t *testing.T) {
	k := client(t)
	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	scope := createScope(t, k, realm, clientID, "read")

	res, err := k.Scopes.Delete(context.Background(), realm, clientID, *scope.ID)
	if err != nil {
		t.Errorf("Scopes.Delete returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestScopesService_Update(t *testing.T) {
	k := client(t)
	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	scope := createScope(t, k, realm, clientID, "read")

	scope.Name = String("updated")

	res, err := k.Scopes.Update(context.Background(), realm, clientID, scope)
	if err != nil {
		t.Errorf("Scopes.Update returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}
