package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// Role representation
type Role struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	Composite   *bool                `json:"composite,omitempty"`
	ClientRole  *bool                `json:"clientRole,omitempty"`
	ContainerID *string              `json:"containerId,omitempty"`
	Attributes  *map[string][]string `json:"attributes,omitempty"`
}

// RolesService ...
type RolesService service

// Create a new role.
func (s *RolesService) Create(ctx context.Context, realm string, role *Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles", realm)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, role)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// RolesListOptions ...
type RolesListOptions struct {
	BriefRepresentation bool `url:"briefRepresentation,omitempty"`
	Search              bool `url:"search,omitempty"`
	Options
}

// List roles.
func (s *RolesService) List(ctx context.Context, realm string, opts *RolesListOptions) ([]*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles", realm)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

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

// GetByName gets role by name.
func (s *RolesService) GetByName(ctx context.Context, realm, name string) (*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles/%s", realm, name)
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

// GetByID gets role by id.
func (s *RolesService) GetByID(ctx context.Context, realm, roleID string) (*Role, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles-by-id/%s", realm, roleID)
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

// GetClientRoleByID gets client role by id.
func (s *RolesService) GetClientRoleByID(ctx context.Context, realm, roleID, clientID string) (*Role, *http.Response, error) {
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

// DeleteClientRoleByID deletes client role by id.
func (s *RolesService) DeleteClientRoleByID(ctx context.Context, realm, roleID, clientID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/roles-by-id/%s?client=%s", realm, roleID, clientID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	return s.keycloak.Do(ctx, req, nil)
}

// delete role

// delete client role
