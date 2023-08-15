package personafactory

import (
	"math/rand"

	"time"

	"github.com/burakgulbay/personafactory/structs"
)

type Creator interface {
	Create() *structs.Person
}

type PersonCreator struct {
	randomFirstNameIndex int
	randomLastNameIndex  int
	randomJobIndex       int
	gender               string
}

func (pc PersonCreator) Create() *structs.Person {

	person := &structs.Person{}

	if pc.gender == "Male" {
		person.FirstName = MaleFirstNames[pc.randomFirstNameIndex]
	} else {
		person.FirstName = FemaleFirstNames[pc.randomFirstNameIndex]
	}

	person.LastName = LastNames[pc.randomLastNameIndex]
	person.Job = Jobs[pc.randomJobIndex]
	person.Gender = pc.gender
	person.BirthDate = RandomDate()

	return person
}

func RandomDate() string {
	min := time.Date(1930, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0)
	return t.Format("2006-01-02")
}

func CreatePersonCreator() Creator {

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	gender := randomizer.Intn(2) == 1

	var personCreator PersonCreator

	if gender { // male
		personCreator = PersonCreator{
			randomFirstNameIndex: randomizer.Intn(len(MaleFirstNames)),
			randomLastNameIndex:  randomizer.Intn(len(LastNames)),
			randomJobIndex:       randomizer.Intn(len(Jobs)),
			gender:               "Male",
		}
	} else { // female
		personCreator = PersonCreator{
			randomFirstNameIndex: randomizer.Intn(len(FemaleFirstNames)),
			randomLastNameIndex:  rand.Intn(len(LastNames)),
			randomJobIndex:       rand.Intn(len(Jobs)),
			gender:               "Female",
		}
	}
	return personCreator
}
