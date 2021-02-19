package main

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
	Attributes        Attributes      `yaml:"attributes"`
	Proficiency       int             `yaml:"proficiency"`
	Languages         []string        `yaml:"languages"`
	SavingThrows      []string        `yaml:"saving-throws"`
	AllSavingThrows   []string
	Skills            []string        `yaml:"skills"`
	AllSkills         []Skill
	ArmorClass        int             `yaml:"armor-class"`
	Speed             int             `yaml:"speed"`
	HitPoints         HitPoints       `yaml:"hit-points"`
	HitDice           HitDice         `yaml:"hit-dice"`
	Weapons           []Weapon        `yaml:"weapons"`
	Equipment         []CountableItem `yaml:"equipment"`
	Consumables       []CountableItem `yaml:"consumables"`
	Features          []Item          `yaml:"features"`
	Spells            Spells          `yaml:"spells"`
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
	Damage     Damage   `yaml:"damage"`
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
	Cantrips []Spell `yaml:"cantrips"`
	Level1   []Spell `yaml:"level-1"`
	Level2   []Spell `yaml:"level-2"`
	Level3   []Spell `yaml:"level-3"`
	Level4   []Spell `yaml:"level-4"`
	Level5   []Spell `yaml:"level-5"`
	Level6   []Spell `yaml:"level-6"`
	Level7   []Spell `yaml:"level-7"`
	Level8   []Spell `yaml:"level-8"`
	Level9   []Spell `yaml:"level-9"`
}
