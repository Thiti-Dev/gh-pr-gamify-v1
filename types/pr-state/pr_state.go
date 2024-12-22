package pr_state

import "fmt"

// PRState represents the state of a pull request.
type PRState string

// Predefined PR states.
const (
	PRStateApproved PRState = "approved"
	PRStateMerged   PRState = "merged"
	PRStateOpen     PRState = "open"
	PRStateClosed   PRState = "closed"
	PRStateUnknown  PRState = "unknown"
)

// NewPRState is a constructor that initializes a PRState and ensures it's valid.
func NewPRState(state string) (PRState, error) {
	prState := PRState(state)

	if !prState.IsValid() {
		return "", fmt.Errorf("invalid PR state: %s", state)
	}

	return prState, nil
}

// IsValid checks if the PRState is one of the predefined valid states.
func (s PRState) IsValid() bool {
	switch s {
	case PRStateApproved, PRStateMerged, PRStateOpen:
		return true
	default:
		return false
	}
}
