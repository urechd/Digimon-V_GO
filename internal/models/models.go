package models

type DigiType int
type DigiStage int

const (
	Data DigiType = iota + 1
	Vaccine
	Virus
	Free
)

const (
	InTraning DigiStage = iota + 1
	InTraningII
	Rookie
	Champion
	Ultimate
	Mega
)

type DigimonInfo struct {
	Id    int       `json:"Id"`
	Name  string    `json:"Name"`
	Type  DigiType  `json:"Type"`
	Stage DigiStage `json:"Stage"`
}

type BattleStats struct {
	MaxHP      int `json:"MaxHP"`
	CurrentHP  int `json:"CurrentHP"`
	Power      int `json:"Power"`
	NumBattles int `json:"NumBattles"`
	NumWins    int `json:"NumWins"`
}

type CaringInfo struct {
	Hunger          int `json:"Hunger"`
	Tiredness       int `json:"Tiredness"`
	Sleepiness      int `json:"Sleepiness"`
	Toilet          int `json:"Toilet"`
	NumCareMistakes int `json:"NumCareMistakes"`
}

type DigimonHistory struct {
	PreviousDigimon []int `json:"PreviousDigimon"`
	NextDigimon     []int `json:"NextDigimon"`
}

type DigivolutionConditions struct {
	BestDigimonId          int `json:"BestDigimonId"`
	HP                     int `json:"HP"`
	Power                  int `json:"Power"`
	NumCareMistakes        int `json:"NumCareMistakes"`
	BeatenStageId          int `json:"BeatenStageId"`
	NumWins                int `json:"NumWins"`
	NumBattles             int `json:"NumBattles"`
	NumMinConditionsNeeded int `json:"NumMinConditionsNeeded"`
}

type Digimon struct {
	Info           DigimonInfo            `json:"Info"`
	Battle         BattleStats            `json:"Battle"`
	Caring         CaringInfo             `json:"Caring"`
	History        DigimonHistory         `json:"History"`
	DigiConditions DigivolutionConditions `json:"DigiConditions"`
}

func (dt DigiType) String() string {
	return [...]string{"Data", "Vaccine", "Virus", "Free"}[dt-1]
}

func (ds DigiStage) String() string {
	return [...]string{"InTraning", "InTraningII", "Rookie", "Champion", "Ultimate", "Mega"}[ds-1]
}
