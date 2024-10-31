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
	UserName     string                 `json:"user_name"`
	Bin          string                 `json:"bin"`
	Photo        string                 `json:"photo"`
	Status       string                 `json:"status"` //inAGame, afk, available
	Privacy      string                 `json:"privacy"`
	Invitations  map[string]*Invitation `json:"invitations,omitempty"`
	GroupIds     []string               `json:"group_ids,omitempty"`
	GameIds      []string               `json:"game_ids,omitempty"`
	Achievements *Achievements          `json:"achievements,omitempt"`
}

type Group struct {
	Bin       string    `json:"bin"`
	Creator   *Player   `json:"creator"`
	Members   []*Player `json:"members,omitempty"`
	GroupName string    `json:"group_name"`
	Capacity  int       `json:"capacity"`
	Status    string    `json:"status"`
}

type Game struct {
	Bin               string                `json:"bin"`
	IsDaytime         bool                  `json:"is_daytime,omitempty"`
	ExplanationSeen   bool                  `json:"explanation_seen,omitempty"`
	FirstDayCompleted bool                  `json:"first_day_completed,omitempty"`
	Level             string                `json:"level"`
	CurrentStep       string                `json:"current_step"`
	GroupId           string                `json:"group_id"`
	Status            string                `json:"status"`
	ServerName        string                `json:"server_name,omitempty"`
	ServerAddress     string                `json:"server_ip,omitempty"`
	ServerPort        int                   `json:"server_port,omitempty"`
	StartTime         string                `json:"start_time,omitempty"`
	EndTime           string                `json:"end_time,omitempty"`
	Creator           *Player               `json:"creator,omitempty"`
	Cycles            int                   `json:"cycles,omitempty"`
	Invited           []string              `json:"invited,omitempty"`
	AllVotesSubmitted bool                  `json:"all_votes_submitted,omitempty"`
	Steps             map[string]*Step      `json:"steps,omitempty"`
	Gamers            map[string]*Gamer     `json:"gamers,omitempty"`
	Characters        map[string]*Character `json:"characters,omitempty"`
}

type Step struct {
	Bin               string                `json:"bin"`
	NextStep          string                `json:"next_step"`
	StepType          string                `json:"step_type"`
	Duration          string                `json:"duration"`
	Command           string                `json:"command"`
	ShowTimer         bool                  `json:"show_timer"`
	CharPhotoUrl      string                `json:"char_photo_url"`
	Characters        map[string]*Character `json:"characters"`
	StepIndex         int                   `json:"step_index"`
	StartTime         string                `json:"start_time"`
	RequiresVote      bool                  `json:"requires_vote"`
	VoteType          string                `json:"vote_type"`
	VillainVoteCount  int                   `json:"villain_vote_count"`
	InnocentVoteCount int                   `json:"innocent_vote_count"`
	EndTime           string                `json:"end_time"`
	Result            map[int]*Result       `json:"result"` // map of game cycles and results
	Allowed           []string              `json:"allowed"`
	SubSteps          map[string]*Step      `json:"sub_steps,omitempty"`
}

type Result struct {
	Bin         string                  `json:"bin"`
	Index       int                     `json:"index"`
	EndTime     string                  `json:"end_time"`
	StepHistory map[string]*StepHistory `json:"decisions,omitempty"` //map of playerId and their stepHistory
	TimeStamp   string                  `json:"timestamp"`
}

type Gamer struct {
	Bin         string     `json:"bin"`
	GameId      string     `json:"game_id"`
	CharacterId string     `json:"character_id"`
	Name        string     `json:"name"`
	ImageUrl    string     `json:"image_url"`
	IsAlive     bool       `json:"is_alive"`
	Abilities   []*Ability `json:"abilities"`
	Metrics     *Metrics   `json:"metrics"`
}

type StepHistory struct {
	Bin            string      `json:"bin"`
	Stamp          string      `json:"stamp"` //millis timestamp same as key for step history entry
	StepBin        string      `json:"step_bin"`
	GameBin        string      `json:"game_bin"`
	VotedAgainstBy []string    `json:"voted_against_by"`
	VoteAction     *VoteAction `json:"actions,omitempty"`
}

type VoteAction struct {
	Bin     string            `json:"bin"`
	StepBin string            `json:"step_bin"`
	GameBin string            `json:"game_bin"`
	Vote    Vote              `json:"vote,omitempty"`
	Media   map[string]*Media `json:"media,omitempty"` // map of string (timestamp millis) and message
}

type Vote struct {
	Bin       string `json:"bin"`
	StepBin   string `json:"step_bin"`
	Target    string `json:"target"`
	Source    string `json:"source"`
	TimeStamp string `json:"time_stamp"`
	VotedBy   string `json:"voted_by"`
	Ability   string `json:"ability"` // what ability to use on the target
}

type Achievements struct {
	Levels     []*Level     `json:"levels"`
	Characters []*Character `json:"characters"`
	Stats      []*GameStats `json:"stats"`
}

type GameStats struct {
	GameIds    []string `json:"game_ids"`
	Roles      []string `json:"roles"`
	Wins       int      `json:"wins"`
	Losses     int      `json:"losses"`
	TotalGames int      `json:"total_games"`
}

type Leaderboard struct {
	Players      []*Player          `json:"players"`
	BestVillain  map[string]*Player `json:"best_villain"`
	BestInnocent map[string]*Player `json:"best_innocent"`
	MostWins     *Player            `json:"most_wins"`
}

type Level struct {
	Bin           string `json:"bin"`
	Value         string `json:"value"`
	AwardedPoints int    `json:"awarded_points"`
	Timestamp     string `json:"timestamp"`
}

type Media struct {
	Bin       string `json:"bin"`
	Text      string `json:"text"`
	Icon      string `json:"icon"`
	HasImage  bool   `json:"has_image"`
	ImagePath string `json:"image_path,omitempty"`
	AudioPath string `json:"audio_path,omitempty"`
	VideoPath string `json:"video_path,omitempty"`
	TimeStamp string `json:"time_stamp,omitempty"`
}

type CharacterPack struct {
	Name       string       `json:"name"`
	OriginDate string       `json:"origin_date"`
	Creator    string       `json:"creator"`
	Characters []*Character `json:"characters,omitempty"`
}

type Character struct {
	Bin         string   `json:"bin"`
	TypeId      int      `json:"type_id"`
	Name        string   `json:"name"`
	ImageUrl    string   `json:"image_url"`
	Description string   `json:"description"`
	Abilities   []string `json:"abilities,omitempty"`
	Role        string   `json:"role"`
}

type Metrics struct {
	LastHeal           string   `json:"last_heal,omitempty"`
	LastVote           string   `json:"last_vote,omitempty"`
	LastKill           string   `json:"last_kill,omitempty"`
	LastPoisoning      string   `json:"last_poisoning,omitempty"`
	LastMimic          string   `json:"last_mimic,omitempty"`
	LastTrick          string   `json:"last_trick,omitempty"`
	Heals              []string `json:"heals,omitempty"`
	TimesHealed        int      `json:"times_healed,omitempty"`
	Kills              []string `json:"kills,omitempty"`
	Votes              []string `json:"votes,omitempty"`
	Mimics             []string `json:"mimics,omitempty"`
	TimesMimicked      int      `json:"times_mimicked,omitempty"`
	Tricks             []string `json:"tricks,omitempty"`
	TimesTricked       int      `json:"times_tricked,omitempty"`
	Poisonings         []string `json:"poisonings,omitempty"`
	TimesPoisoned      int      `json:"times_poisoned,omitempty"`
	VotedAgainstBy     []string `json:"voted_against_by,omitempty"`
	VotedAgainstMostBy string   `json:"voted_against_most_by,omitempty"`
	VoteSuccessRate    float64  `json:"vote_success_rate,omitempty"`
}

type Ability struct {
	Bin            string `json:"bin"`
	Name           string `json:"name"`
	Character      string `json:"character"`
	Description    string `json:"description"`
	CycleUsedIndex int    `json:"cycle_used_index"`
	Frequency      string `json:"frequency"`
	TimesUsed      int    `json:"times_used,omitempty"`
	Instructions   string `json:"instructions,omitempty"`
}

type StepSequence struct {
	SDuration string `json:"s_duration"`
	CStep     string `json:"c_step"`
}
