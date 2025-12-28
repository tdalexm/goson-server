package domain

import (
	"strconv"
	"strings"
)

const (
	FilterEquals   = "equals"
	FilterContains = "contains"
	FilterLT       = "lt"  // <
	FilterLTE      = "lte" // <=
	FilterGT       = "gt"  // >
	FilterGTE      = "gte" // >=
	FilterNE       = "ne"  // !=
)

type Filter struct {
	Field string
	Value string
	Type  string
}

func (f Filter) Matches(value string) (bool, error) {
	if f.Type == FilterGT || f.Type == FilterGTE || f.Type == FilterLT || f.Type == FilterLTE {
		filterFloat, filterIsNum := strconv.ParseFloat(f.Value, 32)
		fieldFloat, fieldIsNum := strconv.ParseFloat(value, 32)
		if fieldIsNum != nil || filterIsNum != nil {
			return false, NewAppError(ErrWrongParams, "Numeric operator for non numeric field")
		}
		switch f.Type {
		case FilterLT:
			return fieldFloat < filterFloat, nil
		case FilterLTE:
			return fieldFloat <= filterFloat, nil
		case FilterGT:
			return fieldFloat > filterFloat, nil
		case FilterGTE:
			return fieldFloat >= filterFloat, nil
		case FilterNE:
			return fieldFloat != filterFloat, nil
		}
	}

	switch f.Type {
	case FilterNE:
		return !strings.EqualFold(value, f.Value), nil
	case FilterContains:
		return strings.Contains(strings.ToLower(value), strings.ToLower(f.Value)), nil
	default:
		return strings.EqualFold(value, f.Value), nil
	}
}
