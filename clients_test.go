package keycloak

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

// create a new client.
func createClient(t *testing.T, k *Keycloak, realm string, clientID string) string {
	t.Helper()

	client := &Client{
		Enabled:                      Bool(true),
		ClientID:                     String(clientID),
		RedirectUris:                 []string{"http://localhost:4200/*"},
		AuthorizationServicesEnabled: Bool(true),
		ServiceAccountsEnabled:       Bool(true),
		PublicClient:                 Bool(false),
	}

	res, err := k.Clients.Create(context.Background(), realm, client)
	if err != nil {
		t.Errorf("Clients.Create returned error: %v", err)
	}

	parts := strings.Split(res.Header.Get("Location"), "/")
	id := parts[len(parts)-1]
	return id
}

func TestClientsService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	ctx := context.Background()

	client := &Client{
		Enabled:  Bool(true),
		ClientID: String("myclient"),
	}

	res, err := k.Clients.Create(ctx, realm, client)
	if err != nil {
		t.Errorf("Clients.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}
}

func TestClientsService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clients, res, err := k.Clients.List(context.Background(), realm)
	if err != nil {
		t.Errorf("Clients.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the "account", "account-console", "admin-cli", "broker", "realm-management", "security-admin-console"
	if len(clients) != 6 {
		t.Errorf("got: %d, want: %d", len(clients), 6)
	}
}

func TestClientsService_Get(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	client, res, err := k.Clients.Get(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Clients.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *client.ClientID != "client" {
		t.Errorf("got: %s, want: %s", *client.ClientID, "client")
	}
}

func TestClientsService_GetSecret(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	credential, res, err := k.Clients.GetSecret(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Clients.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *credential.Type != "secret" {
		t.Errorf("got: %s, want: %s", *credential.Type, "secret")
	}
}
