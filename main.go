
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

//        for _, line := range state.Log {
//            fmt.Println(line)
//        }


    } else {
        fmt.Println(err)
    }

}
