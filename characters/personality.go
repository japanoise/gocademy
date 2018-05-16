package characters

type Personality struct {
	Dialogue map[Id]string
}

func (p *Personality) GetDialogue(id Id) string {
	if p.Dialogue[id] != "" {
		return p.Dialogue[id]
	} else if Default.Dialogue[id] != "" {
		return Default.Dialogue[id]
	} else {
		return "!!! BAD DIALOGUE ID " + string(id) + " !!!"
	}
}
