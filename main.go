
package main

import (
    "fmt"
    "os"
    "strconv"
    "valakut/valakut"
)

// ---------------------------------------------------------------------

func main() {

    valakut.InitRandom()

    name := "debug"

    ntrials := 1
    if len(os.Args) > 1 { ntrials, _ = strconv.Atoi(os.Args[1]) }

    for i := 0; i < ntrials; i++ {
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

    fmt.Println("")

    valakut.PrintSummary(name)



}
