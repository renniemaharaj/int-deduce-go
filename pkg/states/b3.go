package states

// States represents a three-state boolean: True, False, and Unset.
type States int8

const (
	Unset States = iota
)

// ToBool converts Tri to a standard boolean.
// - True  → true
// - False → false
// - Unset → defaults to false
func (t States) ToBool() bool {
	return t == True
}

// ToInt converts Tri to an integer representation.
// - True  → 1
// - False → 0
// - Unset → -1 (or another placeholder for unset state)
func (t States) ToInt() int {
	switch t {
	case True:
		return 1
	case False:
		return 0
	default:
		return -1
	}
}

// IsSet checks whether the value is explicitly True or False.
func (t States) IsSet() bool {
	return t != Unset
}

// Toggle flips the state between True and False.
// If Unset, it defaults to True.
func (t *States) Toggle() {
	switch *t {
	case True:
		*t = False
	case False:
		*t = True
	default:
		*t = True
	}
}

// Unset resets the value to the Unset state.
func (t *States) Unset() {
	*t = Unset
}
