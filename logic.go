package keycloak

// The decision strategy dictates how the policies associated with a given policy are evaluated and how a final decision is obtained.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/Logic.java
const (
	// Defines that this policy follows a positive logic. In other words, the final decision is the policy outcome.
	LogicPositive = "POSITIVE"

	// Defines that this policy uses a logical negation. In other words, the final decision would be a negative of the policy outcome.
	LogicNegative = "NEGATIVE"
)
