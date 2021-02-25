package keycloak

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

// create a new client scope.
func createClientScope(t *testing.T, k *Keycloak, realm string, clientScopeName string) string {
	t.Helper()

	clientScope := &ClientScope{
		Name:        String(clientScopeName),
		Description: String(clientScopeName + " description"),
	}

	res, err := k.ClientScopes.Create(context.Background(), realm, clientScope)
	if err != nil {
		t.Errorf("ClientScopes.Create returned error: %v", err)
	}

	parts := strings.Split(res.Header.Get("Location"), "/")
	clientScopeID := parts[len(parts)-1]
	return clientScopeID
}

func TestClientScopesService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientScopes, res, err := k.ClientScopes.List(context.Background(), realm)
	if err != nil {
		t.Errorf("ClientScopes.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// address, email, microprofile-jwt, offline_access, phone, profile, role_list, roles, web-origins
	if len(clientScopes) != 9 {
		t.Errorf("got: %d, want: %d", len(clientScopes), 9)
	}
}

func TestClientScopesService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	ctx := context.Background()

	clientScope := &ClientScope{
		Name:        String("my-client-scope"),
		Description: String("some description"),
	}

	res, err := k.ClientScopes.Create(ctx, realm, clientScope)
	if err != nil {
		t.Errorf("ClientScopes.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestClientScopesService_Get(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientScopeID := createClientScope(t, k, realm, "my-client-scope")

	clientScopes, res, err := k.ClientScopes.Get(context.Background(), realm, clientScopeID)
	if err != nil {
		t.Errorf("ClientScopes.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *clientScopes.Name != "my-client-scope" {
		t.Errorf("got: %s, want: %s", *clientScopes.Name, "my-client-scope")
	}
}

func TestClientScopesService_Delete(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientScopeID := createClientScope(t, k, realm, "my-client-scope")

	res, err := k.ClientScopes.Delete(context.Background(), realm, clientScopeID)
	if err != nil {
		t.Errorf("ClientScopes.Delete returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}
