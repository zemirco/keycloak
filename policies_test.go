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
			Logic:            String(LogicPositive),
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
			Logic:            String(LogicPositive),
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

	role, _, err := k.RealmRoles.GetByName(context.Background(), realm, "role")
	if err != nil {
		t.Errorf("RealmRoles.GetByName returned error: %v", err)
	}

	policy := &RolePolicy{
		Policy: Policy{
			Type:             String("role"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("policy"),
		},
		Roles: []*RoleDefinition{{
			ID: role.ID,
		}},
	}

	policy, res, err := k.Policies.CreateRolePolicy(context.Background(), realm, clientID, policy)
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

func TestPoliciesService_CreateGroupPolicy(t *testing.T) {
	k := client(t)

	realm := "first"
	createRealm(t, k, realm)
	clientID := createClient(t, k, realm, "client")

	groupID := createGroup(t, k, realm, "group")

	policy := &GroupPolicy{
		Policy: Policy{
			Type:             String("role"),
			Logic:            String(LogicPositive),
			DecisionStrategy: String(DecisionStrategyUnanimous),
			Name:             String("policy"),
		},
		Groups: []*GroupDefinition{{
			ID: &groupID,
		}},
	}

	policy, res, err := k.Policies.CreateGroupPolicy(context.Background(), realm, clientID, policy)
	if err != nil {
		t.Errorf("Policies.CreateGroupPolicy returned error: %v", err)
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
