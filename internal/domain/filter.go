package domain

import (
	"fmt"
	"strings"
)

type Filter struct {
	Field string
	Value string
	Type  string
}

func (f Filter) Matches(value any) bool {
	if f.Type == "contains" {
		if s, ok := value.(string); ok {
			return strings.Contains(strings.ToLower(s), strings.ToLower(f.Value))
		}
	}

	if s, ok := value.(string); ok {
		return strings.Contains(strings.ToLower(s), strings.ToLower(f.Value))
	}

	return fmt.Sprintf("%v", value) == fmt.Sprintf("%v", f.Value)
}
