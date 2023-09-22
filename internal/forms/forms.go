package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	val := r.Form.Get(field)

	if val == "" {
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		val := f.Get(field)

		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	val := r.Form.Get(field)

	if len(val) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long.", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) bool {
	val := f.Get(field)

	if !govalidator.IsEmail(val) {
		f.Errors.Add(field, "Invalid email address entered.")
		return false
	}

	return true
}
