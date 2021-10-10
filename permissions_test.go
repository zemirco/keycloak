package keycloak

import (
	"context"
	"net/http"
	"testing"
)

func TestPermissionsService_CreateResourcePermission(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")
	policy := createUserPolicy(t, k, realm, clientID, userID)
	resource := createResource(t, k, realm, clientID, "resource")

	// create first permission
	permission := &ResourcePermission{
		Permission: Permission{
			Type:             String("resource"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("permission"),
			Resources:        []string{*resource.ID},
			Policies:         []string{*policy.ID},
		},
		ResourceType: String(""),
	}
	permission, res, err := k.Permissions.CreateResourcePermission(context.Background(), realm, clientID, permission)
	if err != nil {
		t.Errorf("Permissions.CreateResourcePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *permission.Name != "permission" {
		t.Errorf("got: %s, want: %s", *permission.Name, "permission")
	}
}

func TestPermissionsService_GetResourcePermission(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")
	policy := createUserPolicy(t, k, realm, clientID, userID)
	resource := createResource(t, k, realm, clientID, "resource")

	// create first permission
	permission := &ResourcePermission{
		Permission: Permission{
			Type:             String("resource"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("permission"),
			Resources:        []string{*resource.ID},
			Policies:         []string{*policy.ID},
		},
		ResourceType: String(""),
	}
	permission, res, err := k.Permissions.CreateResourcePermission(context.Background(), realm, clientID, permission)
	if err != nil {
		t.Errorf("Permissions.CreateResourcePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	permission, res, err = k.Permissions.GetResourcePermission(context.Background(), realm, clientID, *permission.ID)
	if err != nil {
		t.Errorf("Permissions.GetResourcePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *permission.Name != "permission" {
		t.Errorf("got: %s, want: %s", *permission.Name, "permission")
	}
}

func TestPermissionsService_CreateScopePermission(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")
	policy := createUserPolicy(t, k, realm, clientID, userID)
	resource := createResource(t, k, realm, clientID, "resource")
	scope := createScope(t, k, realm, clientID, "scope")

	// create first permission
	permission := &ScopePermission{
		Permission: Permission{
			Type:             String("resource"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("permission"),
			Resources:        []string{*resource.ID},
			Policies:         []string{*policy.ID},
			Scopes:           []string{*scope.ID},
		},
		ResourceType: String(""),
	}
	permission, res, err := k.Permissions.CreateScopePermission(context.Background(), realm, clientID, permission)
	if err != nil {
		t.Errorf("Permissions.CreateScopePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *permission.Name != "permission" {
		t.Errorf("got: %s, want: %s", *permission.Name, "permission")
	}
}

func TestPermissionsService_GetScopePermission(t *testing.T) {
	k := client(t)

	ctx := context.Background()

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")
	policy := createUserPolicy(t, k, realm, clientID, userID)
	resource := createResource(t, k, realm, clientID, "resource")
	scope := createScope(t, k, realm, clientID, "scope")

	// create first permission
	permission := &ScopePermission{
		Permission: Permission{
			Type:             String("resource"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("permission"),
			Resources:        []string{*resource.ID},
			Policies:         []string{*policy.ID},
			Scopes:           []string{*scope.ID},
		},
		ResourceType: String(""),
	}
	permission, res, err := k.Permissions.CreateScopePermission(ctx, realm, clientID, permission)
	if err != nil {
		t.Errorf("Permissions.CreateScopePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	permission, res, err = k.Permissions.GetScopePermission(ctx, realm, clientID, *permission.ID)
	if err != nil {
		t.Errorf("Permissions.GetResourcePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if *permission.Name != "permission" {
		t.Errorf("got: %s, want: %s", *permission.Name, "permission")
	}
}

func TestPermissionsService_List(t *testing.T) {
	k := client(t)

	ctx := context.Background()

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")
	policy := createUserPolicy(t, k, realm, clientID, userID)
	resource := createResource(t, k, realm, clientID, "resource")
	scope := createScope(t, k, realm, clientID, "scope")

	// create first permission
	permission := &ScopePermission{
		Permission: Permission{
			Type:             String("resource"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("permission"),
			Resources:        []string{*resource.ID},
			Policies:         []string{*policy.ID},
			Scopes:           []string{*scope.ID},
		},
		ResourceType: String(""),
	}
	_, res, err := k.Permissions.CreateScopePermission(ctx, realm, clientID, permission)
	if err != nil {
		t.Errorf("Permissions.CreateScopePermission returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	permissions, res, err := k.Permissions.List(ctx, realm, clientID)
	if err != nil {
		t.Errorf("Permissions.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if len(permissions) != 2 {
		t.Errorf("got: %d, want: %d", len(permissions), 2)
	}
}
