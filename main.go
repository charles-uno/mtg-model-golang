
package main

import (
    "fmt"
    "valakut-go/valakut"
)

// ---------------------------------------------------------------------

func main() {

    state, err := valakut.Simulate("debug")

    if err == nil {
        fmt.Println(state.Log)
    } else {
        fmt.Println(err)
    }

}
