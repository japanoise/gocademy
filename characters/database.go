package characters

type Id string

const (
	DialogueGreetingFirst    Id = "dialogue-greeting-first"
	DialogueGreetingStranger    = "dialogue-greeting-stranger"
	DialogueGreetingFriend      = "dialogue-greeting-friend"
	DialogueYes                 = "dialogue-yes"
	DialogueNo                  = "dialogue-no"
	DialogueBye                 = "dialogue-bye"
)

var (
	Default       *Personality
	Personalities map[Id]*Personality
)

func init() {
	Default = &Personality{}
	Default.Dialogue = make(map[Id]string)
	Default.Dialogue[DialogueGreetingFirst] = "A pleasure to meet you."
	Default.Dialogue[DialogueGreetingStranger] = "Can I help you?"
	Default.Dialogue[DialogueGreetingFriend] = "Hey there, buddy!"
	Default.Dialogue[DialogueYes] = "Sure thing!"
	Default.Dialogue[DialogueNo] = "No way."
	Default.Dialogue[DialogueBye] = "See ya."
	Personalities = make(map[Id]*Personality)
}
