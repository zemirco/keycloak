package keycloak_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/zemirco/keycloak"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	realm    = "gorealm"
	clientID = "goclient"
)

func Example_full() {
	// create a new config for the "admin-cli" client
	config := oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/realms/master/protocol/openid-connect/token",
		},
		Scopes: []string{"openid"},
	}

	ctx := context.Background()

	// get a token
	token, err := config.PasswordCredentialsToken(ctx, "admin", "admin")
	if err != nil {
		panic(err)
	}

	// use the token on every http request
	httpClient := config.Client(ctx, token)

	// create a new keycloak client instance
	kc, err := keycloak.NewKeycloak(httpClient, "http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	// create a new realm
	r := &keycloak.Realm{
		Enabled: keycloak.Bool(true),
		ID:      keycloak.String(realm),
		Realm:   keycloak.String(realm),
	}

	if _, err := kc.Realms.Create(ctx, r); err != nil {
		panic(err)
	}

	// clean up
	// remove this or comment out to keep the realm and play around in the keycloak gui
	defer func() {
		if _, err := kc.Realms.Delete(ctx, realm); err != nil {
			panic(err)
		}
	}()

	// create client
	client := &keycloak.Client{
		Enabled:  keycloak.Bool(true),
		ClientID: keycloak.String(clientID),
		Protocol: keycloak.String("openid-connect"),
	}

	res, err := kc.Clients.Create(ctx, realm, client)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(res.Header.Get("Location"), "/")
	id := parts[len(parts)-1]

	// get the client to have all properties
	client, res, err = kc.Clients.Get(ctx, realm, id)
	if err != nil {
		panic(err)
	}

	// update client and enable authorization services
	client.RedirectUris = []string{"http://localhost:4200/*"}
	client.AuthorizationServicesEnabled = keycloak.Bool(true)
	client.ServiceAccountsEnabled = keycloak.Bool(true)
	client.PublicClient = keycloak.Bool(false)

	res, err = kc.Clients.Update(ctx, realm, client)
	if err != nil {
		panic(err)
	}

	// create user
	userA := &keycloak.User{
		Enabled:  keycloak.Bool(true),
		Username: keycloak.String("user_a"),
	}

	res, err = kc.Users.Create(ctx, realm, userA)
	if err != nil {
		panic(err)
	}

	parts = strings.Split(res.Header.Get("Location"), "/")
	userA.ID = &parts[len(parts)-1]

	// set password for user_a
	res, err = kc.Users.ResetPassword(ctx, realm, *userA.ID, &keycloak.Credential{
		Type:      keycloak.String("password"),
		Value:     keycloak.String("mypassword"),
		Temporary: keycloak.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	// create user
	userB := &keycloak.User{
		Enabled:  keycloak.Bool(true),
		Username: keycloak.String("user_b"),
	}

	res, err = kc.Users.Create(ctx, realm, userB)
	if err != nil {
		panic(err)
	}

	parts = strings.Split(res.Header.Get("Location"), "/")
	userB.ID = &parts[len(parts)-1]

	// set password for user_b
	res, err = kc.Users.ResetPassword(ctx, realm, *userB.ID, &keycloak.Credential{
		Type:      keycloak.String("password"),
		Value:     keycloak.String("mypassword"),
		Temporary: keycloak.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	// create scope
	scopeRead := &keycloak.Scope{
		Name: keycloak.String("scope_read"),
	}

	scopeRead, res, err = kc.Scopes.Create(ctx, realm, *client.ID, scopeRead)
	if err != nil {
		panic(err)
	}

	// create scope
	scopeWrite := &keycloak.Scope{
		Name: keycloak.String("scope_write"),
	}

	scopeWrite, res, err = kc.Scopes.Create(ctx, realm, *client.ID, scopeWrite)
	if err != nil {
		panic(err)
	}

	// create first resource
	resourceProject1 := &keycloak.Resource{
		Name: keycloak.String("Project_1"),
		Uris: []string{"/projects/1"},
		Scopes: []*keycloak.Scope{
			{
				ID:   keycloak.String(*scopeRead.ID),
				Name: keycloak.String("scope_read"),
			},
			{
				ID:   keycloak.String(*scopeWrite.ID),
				Name: keycloak.String("scope_write"),
			}},
	}

	resourceProject1, res, err = kc.Resources.Create(ctx, realm, *client.ID, resourceProject1)
	if err != nil {
		panic(err)
	}

	// create second resource
	resourceProject2 := &keycloak.Resource{
		Name: keycloak.String("Project_2"),
		Uris: []string{"/projects/2"},
		Scopes: []*keycloak.Scope{
			{
				ID:   keycloak.String(*scopeRead.ID),
				Name: keycloak.String("scope_read"),
			},
			{
				ID:   keycloak.String(*scopeWrite.ID),
				Name: keycloak.String("scope_write"),
			}},
	}

	resourceProject2, res, err = kc.Resources.Create(ctx, realm, *client.ID, resourceProject2)
	if err != nil {
		panic(err)
	}

	// create user policy with user_a
	userAPolicy := &keycloak.UserPolicy{
		Policy: keycloak.Policy{
			Type:             keycloak.String("user"),
			Logic:            keycloak.String(keycloak.LogicPositive),
			DecisionStrategy: keycloak.String(keycloak.DecisionStrategyUnanimous),
			Name:             keycloak.String("user_a_policy"),
		},
		Users: []string{*userA.ID},
	}

	userAPolicy, res, err = kc.Policies.CreateUserPolicy(ctx, realm, *client.ID, userAPolicy)
	if err != nil {
		panic(err)
	}

	// create user policy with user_b
	userBPolicy := &keycloak.UserPolicy{
		Policy: keycloak.Policy{
			Type:             keycloak.String("user"),
			Logic:            keycloak.String(keycloak.LogicPositive),
			DecisionStrategy: keycloak.String(keycloak.DecisionStrategyUnanimous),
			Name:             keycloak.String("user_b_policy"),
		},
		Users: []string{*userB.ID},
	}

	userBPolicy, res, err = kc.Policies.CreateUserPolicy(ctx, realm, *client.ID, userBPolicy)
	if err != nil {
		panic(err)
	}

	// create first permission
	permission1 := &keycloak.ScopePermission{
		Permission: keycloak.Permission{
			Type:             keycloak.String("resource"),
			Logic:            keycloak.String(keycloak.LogicPositive),
			DecisionStrategy: keycloak.String(keycloak.DecisionStrategyUnanimous),
			Name:             keycloak.String("permission__project_1__user_a_policy"),
			Resources:        []string{*resourceProject1.ID},
			Policies:         []string{*userAPolicy.ID},
			Scopes:           []string{*scopeRead.ID},
		},
		ResourceType: keycloak.String(""),
	}
	permission1, res, err = kc.Permissions.CreateScopePermission(ctx, realm, *client.ID, permission1)
	if err != nil {
		panic(err)
	}

	permission2 := &keycloak.ScopePermission{
		Permission: keycloak.Permission{
			Type:             keycloak.String("resource"),
			Logic:            keycloak.String(keycloak.LogicPositive),
			DecisionStrategy: keycloak.String(keycloak.DecisionStrategyUnanimous),
			Name:             keycloak.String("permission__project_2__user_b_policy"),
			Resources:        []string{*resourceProject2.ID},
			Policies:         []string{*userBPolicy.ID},
			Scopes:           []string{*scopeRead.ID},
		},
		ResourceType: keycloak.String(""),
	}
	permission2, res, err = kc.Permissions.CreateScopePermission(ctx, realm, *client.ID, permission2)
	if err != nil {
		panic(err)
	}

	// get client secret
	credential, _, err := kc.Clients.GetSecret(ctx, realm, *client.ID)
	if err != nil {
		panic(err)
	}

	// pretend to be user_a
	userAConfig := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: *credential.Value,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("http://localhost:8080/realms/%s/protocol/openid-connect/token", realm),
		},
	}

	userAToken, err := userAConfig.PasswordCredentialsToken(ctx, "user_a", "mypassword")
	if err != nil {
		panic(err)
	}

	clientConfig := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: *credential.Value,
		TokenURL:     fmt.Sprintf("http://localhost:8080/realms/%s/protocol/openid-connect/token", realm),
		EndpointParams: url.Values{
			"grant_type": {"urn:ietf:params:oauth:grant-type:uma-ticket"},
			"permission": {*resourceProject1.Name + "#scope_read"},
			"audience":   {clientID},
		},
	}

	ctx = context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.StaticTokenSource(userAToken),
		},
	})

	// ask keycloak for permission for reading Project_1
	token, err = clientConfig.Token(ctx)
	if err != nil {
		panic(err)
	}

	// ask keycloak for permission for reading Project_2
	clientConfig.EndpointParams.Set("permission", *resourceProject2.Name+"#scope_read")

	token, err = clientConfig.Token(ctx)
	if err == nil {
		panic("got no error, expected one")
	}
	retrieveError, ok := err.(*oauth2.RetrieveError)
	if !ok {
		panic(fmt.Errorf("got %T error, expected *RetrieveError; error was: %v", err, err))
	}

	fmt.Println(retrieveError.Response.Status)
	// Output: 403 Forbidden
}
