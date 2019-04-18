
package valakut

// ---------------------------------------------------------------------

import (
    "fmt"
    "math/rand"
    "sort"
    "strings"
    "strconv"
    "time"
)

// ---------------------------------------------------------------------

func InitRandom() {
    rand.Seed(time.Now().UTC().UnixNano())
}

// ---------------------------------------------------------------------

func contains(arr []string, elt string) bool {
    for _, a := range arr {
        if a == elt { return true }
    }
    return false
}



func b2i2a(b bool) string {
    if b { return "1" } else { return "0" }
}



func count(arr []string, r string) (n int) {
    for _, a := range arr {
        if a == r { n += 1 }
    }
    return n
}

// ---------------------------------------------------------------------

func slug(text string) string {
    switch text {
        case "Primeval Titan":
            return "Titan"
        case "Sakura-Tribe Elder":
            return "STE"
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

func remove(arr []string, elt string) []string {
    for i, a := range arr {
        if a == elt {
            arr[i] = arr[len(arr)-1]
            new_arr := arr[:len(arr)-1]
            return new_arr
        }
    }
    fmt.Println("WARNING: Failed to remove", elt, "from", arr)
    return []string{}
}

// ---------------------------------------------------------------------

func shuffled(deck []string) []string {
    ret := make([]string, len(deck))
    for i, j := range rand.Perm(len(deck)) { ret[i] = deck[j] }
    return ret
}

// ---------------------------------------------------------------------

func flip() bool {
    return rand.Intn(2) == 0
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

func unique_strings(arr []string) []string {
    counts := make(map[string]int)
    for _, card := range arr { counts[card] += 1 }
    uarr := []string{}
    for card, _ := range counts { uarr = append(uarr, card) }
    return uarr
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
        case "Primeval Titan":
            return "4GG"
        case "Sakura-Tribe Elder":
            return "1G"
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
        default:
            return ""
    }
}

// ---------------------------------------------------------------------

func is_land(card string) bool {
    lands := []string{
        "Forest",
        "Mountain",
        "Cinder Glade",
        "Sheltered Thicket",
        "Stomping Ground",
        "Taiga",
        "Valakut, the Molten Pinnacle",
        "Wooded Foothills",
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
