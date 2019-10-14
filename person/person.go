package person

import "errors"

// Person used as ds to store empid, name
type Person struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
}

var people []Person

func main() {
	people = []Person{}
}

//GetPersonByID - retrieve a person by their id
func GetPersonByID(id int64) (Person, error) {
	for _, p := range people {
		if p.ID == id {
			return p, nil
		}
	}
	var p Person
	return p, errors.New("No Person found with this ID")
}

//GetPeople - gets all persons in db
func GetPeople() []Person {
	return people
}

//CreatePerson - add a new person to list
func CreatePerson(p Person) Person {
	var id int64
	if len(people) > 0 {
		id = people[len(people)-1].ID
	}
	id++

	people = append(people, p)
	return p
}

//DeletePersonByID - remove an existing person from list
func DeletePersonByID(id int64) (Person, error) {
	for i, p := range people {
		if p.ID == id {
			people = append(people[:i], people[i+1:]...)
			return p, nil
		}
	}
	var p Person
	return p, errors.New("No Person found with this ID")
}

//UpdatePerson - update name of person
func UpdatePerson(person Person) (Person, error) {
	for _, p := range people {
		if p.ID == person.ID {
			p.Name = person.Name
			return p, nil
		}
	}
	var p Person
	return p, errors.New("No Person found with this ID")
}
