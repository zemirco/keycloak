package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// ClientScope representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/ClientScopeRepresentation.java
type ClientScope struct {
	ID          *string            `json:"id,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Description *string            `json:"description,omitempty"`
	Protocol    *string            `json:"protocol,omitempty"`
	Attributes  *map[string]string `json:"attributes,omitempty"`
}

// ClientScopesService ...
type ClientScopesService service

// List all client scopes in realm.
func (s *ClientScopesService) List(ctx context.Context, realm string) ([]*ClientScope, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/client-scopes", realm)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientScopes []*ClientScope
	res, err := s.keycloak.Do(ctx, req, &clientScopes)
	if err != nil {
		return nil, nil, err
	}

	return clientScopes, res, nil
}

// Create a new client scope.
func (s *ClientScopesService) Create(ctx context.Context, realm string, clientScope *ClientScope) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/client-scopes", realm)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, clientScope)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// Get client scope.
func (s *ClientScopesService) Get(ctx context.Context, realm, clientScopeID string) (*ClientScope, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/client-scopes/%s", realm, clientScopeID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientScope ClientScope
	res, err := s.keycloak.Do(ctx, req, &clientScope)
	if err != nil {
		return nil, nil, err
	}

	return &clientScope, res, nil
}

// Delete client scope.
func (s *ClientScopesService) Delete(ctx context.Context, realm, clientScopeID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/client-scopes/%s", realm, clientScopeID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}
