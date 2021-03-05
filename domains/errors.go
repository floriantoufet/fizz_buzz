package domains

import (
	"strings"
)

type Errors []string

func (errs *Errors) Add(err error) {
	*errs = append(*errs, err.Error())
}

func (errs Errors) Error() string {
	var result string
	for _, err := range errs {
		result += err + ", "
	}

	return strings.TrimRight(result, ", ")
}

func (errs Errors) IsEmpty() bool {
	return len(errs) == 0
}

func (errs Errors) Contains(target error) bool {
	for i := 0; i < len(errs); i++ {
		if errs[i] == target.Error() {
			return true
		}
	}
	
	return false
}
