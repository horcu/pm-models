package v1

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
	Bin               string            `json:"bin"`
	Steps             map[string]*Step  `json:"steps"omitempty`
	IsDaytime         bool              `json:"is_daytime", omitempty`
	FirstDayCompleted bool              `json:"first_day_completed", omitempty`
	CurrentStep       string            `json:"current_step"`
	GroupId           string            `json:"group_id"`
	Status            string            `json:"status"`
	ServerName        string            `json:"server_name", omitempty`
	ServerIP          string            `json:"server_ip", omitempty`
	ServerPort        int               `json:"server_port", omitempty`
	StartTime         string            `json:"start_time", omitempty`
	EndTime           string            `json:"end_time", omitempty`
	Creator           *Player           `json:"creator",  omitempty`
	Invited           []string          `json:"invited", omitempty`
	Gamers            map[string]*Gamer `json:"gamers", omitempty`
}

type Step struct {
	Bin          string  `json:"bin"`
	NextStep     string  `json:"next_step"`
	Type         string  `json:"type"`
	StepType     string  `json:"step_type"`
	Duration     string  `json:"duration"`
	Command      string  `json:"command"`
	NcCommand    string  `json:"nc_command"`
	CharPhotoUrl string  `json:"char_photo_url"`
	Text         string  `json:"text"`
	NcText       string  `json:"nc_text"`
	Character    string  `json:"character"`
	StepIndex    int     `json:"step_index"`
	StartTime    string  `json:"start_time"`
	EndTime      string  `json:"end_time"`
	Result       *Result `json:"result"`
	Action       string  `json:"action"`
}

type Result struct {
	Bin     string `json:"bin"`
	Index   int    `json:"index"`    // order
	EndTime string `json:"end_time"` // milliseconds
	// poisoned: [player1, player2] // healed : [player1]
	ResultMap map[string]*Gamer `json:"players_affected"omitempty`
	TimeStamp string            `json:"timestamp"`
}

type Gamer struct {
	Bin          string             `json:"bin"`
	GameId       string             `json:"game_id"`
	CharacterId  string             `json:"character_id"`
	Name         string             `json:"name"`
	ImageUrl     string             `json:"image_url"`
	CharImageUrl string             `json:"char_image_url"`
	IsAlive      bool               `json:"is_alive"`
	Actions      map[string]*Action `json:"actions"` // map of string (stepId) and Action
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
	Name       string      `json:"name"`
	OriginDate string      `json:"origin_date"`
	Creator    string      `json:"creator"`
	Characters []Character `json:"characters"`
}

type Character struct {
	Bin         string    `json:"bin"`
	TypeId      int       `json:"type_id"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Abilities   []Ability `json:"abilities"`
	Role        string    `json:"role"`
}

type Ability struct {
	Bin       string `json:"bin"`
	Name      string `json:"name"`
	Character string `json:"character"`
	Frequency string `json:"frequency"`
}
