package main

type dotaMatchHistory struct {
	Result dotaMatchHistoryResult `json:"result"`
}

type dotaMatchHistoryResult struct {
	Status    int                     `json:"status"`
	StatusMsg string                  `json:"statusDetail"`
	Matches   []dotaMatchHistoryMatch `json:"matches"`
}

type dotaMatchHistoryMatch struct {
	ID uint64 `json:"match_id"`
}

type dotaMatchDetails struct {
	Result dotaMatchDetailsResult `json:"result"`
}

type dotaMatchDetailsResult struct {
	GameMode   int                      `json:"game_mode"`
	LobbyType  int                      `json:"lobby_type"`
	RadiantWin bool                     `json:"radiant_win"`
	StartTime  int64                    `json:"start_time"`
	Duration   uint32                   `json:"duration"`
	Players    []dotaMatchDetailsPlayer `json:"players"`
}

type dotaMatchDetailsPlayer struct {
	AccountID uint64 `json:"account_id"`
	HeroID    int    `json:"hero_id"`
	Slot      uint8  `json:"player_slot"`
	Kills     int    `json:"kills"`
	Deaths    int    `json:"deaths"`
	Assists   int    `json:"assists"`
	XPM       int    `json:"xp_per_min"`
	GPM       int    `json:"gold_per_min"`
}

var dotaLobbyTypes = map[int]string{
	-1: "Invalid",
	0:  "Public matchmaking",
	1:  "Practise",
	2:  "Tournament",
	3:  "Tutorial",
	4:  "Co-op with bots",
	5:  "Team match",
	6:  "Solo Queue",
	7:  "Ranked",
}

var dotaModeNames = map[int]string{
	0:  "None",
	1:  "All Pick",
	2:  "Captain's Mode",
	3:  "Random Draft",
	4:  "Single Draft",
	5:  "All Random",
	6:  "Intro",
	7:  "Diretide",
	8:  "Reverse Captain's Mode",
	9:  "The Greeviling",
	10: "Tutorial",
	11: "Mid Only",
	12: "Least Played",
	13: "New Player Pool",
	14: "Compendium Matchmaking",
	16: "Captain's Draft",
}

var dotaHeroNames = map[int]string{
	1:   "Anti-Mage",
	2:   "Axe",
	3:   "Bane",
	4:   "Bloodseeker",
	5:   "Crystal Maiden",
	6:   "Drow Ranger",
	7:   "Earthshaker",
	8:   "Juggernaut",
	9:   "Mirana",
	11:  "Shadow Fiend",
	10:  "Morphling",
	12:  "Phantom Lancer",
	13:  "Puck",
	14:  "Pudge",
	15:  "Razor",
	16:  "Sand King",
	17:  "Storm Spirit",
	18:  "Sven",
	19:  "Tiny",
	20:  "Vengeful Spirit",
	21:  "Windranger",
	22:  "Zeus",
	23:  "Kunkka",
	25:  "Lina",
	31:  "Lich",
	26:  "Lion",
	27:  "Shadow Shaman",
	28:  "Slardar",
	29:  "Tidehunter",
	30:  "Witch Doctor",
	32:  "Riki",
	33:  "Enigma",
	34:  "Tinker",
	35:  "Sniper",
	36:  "Necrophos",
	37:  "Warlock",
	38:  "Beastmaster",
	39:  "Queen of Pain",
	40:  "Venomancer",
	41:  "Faceless Void",
	42:  "Wraith King",
	43:  "Death Prophet",
	44:  "Phantom Assassin",
	45:  "Pugna",
	46:  "Templar Assassin",
	47:  "Viper",
	48:  "Luna",
	49:  "Dragon Knight",
	50:  "Dazzle",
	51:  "Clockwerk",
	52:  "Leshrac",
	53:  "Nature's Prophet",
	54:  "Lifestealer",
	55:  "Dark Seer",
	56:  "Clinkz",
	57:  "Omniknight",
	58:  "Enchantress",
	59:  "Huskar",
	60:  "Night Stalker",
	61:  "Broodmother",
	62:  "Bounty Hunter",
	63:  "Weaver",
	64:  "Jakiro",
	65:  "Batrider",
	66:  "Chen",
	67:  "Spectre",
	69:  "Doom",
	68:  "Ancient Apparition",
	70:  "Ursa",
	71:  "Spirit Breaker",
	72:  "Gyrocopter",
	73:  "Alchemist",
	74:  "Invoker",
	75:  "Silencer",
	76:  "Outworld Devourer",
	77:  "Lycan",
	78:  "Brewmaster",
	79:  "Shadow Demon",
	80:  "Lone Druid",
	81:  "Chaos Knight",
	82:  "Meepo",
	83:  "Treant Protector",
	84:  "Ogre Magi",
	85:  "Undying",
	86:  "Rubick",
	87:  "Disruptor",
	88:  "Nyx Assassin",
	89:  "Naga Siren",
	90:  "Keeper of the Light",
	91:  "Io",
	92:  "Visage",
	93:  "Slark",
	94:  "Medusa",
	95:  "Troll Warlord",
	96:  "Centaur Warrunner",
	97:  "Magnus",
	98:  "Timbersaw",
	99:  "Bristleback",
	100: "Tusk",
	101: "Skywrath Mage",
	102: "Abaddon",
	103: "Elder Titan",
	104: "Legion Commander",
	106: "Ember Spirit",
	107: "Earth Spirit",
	109: "Terrorblade",
	110: "Phoenix",
}
