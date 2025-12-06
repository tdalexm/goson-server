package domain

import (
	"fmt"
	"strings"
)

type Filter struct {
	Field    string
	Value    string
	Contains string
}

func (f Filter) Matches(value any) (bool, error) {
	if f.Contains != "" {
		s, ok := value.(string)
		if !ok {
			return false, NewAppError(
				ErrFieldNotString,
				fmt.Sprintf("%s field value is not a string.", f.Field),
			)
		}

		return strings.Contains(strings.ToLower(s), strings.ToLower(f.Contains)), nil
	}

	return fmt.Sprintf("%v", value) == fmt.Sprintf("%v", f.Value), nil
}
