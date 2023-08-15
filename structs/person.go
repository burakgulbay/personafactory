package structs

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Job       string
	BirthDate string // dd/mm/yyyy
	Gender    string
}

func (p Person) PrintJSON() {
	marshalledPerson, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(marshalledPerson))
}
