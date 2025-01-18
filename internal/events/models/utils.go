package models

type Recurrence struct {
	Days   []DayOfWeek
	Weeks  []int
	Months []int
}

// Declarar el tipo de enum
type DayOfWeek int

// Definir los valores del enum utilizando iota
const (
	Sunday    DayOfWeek = iota // 0
	Monday                     // 1
	Tuesday                    // 2
	Wednesday                  // 3
	Thursday                   // 4
	Friday                     // 5
	Saturday                   // 6
)

func ConvertToDaysOfWeek(days []int) []DayOfWeek {
	var daysOfWeek []DayOfWeek

	for _, day := range days {
		daysOfWeek = append(daysOfWeek, DayOfWeek(day))
	}
	return daysOfWeek
}
