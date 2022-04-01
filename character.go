package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type SkillSaveBonuser interface {
	GetBonuses() []*SkillSaveBonus
}

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
	Name              string                   `yaml:"name"`
	Class             string                   `yaml:"class"`
	Race              string                   `yaml:"race"`
	Background        string                   `yaml:"background"`
	Alignment         string                   `yaml:"alignment"`
	PersonalityTraits string                   `yaml:"personality-traits"`
	Ideals            string                   `yaml:"ideals"`
	Bonds             string                   `yaml:"bonds"`
	Flaws             string                   `yaml:"flaws"`
	Level             int                      `yaml:"level"`
	Attributes        *Attributes              `yaml:"attributes"`
	Proficiency       int                      `yaml:"proficiency"`
	Proficiencies     Proficiencies            `yaml:"proficiencies"`
	Expertise         Expertise                `yaml:"expertise"`
	Languages         []string                 `yaml:"languages"`
	SavingThrows      []string                 `yaml:"saving-throws"`
	Skills            []string                 `yaml:"skills"`
	ArmorClass        int                      `yaml:"armor-class"`
	Speed             int                      `yaml:"speed"`
	HitPoints         *HitPoints               `yaml:"hit-points"`
	HitDice           *HitDice                 `yaml:"hit-dice"`
	Weapons           []*Weapon                `yaml:"weapons"`
	Equipment         []*CountableItem         `yaml:"equipment"`
	Consumables       []*CountableItem         `yaml:"consumables"`
	Currency          *Currency                `yaml:"currency"`
	Features          []*Item                  `yaml:"features"`
	SpellReset        string                   `yaml:"spell-reset"`
	Spells            map[string]*SpellSection `yaml:"spells"`
	Loot              []*CountableItem         `yaml:"loot"`
	Resources         []*Resource              `yaml:"resources"`
	AllSkills         []Skill                  `yaml:"-"`
	AllSavingThrows   []string                 `yaml:"-"`
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

func (h *HitPoints) SetVal(new int) {
	h.Current = new
}

func (h *HitPoints) CurrentVal() int {
	return h.Current
}

type HitDice struct {
	Current int    `yaml:"current"`
	Max     int    `yaml:"max"`
	Dice    string `yaml:"dice"`
}

func (h *HitDice) SetVal(new int) {
	h.Current = new
}

func (h *HitDice) CurrentVal() int {
	return h.Current
}

type Resource struct {
	Name     string `yaml:"name"`
	Current  int    `yaml:"current"`
	Max      int    `yaml:"max"`
	ResetsOn string `yaml:"resets-on"`
}

func (r *Resource) SetVal(new int) {
	r.Current = new
}

func (r *Resource) CurrentVal() int {
	return r.Current
}

type Weapon struct {
	Name            string            `yaml:"name"`
	Attribute       string            `yaml:"attribute"`
	Damage          *Damage           `yaml:"damage"`
	Properties      []string          `yaml:"properties"`
	Desc            string            `yaml:"desc"`
	Proficient      bool              `yaml:"proficient"`
	Range           string            `yaml:"range"`
	AdditionalToHit int               `yaml:"additional-to-hit,omitempty"`
	Bonuses         []*SkillSaveBonus `yaml:"bonuses,omitempty"`
}

func (w Weapon) GetRange() string {
	if w.Range == "" {
		return "5 ft"
	}
	return w.Range
}

func (w Weapon) GetBonuses() []*SkillSaveBonus {
	return w.Bonuses
}

type Damage struct {
	Dice             string `yaml:"dice"`
	Type             string `yaml:"type"`
	AdditionalDamage int    `yaml:"additional-damage,omitempty"`
}

type Item struct {
	Name    string            `yaml:"name"`
	Desc    string            `yaml:"desc"`
	Bonuses []*SkillSaveBonus `yaml:"bonuses,omitempty"`
}

func (w Item) GetBonuses() []*SkillSaveBonus {
	return w.Bonuses
}

type SkillSaveBonus struct {
	Skill     string `yaml:"skill,omitempty"`
	Save      string `yaml:"save,omitempty"`
	Advantage bool   `yaml:"advantage,omitempty"`
	Bonus     int    `yaml:"bonus,omitempty"`
}

type CountableItem struct {
	Name    string            `yaml:"name"`
	Desc    string            `yaml:"desc"`
	Count   int               `yaml:"count"`
	Bonuses []*SkillSaveBonus `yaml:"bonuses,omitempty"`
}

func (w CountableItem) GetBonuses() []*SkillSaveBonus {
	return w.Bonuses
}

type Currency struct {
	Copper   int `yaml:"copper"`
	Silver   int `yaml:"silver"`
	Electrum int `yaml:"electrum"`
	Gold     int `yaml:"gold"`
	Platinum int `yaml:"platinum"`
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

type SpellSection struct {
	Slots    int      `yaml:"slots"`
	MaxSlots int      `yaml:"max-slots"`
	Spells   []*Spell `yaml:"spells"`
}

func (s *SpellSection) SetVal(new int) {
	s.Slots = new
}

func (s *SpellSection) CurrentVal() int {
	return s.Slots
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


func (c *Character) weaponModString(mod int) string {
	s := c.modString(mod, 0)
	return s[0:len(s) - 4]
}

func (c *Character) modString(mod int, advantage int) string {
	var val string
	if mod < 0 {
		val = strconv.Itoa(mod)
	} else {
		val = fmt.Sprintf("+%v", mod)
	}
	ad := "    "
	if advantage > 0 {
		ad = " (A)"
	} else if advantage < 0 {
		ad = " (D)"
	}
	
	return val + ad
}

func (c *Character) modStringForSkill(s Skill) string {
	mod := c.calcMod(c.attrForString(s.Mod))
	if c.isProficientInSkill(s) {
		mod += c.Proficiency
		if c.isExpertInSkill(s) {
			mod += c.Proficiency
		}
	}

	advantage, bonus := c.advantageAndBonusFromOutside(s.Name)
	mod += bonus

	return c.modString(mod, advantage)
}

func (c *Character) modStringForSave(save string) string {
	mod := c.calcMod(c.attrForString(save))
	if c.isProficientInSave(save) {
		mod += c.Proficiency
	}
	advantage, bonus := c.advantageAndBonusFromOutside(save)
	mod += bonus

	return c.modString(mod, advantage)
}

func (c *Character) modStringForWeapon(w Weapon) string {
	mod := c.calcMod(c.attrForString(w.Attribute)) + w.AdditionalToHit
	if w.Proficient {
		mod += c.Proficiency
	}
	return c.weaponModString(mod)
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

func (c *Character) advantageAndBonusFromOutside(skillOrSave string) (int, int) {
	// This series of checks sort of relies on the fact that more than 2 things don't interact with
	// the same skill. For example, if 2 things gave advantage, and 1 thing gave disadvantage, then
	// this method would incorrectly report advantage, even though RAW that's still a neutral roll.
	// The odds of that are low so *shrug* fuck it.
	advantage := 0
	bonus := 0

	list := []SkillSaveBonuser{}
	for _, f := range c.Features {
		list = append(list, f)
	}
	newAd, newBonus := advantageAndBonusLoop(list, skillOrSave)
	advantage += newAd
	bonus += newBonus

	list = []SkillSaveBonuser{}
	for _, f := range c.Features {
		list = append(list, f)
	}
	newAd, newBonus = advantageAndBonusLoop(list, skillOrSave)
	advantage += newAd
	bonus += newBonus

	list = []SkillSaveBonuser{}
	for _, f := range c.Features {
		list = append(list, f)
	}
	newAd, newBonus = advantageAndBonusLoop(list, skillOrSave)
	advantage += newAd
	bonus += newBonus

	// Normalize multiple sources of advantage/disadvantage to just one
	if advantage > 0 {
		advantage = 1
	} else if advantage < 0 {
		advantage = -1
	}
	
	return advantage, bonus
}

func advantageAndBonusLoop(list []SkillSaveBonuser, skillOrSave string) (int, int) {
	advantage := 0
	bonus := 0
	for _, item := range list {
		bonuses := item.GetBonuses()
		if len(bonuses) == 0 {
			continue
		}

		for _, b := range bonuses {
			if b.Save != skillOrSave && b.Skill != skillOrSave {
				continue
			}
			if b.Advantage {
				advantage += 1
			}
			if b.Bonus != 0 {
				bonus += b.Bonus
			}
		}
	}
	return advantage, bonus
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

	if path == "" {
		return nil
	}

	return ioutil.WriteFile(path, data, 0664)
}
