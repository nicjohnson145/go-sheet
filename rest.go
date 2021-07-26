package main

import (
	"math"
)

const (
	ShortRest = "short-rest"
	LongRest  = "long-rest"
)

func (s *sheet) longRest() {
	// Reset Health
	s.character.HitPoints.Current = s.character.HitPoints.Max
	// Reset up to half hit dice
	halfHitDice := int(math.Floor(float64(s.character.HitDice.Max) / 2.0))
	s.character.HitDice.Current += halfHitDice
	if s.character.HitDice.Current > s.character.HitDice.Max {
		s.character.HitDice.Current = s.character.HitDice.Max
	}
	// Loop through resources and reset any long rest ones
	s.resetResources(LongRest)
	// Assume resources not tagged are long rest resources
	s.resetResources("")
	// Reset any spell slots
	s.resetSpells()
	// Also reset any short rest abilities
	s.shortRest()
}

func (s *sheet) shortRest() {
	// Loop through resources and reset any short rest ones
	s.resetResources(ShortRest)

	// If your spells reset on short resets, reset them
	if s.character.SpellReset == ShortRest {
		s.resetSpells()
	}
}

func (s *sheet) resetSpells() {
	for _, section := range s.character.Spells {
		section.Slots = section.MaxSlots
	}
}

func (s *sheet) resetResources(resetsOn string) {
	for _, resource := range s.character.Resources {
		if resource.ResetsOn == resetsOn {
			resource.Current = resource.Max
		}
	}
}
