package mtgbf


import (
    "io/ioutil"
    "math/rand"
    "strconv"
    "strings"
    "time"
)


func load_list(name string) ([]string, error) {
    lines, err := read_lines("lists/" + name + ".txt")
    cards := []string{}
    if err != nil { return cards, err }
    for _, line := range lines {
        n_card := strings.SplitN(line, " ", 2)
        n, err := strconv.Atoi(n_card[0])
        if err != nil { return []string{}, err }
        for i := 0; i<n; i++ { cards = append(cards, n_card[1]) }
    }
    return shuffled(cards), nil
}


func shuffled(arr_old []string) []string {
    // This shouldn't happen very often. Should be fine to re-seed every time
    rand.Seed(time.Now().UTC().UnixNano())
    arr_new := make([]string, len(arr_old))
    for i, j := range rand.Perm(len(arr_old)) {
        arr_new[i] = arr_old[j]
    }
    return arr_new
}


func read_lines(filename string) ([]string, error) {
    // Load a file and return it as a sequence of strings. Skip empty
    // lines and comments.
    lines := []string{}
    text_bytes, err := ioutil.ReadFile(filename)
    if err != nil { return lines, err }
    for _, line := range strings.Split(string(text_bytes), "\n") {
        if len(line) == 0 { continue }
        if line[:1] == "#" { continue }
        lines = append(lines, line)
    }
    return lines, nil
}
