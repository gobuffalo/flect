package flect

import "github.com/gobuffalo/flect/internal/core"

// Humanize returns first letter of sentence capitalized.
// Common acronyms are capitalized as well.
// Other capital letters in string are left as provided.
//
//	employee_salary = Employee salary
//	employee_id = employee ID
//	employee_mobile_number = Employee mobile number
//	first_Name = First Name
//	firstName = First Name
func Humanize(s string) string {
	return core.Humanize(s)
}
