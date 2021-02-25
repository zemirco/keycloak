package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// ScopesService handles communication with the scopes related methods of the Keycloak API.
type ScopesService service

// Scope represents a Keycloak scope.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/ScopeRepresentation.java
type Scope struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	IconURI     *string `json:"iconUri,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// List lists all resources.
func (s *ScopesService) List(ctx context.Context, realm, clientID string) ([]*Scope, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/scope", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var scopes []*Scope
	res, err := s.keycloak.Do(ctx, req, &scopes)
	if err != nil {
		return nil, nil, err
	}

	return scopes, res, nil
}

// Create creates a new scope.
func (s *ScopesService) Create(ctx context.Context, realm, clientID string, scope *Scope) (*Scope, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/scope", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, scope)
	if err != nil {
		return nil, nil, err
	}

	var created Scope
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}

// Get gets a single scope.
func (s *ScopesService) Get(ctx context.Context, realm, clientID, scopeID string) (*Scope, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/scope/%s", realm, clientID, scopeID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var scope Scope
	res, err := s.keycloak.Do(ctx, req, &scope)
	if err != nil {
		return nil, nil, err
	}

	return &scope, res, nil
}

// Delete deletes a single scope.
func (s *ScopesService) Delete(ctx context.Context, realm, clientID, scopeID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/scope/%s", realm, clientID, scopeID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// Update creates a new scope.
func (s *ScopesService) Update(ctx context.Context, realm, clientID string, scope *Scope) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/scope/%s", realm, clientID, *scope.ID)
	req, err := s.keycloak.NewRequest(http.MethodPut, u, scope)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}
