package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// ClientsService handles communication with the client related methods of the Keycloak API.
type ClientsService service

// Client represents a Keycloak client.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/ClientRepresentation.java
type Client struct {
	ID                                 *string            `json:"id,omitempty"`
	ClientID                           *string            `json:"clientId,omitempty"`
	Name                               *string            `json:"name,omitempty"`
	RootURL                            *string            `json:"rootUrl,omitempty"`
	BaseURL                            *string            `json:"baseUrl,omitempty"`
	SurrogateAuthRequired              *bool              `json:"surrogateAuthRequired,omitempty"`
	Enabled                            *bool              `json:"enabled,omitempty"`
	AlwaysDisplayInConsole             *bool              `json:"alwaysDisplayInConsole,omitempty"`
	ClientAuthenticatorType            *string            `json:"clientAuthenticatorType,omitempty"`
	DefaultRoles                       []string           `json:"defaultRoles,omitempty"`
	RedirectUris                       []string           `json:"redirectUris,omitempty"`
	WebOrigins                         []string           `json:"webOrigins,omitempty"`
	NotBefore                          *int               `json:"notBefore,omitempty"`
	BearerOnly                         *bool              `json:"bearerOnly,omitempty"`
	ConsentRequired                    *bool              `json:"consentRequired,omitempty"`
	StandardFlowEnabled                *bool              `json:"standardFlowEnabled,omitempty"`
	ImplicitFlowEnabled                *bool              `json:"implicitFlowEnabled,omitempty"`
	DirectAccessGrantsEnabled          *bool              `json:"directAccessGrantsEnabled,omitempty"`
	ServiceAccountsEnabled             *bool              `json:"serviceAccountsEnabled,omitempty"`
	AuthorizationServicesEnabled       *bool              `json:"authorizationServicesEnabled,omitempty"`
	PublicClient                       *bool              `json:"publicClient,omitempty"`
	FrontchannelLogout                 *bool              `json:"frontchannelLogout,omitempty"`
	Protocol                           *string            `json:"protocol,omitempty"`
	Attributes                         *map[string]string `json:"attributes,omitempty"`
	AuthenticationFlowBindingOverrides *map[string]string `json:"authenticationFlowBindingOverrides,omitempty"`
	FullScopeAllowed                   *bool              `json:"fullScopeAllowed,omitempty"`
	NodeReRegistrationTimeout          *int               `json:"nodeReRegistrationTimeout,omitempty"`
	DefaultClientScopes                []string           `json:"defaultClientScopes,omitempty"`
	OptionalClientScopes               []string           `json:"optionalClientScopes,omitempty"`
	Access                             *map[string]bool   `json:"access,omitempty"`
}

// List all clients in realm.
func (s *ClientsService) List(ctx context.Context, realm string) ([]*Client, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients", realm)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var clients []*Client
	res, err := s.keycloak.Do(ctx, req, &clients)
	if err != nil {
		return nil, nil, err
	}

	return clients, res, nil
}

// Create a new client.
func (s *ClientsService) Create(ctx context.Context, realm string, client *Client) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients", realm)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, client)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// Update a new client.
func (s *ClientsService) Update(ctx context.Context, realm string, client *Client) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s", realm, *client.ID)
	req, err := s.keycloak.NewRequest(http.MethodPut, u, client)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// Get client.
func (s *ClientsService) Get(ctx context.Context, realm, id string) (*Client, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var client Client
	res, err := s.keycloak.Do(ctx, req, &client)
	if err != nil {
		return nil, nil, err
	}

	return &client, res, nil
}

// Delete ...
func (s *ClientsService) Delete() {

}

// CreateRole creates a new client role.
func (s *ClientsService) CreateRole(ctx context.Context, realm, id string, role *Role) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/roles", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, role)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// ListRoles lists all client roles.
func (s *ClientsService) ListRoles(ctx context.Context, realm, id string) ([]*Role, *http.Response, error) {
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

// GetSecret gets client secret.
func (s *ClientsService) GetSecret(ctx context.Context, realm, id string) (*Credential, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s/client-secret", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var credential Credential
	res, err := s.keycloak.Do(ctx, req, &credential)
	if err != nil {
		return nil, nil, err
	}

	return &credential, res, nil
}

// CreateSecret generates a new secret for the client
func (s *ClientsService) CreateSecret(ctx context.Context, realm, id string) (*Credential, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s/clients/%s", realm, id)
	req, err := s.keycloak.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var credential Credential
	res, err := s.keycloak.Do(ctx, req, &credential)
	if err != nil {
		return nil, nil, err
	}

	return &credential, res, nil
}

// Options ...
type Options struct {
	First int    `url:"first,omitempty"`
	Max   string `url:"max,omitempty"`
}

// GetUsersInRole returns a stream of users that have the specified role name.
func (s *ClientsService) GetUsersInRole(ctx context.Context, realm, clientID, role string, opts *Options) ([]*User, *http.Response, error) {
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
