package keycloak

// DecisionStrategy dictates how the policies associated with a given policy are evaluated and how a final decision is obtained.
//
// https://github.com/keycloak/keycloak/blob/master/core/src/main/java/org/keycloak/representations/idm/authorization/DecisionStrategy.java
const (
	// AFFIRMATIVE defines that at least one policy must evaluate to a positive decision
	// in order to the overall decision be also positive.
	DecisionStrategyAffirmative = "AFFIRMATIVE"

	// UNANIMOUS defines that all policies must evaluate to a positive decision
	// in order to the overall decision be also positive.
	DecisionStrategyUnanimous = "UNANIMOUS"

	// CONSENSUS defines that the number of positive decisions must be greater than the number of negative decisions.
	// If the number of positive and negative is the same, the final decision will be negative.
	DecisionStrategyConsensus = "CONSENSUS"
)
