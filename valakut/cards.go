
package valakut

// ---------------------------------------------------------------------

import (
    "sort"
    "strings"
    "strconv"
)

// ---------------------------------------------------------------------

func slug(text string) string {
    switch text {
        case "Oath of Nissa":
            return "Oath"
        case "Primeval Titan":
            return "Titan"
        case "Sakura-Tribe Elder":
            return "STE"
        case "Scapeshift":
            return "Shift"
        case "Simian Spirit Guide":
            return "SSG"
        case "Stomping Ground":
            return "Shock"
        case "Summoner's Pact":
            return "Pact"
        case "Through the Breach":
            return "Breach"
        case "Valakut, the Molten Pinnacle":
            return "Valakut"
        case "Wooded Foothills":
            return "Fetch"
    }
    // By default, pull out spaces and punctuation.
    ret := text
    for _, c := range []string{" ", "-", "'", ","} {
        ret = strings.ReplaceAll(ret, c, "")
    }
    return ret
}

// ---------------------------------------------------------------------

func tally(arr []string) string {
    counts := make(map[string]int)
    for _, card := range arr { counts[slug(card)] += 1 }
    name_count := []string{}
    for name, count := range counts {
        nc := name
        if count > 1 { nc += "*" + strconv.Itoa(count) }
        name_count = append(name_count, nc)
    }
    sort.Strings(name_count)
    return strings.Join(name_count, " ")
}

// ---------------------------------------------------------------------

func get_cost(card string) string {
    switch card {
        case "Arboreal Grazer":
            return "G"
        case "Deadshot Minotaur":
            return "R"
        case "Desperate Ritual":
            return "1R"
        case "Explore":
            return "1G"
        case "Farseek":
            return "1G"
        case "Oath of Nissa":
            return "G"
        case "Mwonvuli Acid-Moss":
            return "2GG"
        case "Primeval Titan":
            return "4GG"
        case "Prismatic Omen":
            return "1G"
        case "Sakura-Tribe Elder":
            return "1G"
        case "Scapeshift":
            return "2GG"
        // Treat SFT as a 1-cost spell with kicker 2.
        case "Search for Tomorrow":
            return "G"
        case "Shefet Monitor":
            return "3G"
        case "Sleight of Hand":
            return "G"
        case "Summoner's Pact":
            return "4GG"
        case "Through the Breach":
            return "4R"
        case "Wood Elves":
            return "2G"
        case "Woodfall Primus":
            return "5GGG"
        default:
            return ""
    }
}

// ---------------------------------------------------------------------

func nlands(cards []string) int {
    n := 0
    for _, card := range cards {
        if is_land(card) { n += 1 }
    }
    return n
}

// ---------------------------------------------------------------------

func is_land(card string) bool {
    lands := []string{
        "Forest",
        "Mountain",
        "Cinder Glade",
        "Sheltered Thicket",
        "Shivan Oasis",
        "Stomping Ground",
        "Taiga",
        "Tapped Taiga",
        "Valakut, the Molten Pinnacle",
        "Wooded Foothills",
    }
    for _, c := range lands {
        if c == card { return true }
    }
    return false
}

func is_creature(card string) bool {
    lands := []string{
        "Arboreal Grazer",
        "Deadshot Minotaur",
        "Primeval Titan",
        "Simian Spirit Guide",
        "Sakura-Tribe Elder",
        "Shefet Monitor",
        "Wood Elves",
        "Woodfall Primus",
    }
    for _, c := range lands {
        if c == card { return true }
    }
    return false
}

func land_output(card string) string {
    switch card {
        case "Cinder Glade":
            return "G"
        case "Forest":
            return "G"
        case "Mountain":
            return "R"
        case "Sheltered Thicket":
            return "G"
        case "Shivan Oasis":
            return "G"
        case "Stomping Ground":
            return "G"
        case "Taiga":
            return "G"
        case "Valakut, the Molten Pinnacle":
            return "R"
        default:
            return ""
    }
}

// ---------------------------------------------------------------------

func is_spell(card string) bool {
    return card == "Sheltered Thicket" || !is_land(card)
}
