package mtgbf


import (
    "errors"
    "fmt"
    "github.com/smallfish/simpleyaml"
    "io/ioutil"
    "strings"
)


type card_data struct {
    colors []string
    cost mana
    display string
    enters_tapped bool
    name string
    taps_for mana
    types []string
}


var cards_data = make(map[string]card_data)


func is_land(card string) bool {
    for _, t := range get_card_data(card).types {
        if t == "land" { return true }
    }
    return false
}


func enters_tapped(card string) bool {
    return get_card_data(card).enters_tapped
}


func taps_for(card string) mana {
    return get_card_data(card).taps_for
}



func get_lands(cards []string) []string {
    lands := []string{}
    for _, card := range cards {
        if is_land(card) {
            lands = append(lands, card)
        }
    }
    return unique_strings(lands)
}


func get_display(card string) string {
    return get_card_data(card).display
}


func pretty(cards ...string) string {
    names := []string{}
    for _, card := range cards {
        names = append(names, get_display(card))
    }
    return tally_strings(names)
}


func get_card_data(card string) card_data {
    _, ok := cards_data[card]
    if ! ok {
        fmt.Println("creating new card:", card)
        cdy := get_card_data_yaml()
        // Card types
        types_raw, err := cdy.Get(card).Get("types").String()
        if err != nil { panic(errors.New("no types for " + card)) }
        // Colors
        colors_raw, err := cdy.Get(card).Get("colors").String()
        if err != nil { colors_raw = "" }
        // Short name
        display, err := cdy.Get(card).Get("display").String()
        if err != nil { display = slug(card) }
        // Mana cost
        cost_raw, err := cdy.Get(card).Get("cost").String()
        if err != nil { cost_raw = "" }
        // For lands, does it enter the battlefield tapped?
        enters_tapped, err := cdy.Get(card).Get("enters_tapped").Bool()
        if err != nil { enters_tapped = false }
        taps_for_raw, err := cdy.Get(card).Get("taps_for").String()
        if err != nil { taps_for_raw = "" }
        // Stick this all into a card object
        cards_data[card] = card_data{
            colors: strings.Split(colors_raw, ","),
            cost: Mana(cost_raw),
            display: display,
            enters_tapped: enters_tapped,
            name: card,
            taps_for: Mana(taps_for_raw),
            types: strings.Split(types_raw, ","),
        }
    }
    return cards_data[card]
}


func slug(name string) string {
    for _, r := range " -'," {
        name = strings.ReplaceAll(name, string(r), "")
    }
    return name
}


var card_data_yaml *simpleyaml.Yaml = nil


func get_card_data_yaml() *simpleyaml.Yaml {
    if card_data_yaml == nil {
        fmt.Println("loading card metadata")
        source, err := ioutil.ReadFile("carddata.yaml")
        if err != nil { panic(err) }
        card_data_yaml, err = simpleyaml.NewYaml(source)
        if err != nil { panic(err) }
    }
    return card_data_yaml
}
