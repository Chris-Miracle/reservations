package forms

import (
	"net/http"
	"net/mail"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}


// Valid returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Fields required
func (f *Form) Required(fields ...string){
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Required checks for required fields
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Email Validity
func (f *Form) IsEmailValid(field string) {
	value := f.Get(field)
	if field == "email" {
		_, err  := mail.ParseAddress(value)
		if err != nil {
			f.Errors.Add(field, "Email is invalid")
		}
	}
}

func (f *Form) MinLength(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if len(strings.TrimSpace(value)) < 3 {
			f.Errors.Add(field, "This field can't have less than 3 chars")
		}
	}
}