package main

import (
	"fmt"

	"github.com/japanoise/gocademy/characters"
	"github.com/japanoise/termbox-util"
)

func interactDraw(player, target *characters.Character) func(int, int, int) {
	name := target.GetNameString()
	return func(_, sx, sy int) {
		pronouns := pronounStrings[target.Gender]
		termutil.Printstring("Name:", sx-5, 1)
		termutil.Printstring(name, sx-(len(name)), 2)
		termutil.Printstring("Pronouns:", sx-9, 3)
		termutil.Printstring(pronouns, sx-(len(pronouns)), 4)
		target.Face.DrawAscii(sx-12, 5)
	}
}

func Interact(player, target *characters.Character) string {
	var greeting string
	if target.Relationships == nil {
		target.Relationships = make(map[characters.Id]*characters.Relation)
	}
	relation := target.Relationships[player.ID]
	if relation == nil {
		relation = &characters.Relation{0, 0}
		target.Relationships[player.ID] = relation
		greeting = target.SayDialogue(characters.DialogueGreetingFirst)
	} else if relation.Intensity > characters.FriendThreshold {
		greeting = target.SayDialogue(characters.DialogueGreetingFriend)
	} else {
		greeting = target.SayDialogue(characters.DialogueGreetingStranger)
	}
	choice := termutil.ChoiceIndexCallback(greeting, []string{
		"How are you?",
		"How do you feel about...?",
		"<Talk about studying>",
		"<Talk about clubs>",
		"<Talk about love>",
		"<Talk about sex>",
		"[DEBUG] Tell me what your path is.",
		"Bye.",
	}, 6, interactDraw(player, target))
	switch choice {
	case 0:
		return target.SayDialogue(characters.DialogueNo)
	case 1:
		return target.SayDialogue(characters.DialogueNo)
	case 2:
		return target.SayDialogue(characters.DialogueNo)
	case 3:
		return target.SayDialogue(characters.DialogueNo)
	case 4:
		return target.SayDialogue(characters.DialogueNo)
	case 5:
		return target.SayDialogue(characters.DialogueNo)
	case 6:
		return fmt.Sprint(target.Path)
	default:
		return target.SayDialogue(characters.DialogueBye)
	}
}
