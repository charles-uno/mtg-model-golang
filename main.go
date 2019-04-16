
package main

import (
    "fmt"
    "os"
    "strconv"
    "valakut-go/valakut"
)

// ---------------------------------------------------------------------

func main() {

    valakut.InitRandom()

    name := "debug"

    ntrials := 1
    if len(os.Args) > 1 { ntrials, _ = strconv.Atoi(os.Args[1]) }

    filename := "data/" + name + ".out"

    for i := 0; i < ntrials; i++ {
        state, err := valakut.Simulate(name)
        if err == nil {
            fmt.Print(string(state.Line()))
            valakut.Append(filename, state.Line())
            if i == ntrials-1 { fmt.Println("\n" + state.Log) }
        } else {
            valakut.Append(filename, "\n")
            fmt.Println(err)
        }
    }

}
