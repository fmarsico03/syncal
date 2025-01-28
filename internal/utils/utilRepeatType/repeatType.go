package utilRepeatType

type RepeatType string

const (
	OneTime RepeatType = "OneTime"
	Daily   RepeatType = "Daily"
	Weekly  RepeatType = "Weekly"
	Monthly RepeatType = "Monthly"
)

// Repeat representa una configuración de repetición
type Repeat struct {
	TypeOf RepeatType `json:"type_of"`
	Value  int        `json:"value"`
}

func (r *Repeat) IsValid() bool {
	switch r.TypeOf {
	case OneTime, Daily, Weekly, Monthly:
		return r.Value > 0
	default:
		return false
	}
}
