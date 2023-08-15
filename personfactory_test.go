package personafactory

import (
	"testing"
)

func TestPersonCreatorCreate(t *testing.T) {
	personCreator := CreatePersonCreator()
	p := personCreator.Create()

	if p.FirstName == "" {
		t.Fatal("person first name is empty")
	}
	if p.LastName == "" {
		t.Fatal("person last name is empty")
	}
	if p.Gender == "" {
		t.Fatal("person gender is empty")
	}
	if p.Job == "" {
		t.Fatal("person job is empty")
	}
	if p.BirthDate == "" {
		t.Fatal("person birth date is empty")
	}

	p.PrintJSON()
}
