package keycloak_test

import (
	"context"
	"fmt"

	"github.com/zemirco/keycloak"
	"golang.org/x/oauth2"
)

func ExampleNewKeycloak_admin() {
	// create a new oauth2 configuration
	config := oauth2.Config{
		ClientID: "admin-cli",
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/auth/realms/master/protocol/openid-connect/token",
		},
	}

	ctx := context.Background()

	// this should be your KEYCLOAK_USER and your KEYCLOAK_PASSWORD
	token, err := config.PasswordCredentialsToken(ctx, "admin", "admin")
	if err != nil {
		panic(err)
	}

	// create a new http client
	httpClient := config.Client(ctx, token)

	// use the http client to create a Keycloak instance
	kc, err := keycloak.NewKeycloak(httpClient, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	// then use this instance to make requests to the API
	fmt.Println(kc)
}

func ExampleNewKeycloak_user() {
	// create a new oauth2 configuration
	config := oauth2.Config{
		ClientID:     "40349713-a521-48b2-9197-216adfce5f78",
		ClientSecret: "ddafbeba-1402-4969-bbbb-475d657fa9d5",
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("http://localhost:8080/auth/realms/%s/protocol/openid-connect/token", "myrealm"),
		},
	}

	ctx := context.Background()

	// this should be your username and your password
	token, err := config.PasswordCredentialsToken(ctx, "user", "user_password")
	if err != nil {
		panic(err)
	}

	// create a new http client
	httpClient := config.Client(ctx, token)

	// use the http client to create a Keycloak instance
	kc, err := keycloak.NewKeycloak(httpClient, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	// then use this instance to make requests to the API
	fmt.Println(kc)
}

func ExampleRealmsService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	realm := &keycloak.Realm{
		Enabled: keycloak.Bool(true),
		ID:      keycloak.String("myrealm"),
		Realm:   keycloak.String("myrealm"),
	}

	ctx := context.Background()

	res, err := kc.Realms.Create(ctx, realm)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleUsersService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	user := &keycloak.User{
		Enabled:   keycloak.Bool(true),
		Username:  keycloak.String("username"),
		Email:     keycloak.String("user@email.com"),
		FirstName: keycloak.String("first"),
		LastName:  keycloak.String("last"),
	}

	ctx := context.Background()
	res, err := kc.Users.Create(ctx, "myrealm", user)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleClientsService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	client := &keycloak.Client{
		Enabled:  keycloak.Bool(true),
		ClientID: keycloak.String("myclient"),
	}

	ctx := context.Background()
	res, err := kc.Clients.Create(ctx, "myrealm", client)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleScopesService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	scope := &keycloak.Scope{
		Name: keycloak.String("scope_read"),
	}

	ctx := context.Background()

	clientID := "40349713-a521-48b2-9197-216adfce5f78"
	scope, res, err := kc.Scopes.Create(ctx, "myrealm", clientID, scope)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleRolesService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	role := &keycloak.Role{
		Name:        keycloak.String("my name"),
		Description: keycloak.String("my description"),
	}

	ctx := context.Background()
	res, err := kc.Roles.Create(ctx, "myrealm", role)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleResourcesService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	resource := &keycloak.Resource{
		Name:        keycloak.String("resource"),
		DisplayName: keycloak.String("resource"),
	}

	ctx := context.Background()

	clientID := "40349713-a521-48b2-9197-216adfce5f78"
	resource, res, err := kc.Resources.Create(ctx, "myrealm", clientID, resource)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExamplePoliciesService_CreateUserPolicy() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	policy := &keycloak.UserPolicy{
		Policy: keycloak.Policy{
			Type:             keycloak.String("user"),
			Logic:            keycloak.String(keycloak.LogicPositive),
			DecisionStrategy: keycloak.String(keycloak.DecisionStrategyUnanimous),
			Name:             keycloak.String("policy"),
		},
		Users: []string{"89400c55-15fd-4e5d-a7c9-403f431d97f3"},
	}

	ctx := context.Background()

	clientID := "40349713-a521-48b2-9197-216adfce5f78"
	policy, res, err := kc.Policies.CreateUserPolicy(ctx, "myrealm", clientID, policy)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func ExampleGroupsService_Create() {
	kc, err := keycloak.NewKeycloak(nil, "http://localhost:8080/auth/")
	if err != nil {
		panic(err)
	}

	group := &keycloak.Group{
		Name: keycloak.String("mygroup"),
	}

	ctx := context.Background()
	res, err := kc.Groups.Create(ctx, "myrealm", group)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
