package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validation interface methods defined
type Validation interface {
	Validate() (bool, map[string]string)
}

// Validator struct defined
type Validator struct {
	Errors map[string]string
}

// New validation method for Validator is defined
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// IsValid validatio nmethod defined
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

// MinLength validation method defined
func (v *Validator) MinLength(field, value string, maxChar int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < maxChar {
		v.Errors[field] = fmt.Sprintf("%s must be at least (%d) characters long", field, maxChar)
	}

	return true
}

// IsEmail validation method defined
func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}
	if !emailRegexp.MatchString(email) {
		v.Errors[field] = "not a valid email"
		return false
	}
	return true
}

// EqualToField validation method defined
func (v *Validator) EqualToField(field string, value interface{}, toEqualField string, toEqualValue interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if value != toEqualValue {
		v.Errors[field] = fmt.Sprintf("%s must equal %s", field, toEqualField)
		return false
	}

	return true

}

// Required validation method defined.
func (v *Validator) Required(field string, value interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if IsEmpty(value) {
		v.Errors[field] = fmt.Sprintf("%s is required", field)
		return false
	}

	return true

}

// IsEmpty function defiend. This uses REFLECT LIBRARY.
func IsEmpty(value interface{}) bool {
	t := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return t.Len() == 0

	}
	return false
}
