package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// PoliciesService handles communication with the policies related methods of the Keycloak API.
type PoliciesService service

// Policy represents a Keycloak abstract policy.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/AbstractPolicyRepresentation.java
type Policy struct {
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
	Description      *string  `json:"description,omitempty"`
	Type             *string  `json:"type,omitempty"`
	Policies         []string `json:"policies,omitempty"`
	Resources        []string `json:"resources,omitempty"`
	Scopes           []string `json:"scopes,omitempty"`
	Logic            *string  `json:"logic,omitempty"`
	DecisionStrategy *string  `json:"decisionStrategy,omitempty"`
	Owner            *string  `json:"owner,omitempty"`
}

// GroupDefinition represents a Keycloak groupDefinition.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/GroupPolicyRepresentation.java
type GroupDefinition struct {
	ID             *string `json:"id,omitempty"`
	Path           *string `json:"path,omitempty"`
	ExtendChildren *bool   `json:"extendChildren,omitempty"`
}

// GroupPolicy represents a Keycloak group policy.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/GroupPolicyRepresentation.java
type GroupPolicy struct {
	Policy
	GroupsClaim *string            `json:"groupsClaim,omitempty"`
	Groups      []*GroupDefinition `json:"groups,omitempty"`
}

// UserPolicy represents a Keycloak user policy.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/UserPolicyRepresentation.java
type UserPolicy struct {
	Policy
	Users []string `json:"users,omitempty"`
}

// RolePolicy represents a Keycloak role policy.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/RolePolicyRepresentation.java
type RolePolicy struct {
	Policy
	Roles []*RoleDefinition `json:"roles,omitempty"`
}

// RoleDefinition represents a Keycloak role definition.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/RolePolicyRepresentation.java
type RoleDefinition struct {
	ID       *string `json:"id,omitempty"`
	Required *bool   `json:"required,omitempty"`
}

// List lists all policies.
func (s *PoliciesService) List(ctx context.Context, realm, clientID string) ([]*Policy, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/policy?permission=false", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var policies []*Policy
	res, err := s.keycloak.Do(ctx, req, &policies)
	if err != nil {
		return nil, nil, err
	}

	return policies, res, nil
}

// CreateUserPolicy creates a new user policy.
func (s *PoliciesService) CreateUserPolicy(ctx context.Context, realm, clientID string, policy *UserPolicy) (*UserPolicy, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/policy/user", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	var created UserPolicy
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}

// CreateRolePolicy creates a new role policy.
func (s *PoliciesService) CreateRolePolicy(ctx context.Context, realm, clientID string, policy *RolePolicy) (*RolePolicy, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/policy/role", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	var created RolePolicy
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}

// CreateGroupPolicy creates a new group policy.
func (s *PoliciesService) CreateGroupPolicy(ctx context.Context, realm, clientID string, policy *GroupPolicy) (*GroupPolicy, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/policy/group", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	var created GroupPolicy
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}
