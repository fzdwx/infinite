package components

import (
	"fmt"
)

// Validator is a function passed to a Question after a user has provided a response.
// If the function returns an error, then the user will be prompted again for another
// response.
type Validator func(ans interface{}) error

// MinItems requires that the list is longer or equal in length to the specified value
func MinItems(numberItems int) Validator {
	// return a validator that checks the length of the list
	return func(val interface{}) error {
		if list, ok := val.([]int); ok {
			// if the list is shorter than the given value
			if len(list) < numberItems {
				// yell loudly
				return fmt.Errorf("value is too short. Min items is %v", numberItems)
			}
		} else {
			// otherwise we cannot convert the value into a list of answer and cannot enforce length
			return fmt.Errorf("cannot impose the length on something other than a list of answers")
		}
		// the input is fine
		return nil
	}
}

// MaxItems requires that the list is no longer than the specified value
func MaxItems(numberItems int) Validator {
	// return a validator that checks the length of the list
	return func(val interface{}) error {
		if list, ok := val.([]int); ok {
			// if the list is longer than the given value
			if len(list) > numberItems {
				// yell loudly
				return fmt.Errorf("value is too long. Max items is %v", numberItems)
			}
		} else {
			// otherwise we cannot convert the value into a list of answer and cannot enforce length
			return fmt.Errorf("cannot impose the length on something other than a list of answers")
		}
		// the input is fine
		return nil
	}
}
