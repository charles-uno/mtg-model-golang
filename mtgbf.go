package main

import (
    "fmt"
    "mtgbf/mtgbf"
)


func main() {
    fmt.Println("hello")

    m := mtgbf.Mana("GW2URR")

    fmt.Println(m)

    fmt.Println(m.Pretty())

    c := mtgbf.Card("Forest")
    fmt.Println(c)

    c = mtgbf.Card("Ancient Stirrings")
    c = mtgbf.Card("Ancient Stirrings")
    c = mtgbf.Card("Ancient Stirrings")
    fmt.Println(c.Pretty())


    cards := mtgbf.Cards("Explore", "Opt", "Explore")
    fmt.Println(cards.Pretty())

}
