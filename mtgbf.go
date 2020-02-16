package main

import (
    "fmt"
    "mtgbf/mtgbf"
)


func main() {

//    m := mtgbf.Mana("GW2URR")
//    fmt.Println(m.Pretty())

//    c := mtgbf.Card("Forest")
//    fmt.Println(c)

    state := mtgbf.InitialGameState("debug")

    fmt.Println(state.Pretty())

//    cards := mtgbf.Cards("Explore", "Opt", "Explore")
//    fmt.Println(cards.Pretty())
}
