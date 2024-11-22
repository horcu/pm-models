package types

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/horcu/pm-models/enums"
	"time"
)
import s "github.com/horcu/pm-store"

type Gamer struct {
	Bin         string              `json:"bin"`
	GameId      string              `json:"game_id"`
	CharacterId string              `json:"character_id"`
	Name        string              `json:"name"`
	Fate        *Fate               `json:"fate,omitempty"`
	ImageUrl    string              `json:"image_url"`
	IsAlive     bool                `json:"is_alive"`
	SeenSteps   []string            `json:"seen_steps"`  //list of step bins
	VotedSteps  []string            `json:"voted_steps"` // map of cycles and list of step bins
	Abilities   map[string]*Ability `json:"abilities,omitempty"`
	Metrics     *Metrics            `json:"metrics,omitempty"`
	State       enums.State         `json:"state,omitempty"`
}

func (g *Gamer) AutoVote(game *Game, store *s.Store, abilities []*Ability) {

	// choose a random target from the list of gamers
	target := g.getRandomTarget(game)
	ability := g.getRandomAbility(abilities)
	vote := &Vote{
		Bin:       uuid.New().String(),
		Source:    g.Bin,
		Target:    target.Bin,
		GameBin:   game.Bin,
		StepBin:   game.CurrentStep,
		Ability:   ability.Name,
		TimeStamp: time.Now().String(),
		VotedBy:   g.Bin,
	}

	store.Vote(vote)

}

func (g *Gamer) getRandomTarget(game *Game) *Gamer {
	//  choose a truly random gamer from the list of game.Gamers
	var gamers []*Gamer
	for _, gamer := range game.Gamers {
		if gamer.Bin != g.Bin && gamer.IsAlive {
			gamers = append(gamers, gamer)
		}
	}
	var n int
	if len(gamers) == 0 {
		return g
	}

	err := binary.Read(rand.Reader, binary.LittleEndian, &n)
	if err != nil {
		return nil
	} // Read random bytes into n

	index := n % len(gamers)
	if index < 0 {
		index = -index
	}
	return gamers[index]

}

func (g *Gamer) getRandomAbility(abilities []*Ability) *Ability {
	var usableAbilities []*Ability
	for _, ability := range abilities {
		usableAbilities = append(usableAbilities, ability)
	}
	if len(usableAbilities) == 0 {
		return nil
	}
	var n int
	err := binary.Read(rand.Reader, binary.LittleEndian, &n)
	if err != nil {
		return nil
	}
	index := n % len(usableAbilities)
	if index < 0 {
		index = -index
	}
	return usableAbilities[index]

}
