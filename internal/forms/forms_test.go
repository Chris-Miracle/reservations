package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormValid(t *testing.T){
	r := httptest.NewRequest("POST",  "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid")
	}
}


func TestFormRequired(t *testing.T){
	r := httptest.NewRequest("POST",  "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm  = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if  !form.Valid() {
		t.Error("Shows does not have required fields when it does")
	}
}

func TestFormHas(t *testing.T){
	r := httptest.NewRequest("POST",  "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("Form shows has field when it does not")
	}

	postedData :=  url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("Shows form does not have field when it should")
	}
}

func TestFormMinLength(t *testing.T){
	r := httptest.NewRequest("POST",  "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", "t")
	if form.Valid() {
		t.Error("form shows min length  for non-existing fields")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an arror but did not get one")
	}

	postedData :=  url.Values{}
	postedData.Add("first_name", "jas")
	postedData.Add("last_name", "a")

	form = New(postedData)
	form.MinLength("first_name", "last_name")
	if form.Valid() {
		t.Error("Form isn't properly checking minlength for fields")
	}

	postedData =  url.Values{}
	postedData.Add("first_name", "james")
	postedData.Add("last_name", "anais")

	form = New(postedData)
	form.MinLength("first_name", "last_name")
	if !form.Valid() {
		t.Error("Form isn't properly checking minlength for fields")
	}

	isError = form.Errors.Get("first_name")
	if isError != "" {
		t.Error("Should not have an arror but got one")
	}
}

func TestFormIsEmail(t *testing.T){
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmailValid("x")
	if form.Valid() {
		t.Error("Forms shows valid email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "james@gmail.com")

	form = New(postedData)
	form.IsEmailValid("email")

	if !form.Valid() {
		t.Error("Got an invalid email when we should not have")
	}

	postedData = url.Values{}
	postedData.Add("email", "james.com")

	form = New(postedData)
	form.IsEmailValid("email")

	if form.Valid() {
		t.Error("Got a valid email when we should not have")
	}
}

