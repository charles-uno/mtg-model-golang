package mtgbf


import (
    "errors"
    "fmt"
    "github.com/smallfish/simpleyaml"
    "io/ioutil"
    "strings"
)


type card_obj struct {
    colors []string
    cost mana_obj
    display string
    name string
    types []string
}


// Keep card metadata cached so we don't have to keep looking it up
var card_metadata = make(map[string]card_obj)


func Card(name string) card_obj {
    _, ok := card_metadata[name]
    if ! ok {
        fmt.Println("creating new card:", name)
        carddata = get_carddata()
        // Card types
        types_raw, err := carddata.Get(name).Get("types").String()
        if err != nil { panic(errors.New("no types for " + name)) }
        // Colors
        colors_raw, err := carddata.Get(name).Get("colors").String()
        if err != nil { colors_raw = "" }
        // Short name
        display, err := carddata.Get(name).Get("display").String()
        if err != nil { display = slug(name) }
        // Mana cost
        cost_raw, err := carddata.Get(name).Get("cost").String()
        if err != nil { cost_raw = "" }
        // Stick this all into a card object
        card_metadata[name] = card_obj{
            colors: strings.Split(colors_raw, ","),
            cost: Mana(cost_raw),
            display: display,
            name: name,
            types: strings.Split(types_raw, ","),
        }
    }
    return card_metadata[name]
}


func (c *card_obj) Pretty() string {
    return c.display
}


func slug(name string) string {
    for _, r := range " -'," {
        name = strings.ReplaceAll(name, string(r), "")
    }
    return name
}

var carddata *simpleyaml.Yaml = nil


func get_carddata() *simpleyaml.Yaml {
    if carddata == nil {
        fmt.Println("loading card metadata")
        source, err := ioutil.ReadFile("carddata.yaml")
        if err != nil { panic(err) }
        carddata, err = simpleyaml.NewYaml(source)
        if err != nil { panic(err) }
    }
    return carddata
}
