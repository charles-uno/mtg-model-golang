package mtgbf


import (
    "io/ioutil"
    "strconv"
    "strings"
)


func LoadList(name string) (card_seq, error) {
    lines, err := ReadLines("lists/" + name + ".txt")
    if err != nil { return card_seq{}, err }
    card_names := []string{}
    for _, line := range lines {
        n_card := strings.SplitN(line, " ", 2)
        n, err := strconv.Atoi(n_card[0])
        if err != nil { return card_seq{}, err }
        for i := 0; i<n; i++ { card_names = append(card_names, n_card[1]) }
    }
    return Cards(card_names...), nil
}


func ReadLines(filename string) ([]string, error) {
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
