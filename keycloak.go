package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

// Keycloak ...
type Keycloak struct {
	client *http.Client

	BaseURL *url.URL

	common service

	Clients      *ClientsService
	ClientRoles  *ClientRolesService
	ClientScopes *ClientScopesService
	Groups       *GroupsService
	Permissions  *PermissionsService
	Policies     *PoliciesService
	Realms       *RealmsService
	Resources    *ResourcesService
	Roles        *RolesService
	Scopes       *ScopesService
	Users        *UsersService
}

type service struct {
	keycloak *Keycloak
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewKeycloak ...
func NewKeycloak(httpClient *http.Client, baseURL string) (*Keycloak, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	uri, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	k := &Keycloak{
		client:  httpClient,
		BaseURL: uri,
	}

	k.common.keycloak = k
	k.Clients = (*ClientsService)(&k.common)
	k.ClientRoles = (*ClientRolesService)(&k.common)
	k.ClientScopes = (*ClientScopesService)(&k.common)
	k.Groups = (*GroupsService)(&k.common)
	k.Permissions = (*PermissionsService)(&k.common)
	k.Policies = (*PoliciesService)(&k.common)
	k.Realms = (*RealmsService)(&k.common)
	k.Resources = (*ResourcesService)(&k.common)
	k.Roles = (*RolesService)(&k.common)
	k.Scopes = (*ScopesService)(&k.common)
	k.Users = (*UsersService)(&k.common)

	return k, nil
}

// NewRequest ...
func (k *Keycloak) NewRequest(method string, url string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(k.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", k.BaseURL)
	}
	u, err := k.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var b io.ReadWriter
	if body != nil {
		b = &bytes.Buffer{}
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do ...
func (k *Keycloak) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	res, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return nil, err
		}
	}

	return res, err
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// // Int is a helper routine that allocates a new int value
// // to store v and returns a pointer to it.
// func Int(v int) *int { return &v }

// // Int64 is a helper routine that allocates a new int64 value
// // to store v and returns a pointer to it.
// func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
