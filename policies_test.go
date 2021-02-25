package keycloak

import (
	"context"
	"net/http"
	"testing"
)

func createUserPolicy(t *testing.T, k *Keycloak, realm, clientID, userID string) *UserPolicy {
	policy := &UserPolicy{
		Policy: Policy{
			Type:             String("user"),
			Logic:            String("POSITIVE"),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("policy"),
		},
		Users: []string{userID},
	}

	policy, _, err := k.Policies.CreateUserPolicy(context.Background(), realm, clientID, policy)
	if err != nil {
		t.Errorf("Policies.CreateUserPolicy returned error: %v", err)
	}
	return policy
}

func TestPoliciesService_CreateUserPolicy(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")
	userID := createUser(t, k, realm, "john")

	policy := &UserPolicy{
		Policy: Policy{
			Type:             String("user"),
			Logic:            String("POSITIVE"),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("policy"),
		},
		Users: []string{userID},
	}

	policy, res, err := k.Policies.CreateUserPolicy(context.Background(), realm, clientID, policy)
	if err != nil {
		t.Errorf("Policies.CreateUserPolicy returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *policy.Name != "policy" {
		t.Errorf("got: %s, want: %s", *policy.Name, "policy")
	}
}

func TestPoliciesService_CreateRolePolicy(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	createRealmRole(t, k, realm, "role")

	role, res, err := k.Roles.GetByName(context.Background(), realm, "role")
	if err != nil {
		t.Errorf("Roles.GetByName returned error: %v", err)
	}

	policy := &RolePolicy{
		Policy: Policy{
			Type:             String("role"),
			Logic:            String("POSITIVE"),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("policy"),
		},
		Roles: []*RoleDefinition{{
			ID: role.ID,
		}},
	}

	policy, res, err = k.Policies.CreateRolePolicy(context.Background(), realm, clientID, policy)
	if err != nil {
		t.Errorf("Policies.CreateRolePolicy returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusCreated)
	}

	if *policy.Name != "policy" {
		t.Errorf("got: %s, want: %s", *policy.Name, "policy")
	}
}

func TestPoliciesService_List(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	roles, res, err := k.Policies.List(context.Background(), realm, clientID)
	if err != nil {
		t.Errorf("Policies.List returned error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	if len(roles) != 1 {
		t.Errorf("got: %d, want: %d", len(roles), 1)
	}
}
