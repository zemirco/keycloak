package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// ResourcesService handles communication with the resources related methods of the Keycloak API.
type ResourcesService service

// ResourceOwner represents a Keycloak resource owner.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/ResourceOwnerRepresentation.java
type ResourceOwner struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Resource represents a Keycloak resource.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/ResourceRepresentation.java
type Resource struct {
	ID                 *string              `json:"_id,omitempty"`
	Scopes             []*Scope             `json:"scopes,omitempty"`
	Attributes         *map[string][]string `json:"attributes,omitempty"`
	Uris               []string             `json:"uris,omitempty"`
	Name               *string              `json:"name,omitempty"`
	OwnerManagedAccess *bool                `json:"ownerManagedAccess,omitempty"`
	DisplayName        *string              `json:"displayName,omitempty"`
	Owner              *ResourceOwner       `json:"owner,omitempty"`
}

// List lists all resources.
func (s *ResourcesService) List(ctx context.Context, realm, clientID string) ([]*Resource, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/resource", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resources []*Resource
	res, err := s.keycloak.Do(ctx, req, &resources)
	if err != nil {
		return nil, nil, err
	}

	return resources, res, nil
}

// Create creates a new resource.
func (s *ResourcesService) Create(ctx context.Context, realm, clientID string, resource *Resource) (*Resource, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/resource", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, resource)
	if err != nil {
		return nil, nil, err
	}

	var created Resource
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}

// Get gets a single resource.
func (s *ResourcesService) Get(ctx context.Context, realm, clientID, resourceID string) (*Resource, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/resource/%s", realm, clientID, resourceID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resource Resource
	res, err := s.keycloak.Do(ctx, req, &resource)
	if err != nil {
		return nil, nil, err
	}

	return &resource, res, nil
}

// Delete deletes a single resource.
func (s *ResourcesService) Delete(ctx context.Context, realm, clientID, resourceID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/resource/%s", realm, clientID, resourceID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}
