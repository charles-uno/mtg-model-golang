
package valakut

// ---------------------------------------------------------------------

import (
    "errors"

    "fmt"

    "io/ioutil"
    "math/rand"
    "sort"
    "strings"
    "strconv"
    "time"
)

// ---------------------------------------------------------------------

func load(name string) ([]string, error) {
    // Accept the name of a deck. Return (list, err).
    var list []string
    dat, err := ioutil.ReadFile("lists/" + name + ".txt")
    if err != nil {
        return make([]string, 0), errors.New("List not found")
    }
    for _, line := range strings.Split(string(dat), "\n") {
        if len(line) == 0 { continue }
        if line[:1] == "#" { continue }
        n_card := strings.SplitN(line, " ", 2)
        n, err := strconv.Atoi(n_card[0])
        if err != nil {
            return make([]string, 0), errors.New("Invalid integer")
        }
        for i := 0; i<n; i++ { list = append(list, n_card[1]) }
    }
    return list, nil
}

// ---------------------------------------------------------------------

func count(arr []string, r string) (n int) {
    for _, a := range arr {
        if a == r { n += 1 }
    }
    return n
}

// ---------------------------------------------------------------------

func slug(text string) string {
    ret := text
    for _, c := range []string{" ", "-", "'"} {
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
    r := rand.New(rand.NewSource(time.Now().Unix()))
    ret := make([]string, len(deck))
    for i, j := range r.Perm(len(deck)) { ret[i] = deck[j] }
    return ret
}

// ---------------------------------------------------------------------

func flip() bool {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    return r.Intn(2) == 0
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

func uniques(arr []string) []string {
    counts := make(map[string]int)
    for _, card := range arr { counts[card] += 1 }
    uarr := []string{}
    for card, _ := range counts { uarr = append(uarr, card) }
    return uarr
}






// ---------------------------------------------------------------------

func is_land(card string) bool {
    lands := []string{"Forest", "Mountain", "Taiga", "Wooded Foothills", "Stomping Ground", "Cinder Glade", "Valakut, the Molten Pinnacle", "Sheltered Thicket"}
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
