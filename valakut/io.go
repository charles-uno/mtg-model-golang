
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

var turns = [...]float64{3.0, 3.5, 4.0, 4.5, 5.0, 5.5}

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

func PrintSummaries() {
    dirname := "data"
    f, err := os.Open(dirname)
    if err != nil { fmt.Println(err); return }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil { fmt.Println(err); return }

    names := []string{}
    namelen := 0
    for _, file := range files {
        name := strings.Split(file.Name(), ".")[0]
        names = append(names, name)
        if len(name) > namelen { namelen = len(name) }
    }

    sort.Strings(names)

    format := "%-" + strconv.Itoa(namelen+1) + "s"

    fmt.Printf(format, "")
    fmt.Println(get_summary_header())
    for _, name := range names {
        summary, _ := get_summary(name)
        line := get_summary_line(summary)
        fmt.Printf(format, name)
        fmt.Println(line)
    }
}





func get_summary(name string) (map[float64]int, error) {
    summary := make(map[float64]int)
    arr, err := ReadCSV("data/" + name + ".out")
    if err != nil { return summary, err }
    for _, fields := range arr {
        turn, _ := strconv.Atoi(fields[0])
        fast, _ := strconv.Atoi(fields[3])
        adjusted_turn := float64(turn) + 0.5 - 0.5*float64(fast)
        summary[adjusted_turn] += 1
    }
    return summary, nil
}

func get_summary_header() string {
    header := ""
    for _, turn := range turns {
        header += fmt.Sprintf("          T%4.1f", turn)
    }
    return header
}

func get_summary_line(summary map[float64]int) string {
    ntot := 0
    for _, n := range summary { ntot += n }
    ncum := 0
    line := ""
    for _, turn := range turns {
        ncum += summary[turn]
        dn := math.Sqrt(float64(ncum))
        p := float64(100*ncum)/float64(ntot)
        dp := 100.*dn/float64(ntot)
        line += "     " + fmt.Sprintf("%3.0f", p) + "%" + " ± " + fmt.Sprintf("%2.0f", dp) + "%"
    }
    return line
}
