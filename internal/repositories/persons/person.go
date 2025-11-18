package persons

import (
	"task-api/internal/models"

	"github.com/google/uuid"
)

type Persons struct {
	persons []models.Person
}

func New() *Persons {
	return &Persons{persons: make([]models.Person, 0)}
}

func (p Persons) GetPerson(id uuid.UUID) *models.Person {
	for i := range p.persons {
		if p.persons[i].Id == id {
			return &p.persons[i]
		}
	}
	return nil
}

func (p Persons) GetAllPerson() []models.Person {
	return p.persons
}

func (p *Persons) AddPerson(newPerson models.Person) {
	p.persons = append(p.persons, newPerson)
}
