package gordle

type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

// feedback is a list of hints one per character of the word.
type feedback []hint
