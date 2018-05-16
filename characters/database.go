package characters

type Id string

var (
	Default       *Personality
	Personalities map[Id]*Personality
)

func init() {
	Default = &Personality{}
	Default.Dialogue = make(map[Id]string)
	Personalities = make(map[Id]*Personality)
}
