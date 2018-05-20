package characters

type RLevel int8

type Relation struct {
	Intensity RLevel
	Love      RLevel
}

const (
	FriendThreshold RLevel = 10
)
