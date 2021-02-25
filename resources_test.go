package keycloak

import (
	"context"
	"net/http"
	"testing"
)

// create a new resource.
func createResource(t *testing.T, k *Keycloak, realm, clientID, name string) *Resource {
	t.Helper()

	resource := &Resource{
		Name:        String(name),
		DisplayName: String(name),
	}

	created, _, err := k.Resources.Create(context.Background(), realm, clientID, resource)
	if err != nil {
		t.Errorf("Resources.Create returned error: %v", err)
	}

	return created
}

func TestResourcesService_Create(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	ctx := context.Background()

	resource := &Resource{
		Name:        String("resource"),
		DisplayName: String("resource"),
	}

	created, res, err := k.Resources.Create(ctx, realm, clientID, resource)
	if err != nil {
		t.Errorf("Clients.CreateResource returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *created.Name != "resource" {
		t.Errorf("got: %s, want: %s", *resource.Name, "Default Resource")
	}
}

func TestResourcesService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	resources, res, err := k.Resources.List(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Clients.ListResources returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if len(resources) != 1 {
		t.Errorf("got: %d, want: %d", len(resources), 1)
	}
}

func TestResourcesService_Get(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	// list all resources first
	resources, res, err := k.Resources.List(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Clients.ListResources returned error: %v", err)
	}

	resource, res, err := k.Resources.Get(context.Background(), realm, clientID, *resources[0].ID)
	if err != nil {
		t.Errorf("Clients.GetResource returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *resource.Name != "Default Resource" {
		t.Errorf("got: %s, want: %s", *resource.Name, "Default Resource")
	}
}

func TestResourcesService_Deletee(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)

	clientID := createClient(t, k, realm, "client")

	ctx := context.Background()

	resource := &Resource{
		Name:        String("resource"),
		DisplayName: String("resource"),
	}

	created, res, err := k.Resources.Create(ctx, realm, clientID, resource)
	if err != nil {
		t.Errorf("Clients.CreateResource returned error: %v", err)
	}

	res, err = k.Resources.Delete(ctx, realm, clientID, *created.ID)
	if err != nil {
		t.Errorf("Clients.DeleteResource returned error: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusNoContent)
	}
}
