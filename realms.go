package keycloak

import (
	"context"
	"fmt"
	"net/http"
)

// IdentityProvider representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/IdentityProviderRepresentation.java
type IdentityProvider struct {
}

// IdentityProviderMapper representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/IdentityProviderMapperRepresentation.java
type IdentityProviderMapper struct {
}

// Realm representation.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/RealmRepresentation.java
type Realm struct {
	ID                                                        *string                   `json:"id,omitempty"`
	Realm                                                     *string                   `json:"realm,omitempty"`
	DisplayName                                               *string                   `json:"displayName,omitempty"`
	DisplayNameHTML                                           *string                   `json:"displayNameHtml,omitempty"`
	NotBefore                                                 *int                      `json:"notBefore,omitempty"`
	RevokeRefreshToken                                        *bool                     `json:"revokeRefreshToken,omitempty"`
	RefreshTokenMaxReuse                                      *int                      `json:"refreshTokenMaxReuse,omitempty"`
	AccessTokenLifespan                                       *int                      `json:"accessTokenLifespan,omitempty"`
	AccessTokenLifespanForImplicitFlow                        *int                      `json:"accessTokenLifespanForImplicitFlow,omitempty"`
	SsoSessionIdleTimeout                                     *int                      `json:"ssoSessionIdleTimeout,omitempty"`
	SsoSessionMaxLifespan                                     *int                      `json:"ssoSessionMaxLifespan,omitempty"`
	SsoSessionIdleTimeoutRememberMe                           *int                      `json:"ssoSessionIdleTimeoutRememberMe,omitempty"`
	SsoSessionMaxLifespanRememberMe                           *int                      `json:"ssoSessionMaxLifespanRememberMe,omitempty"`
	OfflineSessionIdleTimeout                                 *int                      `json:"offlineSessionIdleTimeout,omitempty"`
	OfflineSessionMaxLifespanEnabled                          *bool                     `json:"offlineSessionMaxLifespanEnabled,omitempty"`
	OfflineSessionMaxLifespan                                 *int                      `json:"offlineSessionMaxLifespan,omitempty"`
	ClientSessionIdleTimeout                                  *int                      `json:"clientSessionIdleTimeout,omitempty"`
	ClientSessionMaxLifespan                                  *int                      `json:"clientSessionMaxLifespan,omitempty"`
	ClientOfflineSessionIdleTimeout                           *int                      `json:"clientOfflineSessionIdleTimeout,omitempty"`
	ClientOfflineSessionMaxLifespan                           *int                      `json:"clientOfflineSessionMaxLifespan,omitempty"`
	AccessCodeLifespan                                        *int                      `json:"accessCodeLifespan,omitempty"`
	AccessCodeLifespanUserAction                              *int                      `json:"accessCodeLifespanUserAction,omitempty"`
	AccessCodeLifespanLogin                                   *int                      `json:"accessCodeLifespanLogin,omitempty"`
	ActionTokenGeneratedByAdminLifespan                       *int                      `json:"actionTokenGeneratedByAdminLifespan,omitempty"`
	ActionTokenGeneratedByUserLifespan                        *int                      `json:"actionTokenGeneratedByUserLifespan,omitempty"`
	Enabled                                                   *bool                     `json:"enabled,omitempty"`
	SslRequired                                               *string                   `json:"sslRequired,omitempty"`
	RegistrationAllowed                                       *bool                     `json:"registrationAllowed,omitempty"`
	RegistrationEmailAsUsername                               *bool                     `json:"registrationEmailAsUsername,omitempty"`
	RememberMe                                                *bool                     `json:"rememberMe,omitempty"`
	VerifyEmail                                               *bool                     `json:"verifyEmail,omitempty"`
	LoginWithEmailAllowed                                     *bool                     `json:"loginWithEmailAllowed,omitempty"`
	DuplicateEmailsAllowed                                    *bool                     `json:"duplicateEmailsAllowed,omitempty"`
	ResetPasswordAllowed                                      *bool                     `json:"resetPasswordAllowed,omitempty"`
	EditUsernameAllowed                                       *bool                     `json:"editUsernameAllowed,omitempty"`
	BruteForceProtected                                       *bool                     `json:"bruteForceProtected,omitempty"`
	PermanentLockout                                          *bool                     `json:"permanentLockout,omitempty"`
	MaxFailureWaitSeconds                                     *int                      `json:"maxFailureWaitSeconds,omitempty"`
	MinimumQuickLoginWaitSeconds                              *int                      `json:"minimumQuickLoginWaitSeconds,omitempty"`
	WaitIncrementSeconds                                      *int                      `json:"waitIncrementSeconds,omitempty"`
	QuickLoginCheckMilliSeconds                               *int                      `json:"quickLoginCheckMilliSeconds,omitempty"`
	MaxDeltaTimeSeconds                                       *int                      `json:"maxDeltaTimeSeconds,omitempty"`
	FailureFactor                                             *int                      `json:"failureFactor,omitempty"`
	DefaultRoles                                              []string                  `json:"defaultRoles,omitempty"`
	RequiredCredentials                                       []string                  `json:"requiredCredentials,omitempty"`
	OtpPolicyType                                             *string                   `json:"otpPolicyType,omitempty"`
	OtpPolicyAlgorithm                                        *string                   `json:"otpPolicyAlgorithm,omitempty"`
	OtpPolicyInitialCounter                                   *int                      `json:"otpPolicyInitialCounter,omitempty"`
	OtpPolicyDigits                                           *int                      `json:"otpPolicyDigits,omitempty"`
	OtpPolicyLookAheadWindow                                  *int                      `json:"otpPolicyLookAheadWindow,omitempty"`
	OtpPolicyPeriod                                           *int                      `json:"otpPolicyPeriod,omitempty"`
	OtpSupportedApplications                                  []string                  `json:"otpSupportedApplications,omitempty"`
	WebAuthnPolicyRpEntityName                                *string                   `json:"webAuthnPolicyRpEntityName,omitempty"`
	WebAuthnPolicySignatureAlgorithms                         []string                  `json:"webAuthnPolicySignatureAlgorithms,omitempty"`
	WebAuthnPolicyRpID                                        *string                   `json:"webAuthnPolicyRpId,omitempty"`
	WebAuthnPolicyAttestationConveyancePreference             *string                   `json:"webAuthnPolicyAttestationConveyancePreference,omitempty"`
	WebAuthnPolicyAuthenticatorAttachment                     *string                   `json:"webAuthnPolicyAuthenticatorAttachment,omitempty"`
	WebAuthnPolicyRequireResidentKey                          *string                   `json:"webAuthnPolicyRequireResidentKey,omitempty"`
	WebAuthnPolicyUserVerificationRequirement                 *string                   `json:"webAuthnPolicyUserVerificationRequirement,omitempty"`
	WebAuthnPolicyCreateTimeout                               *int                      `json:"webAuthnPolicyCreateTimeout,omitempty"`
	WebAuthnPolicyAvoidSameAuthenticatorRegister              *bool                     `json:"webAuthnPolicyAvoidSameAuthenticatorRegister,omitempty"`
	WebAuthnPolicyAcceptableAaguids                           []string                  `json:"webAuthnPolicyAcceptableAaguids,omitempty"`
	WebAuthnPolicyPasswordlessRpEntityName                    *string                   `json:"webAuthnPolicyPasswordlessRpEntityName,omitempty"`
	WebAuthnPolicyPasswordlessSignatureAlgorithms             []string                  `json:"webAuthnPolicyPasswordlessSignatureAlgorithms,omitempty"`
	WebAuthnPolicyPasswordlessRpID                            *string                   `json:"webAuthnPolicyPasswordlessRpId,omitempty"`
	WebAuthnPolicyPasswordlessAttestationConveyancePreference *string                   `json:"webAuthnPolicyPasswordlessAttestationConveyancePreference,omitempty"`
	WebAuthnPolicyPasswordlessAuthenticatorAttachment         *string                   `json:"webAuthnPolicyPasswordlessAuthenticatorAttachment,omitempty"`
	WebAuthnPolicyPasswordlessRequireResidentKey              *string                   `json:"webAuthnPolicyPasswordlessRequireResidentKey,omitempty"`
	WebAuthnPolicyPasswordlessUserVerificationRequirement     *string                   `json:"webAuthnPolicyPasswordlessUserVerificationRequirement,omitempty"`
	WebAuthnPolicyPasswordlessCreateTimeout                   *int                      `json:"webAuthnPolicyPasswordlessCreateTimeout,omitempty"`
	WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister  *bool                     `json:"webAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister,omitempty"`
	WebAuthnPolicyPasswordlessAcceptableAaguids               []string                  `json:"webAuthnPolicyPasswordlessAcceptableAaguids,omitempty"`
	BrowserSecurityHeaders                                    *map[string]string        `json:"browserSecurityHeaders,omitempty"`
	SMTPServer                                                *map[string]string        `json:"smtpServer,omitempty"`
	EventsEnabled                                             *bool                     `json:"eventsEnabled,omitempty"`
	EventsListeners                                           []string                  `json:"eventsListeners,omitempty"`
	EnabledEventTypes                                         []string                  `json:"enabledEventTypes,omitempty"`
	AdminEventsEnabled                                        *bool                     `json:"adminEventsEnabled,omitempty"`
	AdminEventsDetailsEnabled                                 *bool                     `json:"adminEventsDetailsEnabled,omitempty"`
	IdentityProviders                                         []*IdentityProvider       `json:"identityProviders,omitempty"`
	IdentityProviderMappers                                   []*IdentityProviderMapper `json:"identityProviderMappers,omitempty"`
	InternationalizationEnabled                               *bool                     `json:"internationalizationEnabled,omitempty"`
	SupportedLocales                                          []string                  `json:"supportedLocales,omitempty"`
	BrowserFlow                                               *string                   `json:"browserFlow,omitempty"`
	RegistrationFlow                                          *string                   `json:"registrationFlow,omitempty"`
	DirectGrantFlow                                           *string                   `json:"directGrantFlow,omitempty"`
	ResetCredentialsFlow                                      *string                   `json:"resetCredentialsFlow,omitempty"`
	ClientAuthenticationFlow                                  *string                   `json:"clientAuthenticationFlow,omitempty"`
	DockerAuthenticationFlow                                  *string                   `json:"dockerAuthenticationFlow,omitempty"`
	Attributes                                                *map[string]string        `json:"attributes,omitempty"`
	UserManagedAccessAllowed                                  *bool                     `json:"userManagedAccessAllowed,omitempty"`
}

// Configuration represents a UMA configuration.
type Configuration struct {
	Issuer                                     *string  `json:"issuer,omitempty"`
	AuthorizationEndpoint                      *string  `json:"authorization_endpoint,omitempty"`
	TokenEndpoint                              *string  `json:"token_endpoint,omitempty"`
	IntrospectionEndpoint                      *string  `json:"introspection_endpoint,omitempty"`
	EndSessionEndpoint                         *string  `json:"end_session_endpoint,omitempty"`
	JwksURI                                    *string  `json:"jwks_uri,omitempty"`
	GrantTypesSupported                        []string `json:"grant_types_supported,omitempty"`
	ResponseTypesSupported                     []string `json:"response_types_supported,omitempty"`
	ResponseModesSupported                     []string `json:"response_modes_supported,omitempty"`
	RegistrationEndpoint                       *string  `json:"registration_endpoint,omitempty"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	ScopesSupported                            []string `json:"scopes_supported,omitempty"`
	ResourceRegistrationEndpoint               *string  `json:"resource_registration_endpoint,omitempty"`
	PermissionEndpoint                         *string  `json:"permission_endpoint,omitempty"`
	PolicyEndpoint                             *string  `json:"policy_endpoint,omitempty"`
}

// RealmsService ...
type RealmsService service

// Create a new realm.
func (s *RealmsService) Create(ctx context.Context, realm *Realm) (*http.Response, error) {
	req, err := s.keycloak.NewRequest(http.MethodPost, "admin/realms", realm)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// List all realms.
func (s *RealmsService) List(ctx context.Context) ([]*Realm, *http.Response, error) {
	req, err := s.keycloak.NewRequest(http.MethodGet, "admin/realms", nil)
	if err != nil {
		return nil, nil, err
	}

	var realms []*Realm
	res, err := s.keycloak.Do(ctx, req, &realms)
	if err != nil {
		return nil, nil, err
	}

	return realms, res, nil
}

// Get realm.
func (s *RealmsService) Get(ctx context.Context, name string) (*Realm, *http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s", name)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var realm Realm
	res, err := s.keycloak.Do(ctx, req, &realm)
	if err != nil {
		return nil, nil, err
	}

	return &realm, res, nil
}

// Delete realm.
func (s *RealmsService) Delete(ctx context.Context, name string) (*http.Response, error) {
	u := fmt.Sprintf("admin/realms/%s", name)
	req, err := s.keycloak.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.keycloak.Do(ctx, req, nil)
}

// GetConfig gets realm configuration.
func (s *RealmsService) GetConfig(ctx context.Context, name string) (*Configuration, *http.Response, error) {
	u := fmt.Sprintf("realms/%s/.well-known/uma2-configuration", name)
	req, err := s.keycloak.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var config Configuration
	res, err := s.keycloak.Do(ctx, req, &config)
	if err != nil {
		return nil, nil, err
	}

	return &config, res, nil
}
