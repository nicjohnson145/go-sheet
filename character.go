package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Character struct {
	Name              string          `yaml:"name"`
	Class             string          `yaml:"class"`
	Race              string          `yaml:"race"`
	Background        string          `yaml:"background"`
	Alignment         string          `yaml:"alignment"`
	PersonalityTraits string          `yaml:"personality-traits"`
	Ideals            string          `yaml:"ideals"`
	Bonds             string          `yaml:"bonds"`
	Flaws             string          `yaml:"flaws"`
	Level             int             `yaml:"level"`
	Attributes        *Attributes      `yaml:"attributes"`
	Proficiency       int             `yaml:"proficiency"`
	Languages         []string        `yaml:"languages"`
	SavingThrows      []string        `yaml:"saving-throws"`
	AllSavingThrows   []string
	Skills            []string        `yaml:"skills"`
	AllSkills         []Skill
	ArmorClass        int             `yaml:"armor-class"`
	Speed             int             `yaml:"speed"`
	HitPoints         *HitPoints       `yaml:"hit-points"`
	HitDice           *HitDice         `yaml:"hit-dice"`
	Weapons           []*Weapon        `yaml:"weapons"`
	Equipment         []*CountableItem `yaml:"equipment"`
	Consumables       []*CountableItem `yaml:"consumables"`
	Features          []*Item          `yaml:"features"`
	Spells            *Spells          `yaml:"spells"`
}

type Attributes struct {
	Strength     int `yaml:"strength"`
	Dexterity    int `yaml:"dexterity"`
	Constitution  int `yaml:"constitution"`
	Intelligence int `yaml:"intelligence"`
	Wisdom       int `yaml:"wisdom"`
	Charisma     int `yaml:"charisma"`
}

type Skill struct {
	Name string
	Mod  string
}

type HitPoints struct {
	Current int `yaml:"current"`
	Max     int `yaml:"max"`
	Temp    int `yaml:"temp"`
}

type HitDice struct {
	Current int    `yaml:"current"`
	Max     int    `yaml:"max"`
	dice    string `yaml:"dice"`
}

type Weapon struct {
	Name       string   `yaml:"name"`
	Damage     *Damage   `yaml:"damage"`
	Properties []string `yaml:"properties"`
	Desc       string   `yaml:"desc"`
	Proficient bool     `yaml:"proficient"`
}

type Damage struct {
	Dice string `yaml:"dice"`
	Type string `yaml:"type"`
}

type Item struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
}

type CountableItem struct {
	Item
	Count int `yaml:"count"`
}

type Spell struct {
	Name          string   `yaml:"name"`
	Range         string   `yaml:"range"`
	Duration      string   `yaml:"duration"`
	Concentration bool     `yaml:"concentration"`
	Components    []string `yaml:"components"`
	Ritual        bool     `yaml:"ritual"`
	CastingTime   string   `yaml:"casting-time"`
	Desc          string   `yaml:"desc"`
}

type Spells struct {
	Cantrips []*Spell `yaml:"cantrips"`
	Level1   []*Spell `yaml:"level-1"`
	Level2   []*Spell `yaml:"level-2"`
	Level3   []*Spell `yaml:"level-3"`
	Level4   []*Spell `yaml:"level-4"`
	Level5   []*Spell `yaml:"level-5"`
	Level6   []*Spell `yaml:"level-6"`
	Level7   []*Spell `yaml:"level-7"`
	Level8   []*Spell `yaml:"level-8"`
	Level9   []*Spell `yaml:"level-9"`
}

func (c *Character) setDefaultData() {
	c.AllSavingThrows = []string{
		"Strength",
		"Dexterity",
		"Constitution",
		"Intelligence",
		"Wisdom",
		"Charisma",
	}

	c.AllSkills = []Skill{
		{Name: "Acrobatics", Mod: "Dexterity"},
		{Name: "Animal Handling", Mod: "Wisdom"},
		{Name: "Arcana", Mod: "Intelligence"},
		{Name: "Athletics", Mod: "Strength"},
		{Name: "Deception", Mod: "Charisma"},
		{Name: "History", Mod: "Intelligence"},
		{Name: "Insight", Mod: "Wisdom"},
		{Name: "Intimidation", Mod: "Charisma"},
		{Name: "Investigation", Mod: "Intelligence"},
		{Name: "Medicine", Mod: "Wisdom"},
		{Name: "Nature", Mod: "Intelligence"},
		{Name: "Perception", Mod: "Wisdom"},
		{Name: "Performance", Mod: "Charisma"},
		{Name: "Persuasion", Mod: "Charisma"},
		{Name: "Religion", Mod: "Intelligence"},
		{Name: "Sleight of Hand", Mod: "Dexterity"},
		{Name: "Stealth", Mod: "Dexterity"},
		{Name: "Survival", Mod: "Wisdom"},
	}
}

func (c *Character) calcMod(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

func (c *Character) modString(mod int) string {
	if mod < 0 {
		return strconv.Itoa(mod)
	} else {
		return fmt.Sprintf("+%v", mod)
	}
}

func (c *Character) modStringForSkill(s Skill) string {
	mod := c.calcMod(c.attrForString(s.Mod))
	if c.isProficientInSkill(s) {
		mod += c.Proficiency
	}

	return c.modString(mod)
}

func (c *Character) modStringForSave(save string) string {
	mod := c.calcMod(c.attrForString(save))
	if c.isProficientInSave(save) {
		mod += c.Proficiency
	}
	return c.modString(mod)
}

func (c *Character) attrForString(s string) int {
	switch s {
	case "Strength":
		return c.Attributes.Strength
	case "Dexterity":
		return c.Attributes.Dexterity
	case "Constitution":
		return c.Attributes.Constitution
	case "Intelligence":
		return c.Attributes.Intelligence
	case "Wisdom":
		return c.Attributes.Wisdom
	case "Charisma":
		return c.Attributes.Charisma
	default:
		panic("wtf did you do")
	}
}

func (c *Character) isProficientInSkill(s Skill) bool {
	return c.inList(s.Name, c.Skills)
}

func (c *Character) isProficientInSave(s string) bool {
	return c.inList(s, c.SavingThrows)
}

func (c *Character) inList(item string, list []string) bool {
	term := strings.ToUpper(item)
	for _, i := range list {
		if strings.ToUpper(i) == term {
			return true
		}
	}
	return false
}

func newCharacter(path string) (*Character, error) {
	c := &Character{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return c, err
	}

	c.setDefaultData()

	return c, nil
}

