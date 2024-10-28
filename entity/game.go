package entity 
 
type Came struct{
	ID uint 
	CategoryID unit 
	QuestionIDs []uint
	Players []Players
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choice     PossibleAnswerChoice
}