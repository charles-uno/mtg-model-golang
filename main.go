
package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "valakut/valakut"
)

// ---------------------------------------------------------------------

func main() {

    valakut.InitRandom()

    ntrials := 0
    if len(os.Args) > 1 {
        ntrials, _ = strconv.Atoi(os.Args[1])
    }

    if ntrials == 0 {
        valakut.PrintSummaries()
        return
    }

    names := []string{}
    if len(os.Args) > 2 {
        names = os.Args[2:]
    }

    namefmt := "%-" + strconv.Itoa(maxlen(names)+2) + "s"
    numfmt := ""
    if len(os.Args) > 1 {
        numfmt = "%-" + strconv.Itoa(len(os.Args[1])+2) + "d"
    }

    for i := 0; i < ntrials; i++ {

        if ntrials == 0 || len(names) == 0 { break }

        name := choice(names)

        fmt.Printf(numfmt, i)
        fmt.Printf(namefmt, name)

        state, err := valakut.Simulate(name)
        // It'll be all zeros if we failed to converge.
        valakut.SaveResult(name, state.Line())

        if err == nil {
            fmt.Println(state.Line())
            if i == ntrials-1 { fmt.Println("\n" + state.Log) }
        } else {
            fmt.Println(err)
        }
    }
}

// ---------------------------------------------------------------------

func choice(names []string) string {
    return names[rand.Intn(len(names))]
}


func maxlen(names []string) int {
    ml := 0
    for _, name := range names {
        if len(name) > ml { ml = len(name) }
    }
    return ml
}
