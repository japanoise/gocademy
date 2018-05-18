package characters

type Character struct {
	GivenName     string
	Surname       string
	Gender        CGender
	Relationships map[Id]Relation
	ID            Id
	Personality   Id
	Loc           Location
	// Add like, hair and shit here
}

func (c *Character) GetDialogue(id Id) string {
	if Personalities[id] != nil {
		return Personalities[id].GetDialogue(id)
	} else {
		return Default.GetDialogue(id)
	}
}

func NewCharacter(forename, surname string, id Id) *Character {
	ret := &Character{}
	ret.GivenName = forename
	ret.Surname = surname
	return ret
}
