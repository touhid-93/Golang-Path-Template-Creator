package oldtests

import "fmt"

func GetAssertMessage(actual, expected interface{}, counter int) string {
	return fmt.Sprintf("----------------------\n%d )\tActual:%#v , Expected:%#v",
		counter,
		actual,
		expected,
	)
}
