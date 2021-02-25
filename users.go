package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// User representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/UserRepresentation.java
type User struct {
	ID                         *string              `json:"id,omitempty"`
	CreatedTimestamp           *int64               `json:"createdTimestamp,omitempty"`
	Username                   *string              `json:"username,omitempty"`
	Enabled                    *bool                `json:"enabled,omitempty"`
	Totp                       *bool                `json:"totp,omitempty"`
	EmailVerified              *bool                `json:"emailVerified,omitempty"`
	FirstName                  *string              `json:"firstName,omitempty"`
	LastName                   *string              `json:"lastName,omitempty"`
	Email                      *string              `json:"email,omitempty"`
	DisableableCredentialTypes []string             `json:"disableableCredentialTypes,omitempty"`
	RequiredActions            []string             `json:"requiredActions,omitempty"`
	NotBefore                  *int                 `json:"notBefore,omitempty"`
	Access                     *map[string]bool     `json:"access,omitempty"`
	Attributes                 *map[string][]string `json:"attributes,omitempty"`
}

// Credential representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/CredentialRepresentation.java
type Credential struct {
	Type      *string `json:"type,omitempty"`
	Value     *string `json:"value,omitempty"`
	Temporary *bool   `json:"temporary,omitempty"`
}

// UsersService ...
type UsersService service

// Create a new user.
func (s *UsersService) Create(ctx context.Context, realm string, user *User) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users", realm)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, user)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// List users.
func (s *UsersService) List(ctx context.Context, realm string) ([]*User, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users", realm)
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

// GetByUsername get a single user by username.
func (s *UsersService) GetByUsername(ctx context.Context, realm, username string) ([]*User, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users?username=%s", realm, username)
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

// Delete user.
func (s *UsersService) Delete(ctx context.Context, realm, userID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s", realm, userID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// ResetPassword sets or resets the user's password.
func (s *UsersService) ResetPassword(ctx context.Context, realm, userID string, credential *Credential) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/reset-password", realm, userID)
	req, err := s.keycloak.NewRequest(http.MethodPut, u, credential)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// Update user.

// JoinGroup adds user to a group.
func (s *UsersService) JoinGroup(ctx context.Context, realm, userID, groupID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/groups/%s", realm, userID, groupID)
	req, err := s.keycloak.NewRequest(http.MethodPut, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// LeaveGroup removes a user from a group.
func (s *UsersService) LeaveGroup(ctx context.Context, realm, userID, groupID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/groups/%s", realm, userID, groupID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// AddRealmRoles adds realm roles to user.
func (s *UsersService) AddRealmRoles(ctx context.Context, realm, userID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/role-mappings/realm", realm, userID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// RemoveRealmRoles removes assigned realm roles from user.
func (s *UsersService) RemoveRealmRoles(ctx context.Context, realm, userID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/role-mappings/realm", realm, userID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// AddClientRoles adds client roles to user.
func (s *UsersService) AddClientRoles(ctx context.Context, realm, userID, clientID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/role-mappings/clients/%s", realm, userID, clientID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// RemoveClientRoles removes assigned client roles from user.
func (s *UsersService) RemoveClientRoles(ctx context.Context, realm, userID, clientID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/users/%s/role-mappings/clients/%s", realm, userID, clientID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}
