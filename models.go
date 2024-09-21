package v1

import (
	"fmt"
)

type Invitation struct {
	Bin        string `json:"bin"`
	GameGroup  string `json:"game_group"`
	CreatorId  string `json:"creator_id"`
	Status     string `json:"status"`
	Invitation string `json:"invitation"` // type - game/group/message etc.
	Message    string `json:"message"`
	Time       string `json:"time"`
	GameId     string `json:"game_id"`
	Accepted   bool   `json:"accepted"`
	Declined   bool   `json:"declined"`
}

type Player struct {
	UserName    string                 `json:"user_name"`
	Bin         string                 `json:"bin"`
	Photo       string                 `json:"photo"`
	Status      string                 `json:"status"` //inAGame, afk, available
	Privacy     string                 `json:"privacy"`
	Invitations map[string]*Invitation `json:"invitations",omitempty`
	GroupIds    []string               `json:"group_ids",omitempty`
	GameIds     []string               `json:"game_ids",omitempty`
}

type Group struct {
	Bin       string    `json:"bin"`
	Creator   *Player   `json:"creator"`
	Members   []*Player `json:"members", omitempty`
	GroupName string    `json:"group_name"`
	Capacity  int       `json:"capacity"`
	Status    string    `json:"status"`
}

type Game struct {
	Bin               string                `json:"bin"`
	IsDaytime         bool                  `json:"is_daytime", omitempty`
	ExplanationSeen   bool                  `json:"explanation_seen", omitempty`
	FirstDayCompleted bool                  `json:"first_day_completed", omitempty`
	CurrentStep       string                `json:"current_step"`
	GroupId           string                `json:"group_id"`
	Status            string                `json:"status"`
	ServerName        string                `json:"server_name", omitempty`
	ServerAddress     string                `json:"server_ip", omitempty`
	ServerPort        int                   `json:"server_port", omitempty`
	StartTime         string                `json:"start_time", omitempty`
	EndTime           string                `json:"end_time", omitempty`
	Creator           *Player               `json:"creator",  omitempty`
	Cycles            int                   `json:"cycles", omitempty`
	Invited           []string              `json:"invited", omitempty`
	Steps             map[string]*Step      `json:"steps"omitempty`
	Gamers            map[string]*Gamer     `json:"gamers", omitempty`
	Characters        map[string]*Character `json:"characters" omitempty`
}

type Step struct {
	Bin          string   `json:"bin"`
	NextStep     string   `json:"next_step"`
	Type         string   `json:"type"`
	StepType     string   `json:"step_type"`
	Duration     string   `json:"duration"`
	Command      string   `json:"command"`
	NcCommand    string   `json:"nc_command"`
	ShowTimer    bool     `json:"show_timer"`
	CharPhotoUrl string   `json:"char_photo_url"`
	Text         string   `json:"text"`
	NcText       string   `json:"nc_text"`
	Character    string   `json:"character"`
	StepIndex    int      `json:"step_index"`
	StartTime    string   `json:"start_time"`
	EndTime      string   `json:"end_time"`
	Result       *Result  `json:"result"`
	Allowed      []string `json:"allowed"`
}

type Result struct {
	Bin       string            `json:"bin"`
	Index     int               `json:"index"`                     // order
	EndTime   string            `json:"end_time"`                  // milliseconds
	ResultMap map[string]*Gamer `json:"players_affected"omitempty` // poisoned: [player1, player2] // healed : [player1]
	TimeStamp string            `json:"timestamp"`
}

type Gamer struct {
	Bin         string             `json:"bin"`
	GameId      string             `json:"game_id"`
	CharacterId string             `json:"character_id"`
	Name        string             `json:"name"`
	ImageUrl    string             `json:"image_url"`
	IsAlive     bool               `json:"is_alive"`
	Actions     map[string]*Action `json:"actions"omitempty` // map of string (stepId) and Action
	Abilities   []*Ability         `json:"abilities"`
}

type Action struct {
	Bin      string              `json:"bin"`
	Vote     Vote                `json:"vote"omitempty`
	Messages map[string]*Message `json:"messages"omitempty` // map of string (timestamp millis) and message
}

type Vote struct {
	Bin       string `json:"bin"`
	StepId    string `json:"step_id"`
	Target    string `json:"target"`
	TimeStamp string `json:"time_stamp"`
	Decision  string `json:"decision"`
}

type Message struct {
	Bin       string `json:"bin"`
	StepId    string `json:"step_id"`
	PlayerId  string `json:"player_id"`
	Text      string `json:"text"`
	Icon      string `json:"icon"`
	HasImage  bool   `json:"has_image"`
	ImagePath string `json:"image_path"omitempty`
	AudioPath string `json:"audio_path"omitempty`
	VideoPath string `json:"video_path"omitempty`
	TimeStamp string `json:"time_stamp"omitempty`
}

type CharacterPack struct {
	Name       string       `json:"name"`
	OriginDate string       `json:"origin_date"`
	Creator    string       `json:"creator"`
	Characters []*Character `json:"characters"omitempty`
}

type Character struct {
	Bin         string   `json:"bin"`
	TypeId      int      `json:"type_id"`
	Name        string   `json:"name"`
	ImageUrl    string   `json:"image_url"`
	Description string   `json:"description"`
	Abilities   []string `json:"abilities"omitempty`
	Role        string   `json:"role"`
}

type Ability struct {
	Bin       string `json:"bin"`
	Name      string `json:"name"`
	Character string `json:"character"`
	Frequency string `json:"frequency"`
	TimesUsed int    `json:"times_used"omitempty"`
}

// RULES ENGINE
type RulesEngine struct {
	game *Game
}

// NewRulesEngine creates a new RulesEngine instance.
func NewRulesEngine(game *Game) (*RulesEngine, error) {
	return &RulesEngine{game: game}, nil
}

// GetAllowedAbilities determines the allowed abilities for a player in the current step.
func (re *RulesEngine) GetAllowedAbilities(playerBin string, currentStep Step) ([]*Ability, error) {
	// 1. Find the current step

	for _, s := range re.game.Steps {
		if s.Bin == currentStep.Bin {
			currentStep = *s
			break
		}
	}
	if currentStep.Bin == "" {
		return nil, fmt.Errorf("invalid step: %s", currentStep.Bin)
	}

	// 2. Find the player's abilities
	var gamer *Gamer
	for _, c := range re.game.Gamers {
		if c.Bin == playerBin {
			gamer = c
			break
		}
	}
	if gamer == nil {
		return nil, fmt.Errorf("invalid player: %s", playerBin)
	}

	// 3. Determine allowed abilities
	var allowedAbilities []*Ability
	for _, ability := range gamer.Abilities {
		// Check if the ability is allowed in the current step
		isAllowedInStep := false
		for _, allowed := range currentStep.Allowed {
			if allowed == ability.Name {
				isAllowedInStep = true
				break
			}
		}

		if isAllowedInStep {
			// Check ability frequency
			switch ability.Frequency {
			case "every":
				allowedAbilities = append(allowedAbilities, ability)
			case "once":
				if ability.TimesUsed == 0 {
					allowedAbilities = append(allowedAbilities, ability)
				}
			case "every_other":
				if re.game.Cycles%2 == 0 {
					allowedAbilities = append(allowedAbilities, ability)
				}
			default:
				return nil, fmt.Errorf("invalid ability frequency: %s", ability.Frequency)
			}
		}
	}

	return allowedAbilities, nil
}

//END RULES ENGINE
