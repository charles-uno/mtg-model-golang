
package valakut

// ---------------------------------------------------------------------

import (
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "path"
    "sort"
    "strconv"
    "strings"
)

// ---------------------------------------------------------------------




func SaveResult(name string, line string) error {
    return AppendLine("data/" + name + ".out", line)
}





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


func ReadCSV(filename string) ([][]string, error) {
    arr := [][]string{}
    lines, err := ReadLines(filename)
    if err != nil { return arr, err }
    for _, line := range lines {
        fields := strings.Split(line, ",")
        arr = append(arr, fields)
    }
    return arr, nil
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

// ---------------------------------------------------------------------

func PrintSummary(name string) error {
    arr, err := ReadCSV("data/" + name + ".out")
    if err != nil { return err }

    ntrials := len(arr)

    turn_tally := make(map[float64]int)
    for _, fields := range arr {
        turn, _ := strconv.Atoi(fields[0])
        fast, _ := strconv.Atoi(fields[3])
        adjusted_turn := float64(turn) + 0.5 - 0.5*float64(fast)
        turn_tally[adjusted_turn]++
    }

    var turns []float64
    for t, _ := range turn_tally {
        turns = append(turns, t)
    }
    sort.Float64s(turns)

    header := ""
    line := ""
    n := 0
    for _, t := range turns {
        if t < 1 { continue }
        header += fmt.Sprintf("         T%5.1f", t)
        n += turn_tally[t]
        dn := math.Sqrt(float64(n))

        p := float64(100*n)/float64(ntrials)
        dp := 100.*dn/float64(ntrials)

        line += "     " + fmt.Sprintf("%3.0f", p) + "%" + " ± " + fmt.Sprintf("%2.0f", dp) + "%"
    }
    fmt.Println(header)
    fmt.Println(line)

    return nil

}
