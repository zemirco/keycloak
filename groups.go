package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// Group ...
type Group struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Path        *string              `json:"path,omitempty"`
	Attributes  *map[string][]string `json:"attributes,omitempty"`
	RealmRoles  []string             `json:"realmRoles,omitempty"`
	ClientRoles *map[string][]string `json:"clientRoles,omitempty"`
	SubGroups   []*Group             `json:"subGroups,omitempty"`
	Access      *map[string]bool     `json:"access,omitempty"`
}

// GroupsService ...
type GroupsService service

// Create a new group.
func (s *GroupsService) Create(ctx context.Context, realm string, group *Group) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups", realm)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, group)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// List groups.
func (s *GroupsService) List(ctx context.Context, realm string) ([]*Group, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups", realm)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var groups []*Group
	res, err := s.keycloak.Do(ctx, req, &groups)
	if err != nil {
		return nil, nil, err
	}

	return groups, res, nil
}

// Get group.
func (s *GroupsService) Get(ctx context.Context, realm, groupID string) (*Group, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups/%s", realm, groupID)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var group Group
	res, err := s.keycloak.Do(ctx, req, &group)
	if err != nil {
		return nil, nil, err
	}

	return &group, res, nil
}

// update group

// Delete group.
func (s *GroupsService) Delete(ctx context.Context, realm, groupID string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups/%s", realm, groupID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

func (s *GroupsService) AddRealmRoles(ctx context.Context, realm, groupID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups/%s/role-mappings/realm", realm, groupID)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

func (s *GroupsService) RemoveRealmRoles(ctx context.Context, realm, groupID string, roles []*Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/groups/%s/role-mappings/realm", realm, groupID)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, roles)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}
