package keycloak

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

// create a new keycloak instance.
func client(t *testing.T) *Keycloak {
	t.Helper()

	config := oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/auth/realms/master/protocol/openid-connect/token",
		},
	}

	ctx := context.Background()

	token, err := config.PasswordCredentialsToken(ctx, "admin", "admin")
	if err != nil {
		t.Error(err)
	}

	client := config.Client(ctx, token)

	return NewKeycloak(client)
}

func TestKeycloak_New(t *testing.T) {

}

func TestKeycloak_NewRequest(t *testing.T) {

}

func TestKeycloak_Do(t *testing.T) {

}
