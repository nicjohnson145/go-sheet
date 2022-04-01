// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package main

import "fyne.io/fyne/v2"

var resourceExampleYml = &fyne.StaticResource{
	StaticName: "example.yml",
	StaticContent: []byte(
		"name: Example Character\nclass: Sorcerer\nrace: Dragonborn\nbackground: Urchin\nalignment: Lawful Evil\npersonality-traits: |\n  this is where they go\nideals: |\n  this is where they go\nbonds: |\n  this is where they go\nflaws: |\n  this is where they go\nlevel: 5\nattributes:\n  strength: 8\n  dexterity: 10\n  constitution: 10\n  intelligence: 14\n  wisdom: 14\n  charisma: 18\nproficiency: 2\nproficiencies:\n  armor:\n  - light\n  tools:\n  - herbalism kit\n  weapons:\n  - martial\n  - simple\nexpertise:\n  tools:\n  - herbalism kit\n  skills:\n  - survival\n  - sleight of hand\nlanguages:\n- Draconic\n- Common\nsaving-throws:\n- charisma\n- wisdom\nskills:\n- acrobatics\n- intimidation\n- sleight of hand\n- survival\narmor-class: 13\nspeed: 30\nhit-points:\n  current: 20\n  max: 20\n  temp: 0\nhit-dice:\n  current: 5\n  max: 5\n  dice: d8\nweapons:\n- name: Quarterstaff\n  attribute: strength\n  damage:\n    dice: 1d6 (1d8)\n    type: bludgeoning\n  properties:\n  - versatile\n  desc: \"\"\n  proficient: true\n  range: 5 ft\n- name: Longsword\n  attribute: strength\n  damage:\n    dice: 1d8\n    type: slashing\n  properties:\n  - versatile\n  desc: |\n    A lengthy straight blade designed for slashing foes, the longsword is a highly versatile weapon that may also be wielded with two hands for more punishing strikes.\n  proficient: false\n  range: \"\"\n- name: +1 Bow\n  attribute: dexterity\n  damage:\n    dice: 1d8\n    type: piercing\n    additional-damage: 1\n  additional-to-hit: 1\n  proficient: true\n  range: 60/120 ft\nequipment:\n- name: Hat of Disguise\n  desc: |\n    While wearing this hat, you can use an action to cast the disquise self spell from it at will. The spell ends if the hat is removed\n  count: 0\n- name: Boots of Elvenkind\n  desc: |\n    While you wear these boots, your steps make no sound, regardless of the surface you are moving across. You also have advantage on Dexterity (Stealth) checks that rely on moving silently.\n  count: 0\n  bonuses:\n  - skill: Stealth\n    Advantage: true\nconsumables:\n- name: Potion of Healing\n  desc: |\n    You regain 2d4 + 2 hit points when you drink this potion. The potions red liquid glimmers when agitated\n  count: 3\ncurrency:\n  copper: 0\n  silver: 5\n  electrum: 0\n  gold: 10\n  platinum: 1\nfeatures:\n- name: Fey Ancestry\n  desc: |\n    You have advantage on saving throws against being charmed, and magic can't put you to sleep\n- name: Draconic Resistance\n  desc: |\n    As magic flows through your body, it causes physical Traits of your Dragon ancestors to emerge. At 1st level, your hit point maximum increases by 1 and increases by 1 again whenever you gain a level in this class.  Additionally, parts of your skin are covered by a thin sheen of dragon-like scales. When you aren’t wearing armor, your AC equals 13 + your Dexterity modifier.\n- name: Incisive Sense\n  desc: |\n    You have advantage on Intelligence (Investigation) and Wisdom (Insight) checks\n  bonuses:\n  - skill: Investigation\n    advantage: true\n  - skill: Insight\n    advantage: true\nspells:\n  cantrips:\n    slots: 0\n    max-slots: 0\n    spells:\n    - name: Create Bonfire\n      range: 60 ft\n      duration: Up to 1 minute\n      concentration: true\n      components:\n      - V\n      - S\n      ritual: false\n      casting-time: 1 action\n      desc: |\n        You create a bonfire on ground that you can see within range. Until the spell ends, the magic bonfire fills a 5-foot cube. Any creature in the bonfire’s space when you cast the spell must succeed on a Dexterity saving throw or take 1d8 fire damage. A creature must also make the saving throw when it moves into the bonfire’s space for the first time on a turn or ends its turn there.  The bonfire ignites flammable objects in its area that aren’t being worn or carried.  The spell’s damage increases by 1d8 when you reach 5th level (2d8), 11th level (3d8), and 17th level (4d8).\n  level-1:\n    slots: 3\n    max-slots: 4\n    spells:\n    - name: Disguise Self\n      range: self\n      duration: 1 hour\n      concentration: false\n      components:\n      - V\n      - S\n      ritual: false\n      casting-time: 1 action\n      desc: |\n        You make yourself—including your clothing, armor, weapons, and other belongings on your person—look different until the spell ends or until you use your action to dismiss it. You can seem 1 foot shorter or taller and can appear thin, fat, or in between. You can’t change your body type, so you must adopt a form that has the same basic arrangement of limbs. Otherwise, the extent of the illusion is up to you.  The changes wrought by this spell fail to hold up to physical inspection. For example, if you use this spell to add a hat to your outfit, objects pass through the hat, and anyone who touches it would feel nothing or would feel your head and hair. If you use this spell to appear thinner than you are, the hand of someone who reaches out to touch you would bump into you while it was seemingly still in midair.  To discern that you are disguised, a creature can use its action to inspect your appearance and must succeed on an Intelligence (Investigation) check against your spell save DC.\nloot:\n- name: Grell Poison\n  desc: Posion scavenged from a grell\n  count: 3\n- name: Gold\n  desc: \"\"\n  count: 205\nresources:\n- name: Sorcery Points\n  current: 5\n  max: 5\n- name: Wand of Web Charges\n  current: 3\n  max: 4\n"),
}
