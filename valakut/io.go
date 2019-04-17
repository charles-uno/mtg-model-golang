
package valakut

// ---------------------------------------------------------------------

import (
    "io/ioutil"
    "os"
    "path"
    "strconv"
    "strings"
)

// ---------------------------------------------------------------------

func AppendLine(filename string, line string) error {
    filedir := path.Dir(filename)
    if _, err := os.Stat(filedir); os.IsNotExist(err) {
        os.Mkdir(filedir, 0755)
    }
    handle, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil { return err }
    defer handle.Close()
    _, err = handle.WriteString(line + "\n")
    return err
}

// ---------------------------------------------------------------------

func LoadDeck(name string) ([]string, error) {
    lines, err := ReadLines("lists/" + name + ".txt")
    if err != nil { return lines, err }
    list := []string{}
    for _, line := range lines {
        n_card := strings.SplitN(line, " ", 2)
        n, err := strconv.Atoi(n_card[0])
        if err != nil { return []string{}, err }
        for i := 0; i<n; i++ { list = append(list, n_card[1]) }
    }
    return list, nil
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
