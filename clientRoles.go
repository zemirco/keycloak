package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// ClientRolesService handles communication with the client roles related methods of the Keycloak API.
type ClientRolesService service

// Create creates a new client role.
func (s *ClientRolesService) Create(ctx context.Context, realm, id string, role *Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/roles", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, role)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// List lists all client roles.
func (s *ClientRolesService) List(ctx context.Context, realm, id string) ([]*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/roles", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var roles []*Role
	res, err := s.keycloak.Do(ctx, req, &roles)
	if err != nil {
		return nil, nil, err
	}

	return roles, res, nil
}

// Get retrieves a single client role.
func (s *ClientRolesService) Get(ctx context.Context, realm, id, roleName string) (*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/roles/%s", realm, id, roleName)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var role *Role
	res, err := s.keycloak.Do(ctx, req, &role)
	if err != nil {
		return nil, nil, err
	}

	return role, res, nil
}

// GetUsers returns a stream of users that have the specified role name.
func (s *ClientRolesService) GetUsers(ctx context.Context, realm, clientID, role string, opts *Options) ([]*User, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/roles/%s/users", realm, clientID, role)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	res, err := s.keycloak.Do(ctx, req, &users)
	if err != nil {
		return nil, nil, err
	}

	return users, res, nil
}

// GetByID gets client role by id.
func (s *ClientRolesService) GetByID(ctx context.Context, realm, roleID, clientID string) (*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles-by-id/%s?client=%s", realm, roleID, clientID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var role Role
	res, err := s.keycloak.Do(ctx, req, &role)
	if err != nil {
		return nil, nil, err
	}

	return &role, res, nil
}

// DeleteByID deletes client role by id.
func (s *ClientRolesService) DeleteByID(ctx context.Context, realm, roleID, clientID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles-by-id/%s?client=%s", realm, roleID, clientID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	return s.keycloak.Do(ctx, req, nil)
}

// Returns a stream of groups that have the specified role name
// GET /{realm}/clients/{id}/roles/{role-name}/groups

// Add a composite to the role
// POST /{realm}/clients/{id}/roles/{role-name}/composites

// Get composites of the role
// GET /{realm}/clients/{id}/roles/{role-name}/composites

// Remove roles from the role’s composite
// DELETE /{realm}/clients/{id}/roles/{role-name}/composites
