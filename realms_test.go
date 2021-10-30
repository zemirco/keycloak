package keycloak

import (
	"context"
	"net/http"
	"testing"
)

// create a new realm and delete it afterwards.
func createRealm(t *testing.T, k *Keycloak, name string) {
	t.Helper()

	realm := &Realm{
		Enabled: Bool(true),
		ID:      String(name),
		Realm:   String(name),
		SMTPServer: &map[string]string{
			"from": "john@wayne.com",
			"host": "mailhog",
			"port": "1025",
		},
	}

	ctx := context.Background()

	// create a new realm
	if _, err := k.Realms.Create(ctx, realm); err != nil {
		t.Errorf("Realms.Create returned error: %v", err)
	}

	t.Cleanup(func() {
		if _, err := k.Realms.Delete(ctx, name); err != nil {
			t.Errorf("Realms.Delete returned error: %v", err)
		}
	})
}

func TestRealmsService_Create(t *testing.T) {
	k := client(t)

	name := "supernice"
	ctx := context.Background()

	realm := &Realm{
		Enabled: Bool(true),
		ID:      String(name),
		Realm:   String(name),
	}

	res, err := k.Realms.Create(ctx, realm)
	if err != nil {
		t.Errorf("Realms.Create returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if res.Header.Get("Location") != "http://localhost:8080/auth/admin/realms/supernice" {
		t.Errorf("got: %s, want: %s", res.Header.Get("Location"), "http://localhost:8080/auth/admin/realms/supernice")
	}

	// manually clean up
	if _, err := k.Realms.Delete(ctx, name); err != nil {
		t.Errorf("Realms.Delete returned error: %v", err)
	}
}

func TestRealmsService_List(t *testing.T) {
	k := client(t)

	createRealm(t, k, "first")
	createRealm(t, k, "second")

	realms, res, err := k.Realms.List(context.Background())
	if err != nil {
		t.Errorf("Realms.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	// it includes the master realm
	if len(realms) != 3 {
		t.Errorf("got: %d, want: %d", len(realms), 3)
	}
}

func TestRealmsService_Get(t *testing.T) {
	k := client(t)

	createRealm(t, k, "first")

	realm, res, err := k.Realms.Get(context.Background(), "first")
	if err != nil {
		t.Errorf("Realms.Get returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *realm.ID != "first" {
		t.Errorf("got: %s, want: %s", *realm.ID, "first")
	}
}

func TestRealmsService_Delete(t *testing.T) {
	k := client(t)

	createRealm(t, k, "first")

	res, err := k.Realms.Delete(context.Background(), "first")
	if err != nil {
		t.Errorf("Realms.Delete returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}

func TestRealmsService_GetConfig(t *testing.T) {
	k := client(t)

	createRealm(t, k, "first")

	config, res, err := k.Realms.GetConfig(context.Background(), "first")
	if err != nil {
		t.Errorf("Realms.GetConfig returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *config.Issuer != "http://localhost:8080/auth/realms/first" {
		t.Errorf("got: %s, want: %s", *config.Issuer, "http://localhost:8080/auth/realms/first")
	}
}
