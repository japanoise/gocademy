package characters

type Character struct {
	GivenName     string
	Surname       string
	Relationships map[Id]Relation
	ID            Id
	Personality   Id
	// Add like, hair and shit here
}

func (c *Character) GetDialogue(id Id) string {
	if Personalities[id] != nil {
		return Personalities[id].GetDialogue(id)
	} else {
		return Default.GetDialogue(id)
	}
}
