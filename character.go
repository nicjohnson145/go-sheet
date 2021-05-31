package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	Cantrips = "cantrips"
	Level1   = "level-1"
	Level2   = "level-2"
	Level3   = "level-3"
	Level4   = "level-4"
	Level5   = "level-5"
	Level6   = "level-6"
	Level7   = "level-7"
	Level8   = "level-8"
	Level9   = "level-9"
)

type Character struct {
	Name              string                 `yaml:"name"`
	Class             string                 `yaml:"class"`
	Race              string                 `yaml:"race"`
	Background        string                 `yaml:"background"`
	Alignment         string                 `yaml:"alignment"`
	PersonalityTraits string                 `yaml:"personality-traits"`
	Ideals            string                 `yaml:"ideals"`
	Bonds             string                 `yaml:"bonds"`
	Flaws             string                 `yaml:"flaws"`
	Level             int                    `yaml:"level"`
	Attributes        *Attributes            `yaml:"attributes"`
	Proficiency       int                    `yaml:"proficiency"`
	Proficiencies     Proficiencies          `yaml:"proficiencies"`
	Expertise         Expertise              `yaml:"expertise"`
	Languages         []string               `yaml:"languages"`
	SavingThrows      []string               `yaml:"saving-throws"`
	Skills            []string               `yaml:"skills"`
	ArmorClass        int                    `yaml:"armor-class"`
	Speed             int                    `yaml:"speed"`
	HitPoints         *HitPoints             `yaml:"hit-points"`
	HitDice           *HitDice               `yaml:"hit-dice"`
	Weapons           []*Weapon              `yaml:"weapons"`
	Equipment         []*CountableItem       `yaml:"equipment"`
	Consumables       []*CountableItem       `yaml:"consumables"`
	Features          []*Item                `yaml:"features"`
	Spells            map[string]SpellSecion `yaml:"spells"`
	Loot              []*CountableItem       `yaml:"loot"`
	Resources         []*Resource            `yaml:"resources"`
	AllSkills         []Skill
	AllSavingThrows   []string
}

type Attributes struct {
	Strength     int `yaml:"strength"`
	Dexterity    int `yaml:"dexterity"`
	Constitution int `yaml:"constitution"`
	Intelligence int `yaml:"intelligence"`
	Wisdom       int `yaml:"wisdom"`
	Charisma     int `yaml:"charisma"`
}

type Proficiencies struct {
	Armor   []string `yaml:"armor"`
	Tools   []string `yaml:"tools"`
	Weapons []string `yaml:"weapons"`
}

type Expertise struct {
	Tools  []string `yaml:"tools"`
	Skills []string `yaml:"skills"`
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
	Dice    string `yaml:"dice"`
}

type Resource struct {
	Name    string `yaml:"name"`
	Current int    `yaml:"current"`
	Max     int    `yaml:"max"`
}

type Weapon struct {
	Name       string   `yaml:"name"`
	Attribute  string   `yaml:"attribute"`
	Damage     *Damage  `yaml:"damage"`
	Properties []string `yaml:"properties"`
	Desc       string   `yaml:"desc"`
	Proficient bool     `yaml:"proficient"`
	Range      string   `yaml:"range"`
}

func (w Weapon) GetRange() string {
	if w.Range == "" {
		return "5 ft"
	}
	return w.Range
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
	Name  string `yaml:"name"`
	Desc  string `yaml:"desc"`
	Count int    `yaml:"count"`
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

type SpellSecion struct {
	Slots    int      `yaml:"slots"`
	MaxSlots int      `yaml:"max-slots"`
	Spells   []*Spell `yaml:"spells"`
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
		if c.isExpertInSkill(s) {
			mod += c.Proficiency
		}
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

func (c *Character) modStringForWeapon(w Weapon) string {
	mod := c.calcMod(c.attrForString(w.Attribute))
	if w.Proficient {
		mod += c.Proficiency
	}
	return c.modString(mod)
}

func (c *Character) attrForString(s string) int {
	switch strings.ToUpper(s) {
	case "STRENGTH":
		return c.Attributes.Strength
	case "DEXTERITY":
		return c.Attributes.Dexterity
	case "CONSTITUTION":
		return c.Attributes.Constitution
	case "INTELLIGENCE":
		return c.Attributes.Intelligence
	case "WISDOM":
		return c.Attributes.Wisdom
	case "CHARISMA":
		return c.Attributes.Charisma
	default:
		panic("wtf did you do")
	}
}

func (c *Character) isProficientInSkill(s Skill) bool {
	return c.inList(s.Name, c.Skills)
}

func (c *Character) isExpertInSkill(s Skill) bool {
	return c.inList(s.Name, c.Expertise.Skills)
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

	var data []byte
	if path != "" {
		d, err := ioutil.ReadFile(path)
		data = d
		if err != nil {
			return c, err
		}
	} else {
		data = resourceExampleYml.StaticContent
	}

	err := yaml.Unmarshal(data, c)
	if err != nil {
		return c, err
	}

	c.setDefaultData()

	return c, nil
}

func persistCharacter(c Character, path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0664)
}
