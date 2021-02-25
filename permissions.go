package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// PermissionsService handles communication with the permissions related methods of the Keycloak API.
type PermissionsService service

// Permission represents a Keycloak abstract permission.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/AbstractPolicyRepresentation.java
type Permission struct {
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

// ResourcePermission represents a Keycloak resource permission.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/ResourcePermissionRepresentation.java
type ResourcePermission struct {
	Permission
	ResourceType *string `json:"resourceType,omitempty"`
}

// ScopePermission represents a Keycloak scope permission.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/ScopePermissionRepresentation.java
type ScopePermission struct {
	Permission
	ResourceType *string `json:"resourceType,omitempty"`
}

// List lists all permissions.
func (s *PermissionsService) List(ctx context.Context, realm, clientID string) ([]*Permission, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/permission", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var permission []*Permission
	res, err := s.keycloak.Do(ctx, req, &permission)
	if err != nil {
		return nil, nil, err
	}

	return permission, res, nil
}

// CreateResourcePermission creates a new resource based permission.
func (s *PermissionsService) CreateResourcePermission(ctx context.Context, realm, clientID string, permission *ResourcePermission) (*ResourcePermission, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/permission/resource", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, permission)
	if err != nil {
		return nil, nil, err
	}

	var created ResourcePermission
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}

// GetResourcePermission gets resource based permission by id.
func (s *PermissionsService) GetResourcePermission(ctx context.Context, realm, clientID, permissionID string) (*ResourcePermission, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/permission/resource/%s", realm, clientID, permissionID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var permission ResourcePermission
	res, err := s.keycloak.Do(ctx, req, &permission)
	if err != nil {
		return nil, nil, err
	}

	return &permission, res, nil
}

// GetScopePermission gets scope based permission by id.
func (s *PermissionsService) GetScopePermission(ctx context.Context, realm, clientID, permissionID string) (*ScopePermission, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/permission/scope/%s", realm, clientID, permissionID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var permission ScopePermission
	res, err := s.keycloak.Do(ctx, req, &permission)
	if err != nil {
		return nil, nil, err
	}

	return &permission, res, nil
}

// CreateScopePermission creates a new scope based permission.
func (s *PermissionsService) CreateScopePermission(ctx context.Context, realm, clientID string, permission *ScopePermission) (*ScopePermission, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/authz/resource-server/permission/scope", realm, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, permission)
	if err != nil {
		return nil, nil, err
	}

	var created ScopePermission
	res, err := s.keycloak.Do(ctx, req, &created)
	if err != nil {
		return nil, nil, err
	}

	return &created, res, nil
}
