
package valakut

// ---------------------------------------------------------------------

import (
    "fmt"
    "math/rand"
    "time"
)

// ---------------------------------------------------------------------

func InitRandom() {
    rand.Seed(time.Now().UTC().UnixNano())
}

// ---------------------------------------------------------------------

func contains(arr []string, elt string) bool {
    for _, a := range arr {
        if a == elt { return true }
    }
    return false
}



func b2i2a(b bool) string {
    if b { return "1" } else { return "0" }
}



func count(arr []string, r string) (n int) {
    for _, a := range arr {
        if a == r { n += 1 }
    }
    return n
}

// ---------------------------------------------------------------------

func remove(arr []string, elt string) []string {
    for i, a := range arr {
        if a == elt {
            arr[i] = arr[len(arr)-1]
            new_arr := arr[:len(arr)-1]
            return new_arr
        }
    }
    fmt.Println("WARNING: Failed to remove", elt, "from", arr)
    return []string{}
}

// ---------------------------------------------------------------------

func shuffled(deck []string) []string {
    ret := make([]string, len(deck))
    for i, j := range rand.Perm(len(deck)) { ret[i] = deck[j] }
    return ret
}

// ---------------------------------------------------------------------

func flip() bool {
    return rand.Intn(2) == 0
}

// ---------------------------------------------------------------------

func unique_strings(arr []string) []string {
    counts := make(map[string]int)
    for _, card := range arr { counts[card] += 1 }
    uarr := []string{}
    for card, _ := range counts { uarr = append(uarr, card) }
    return uarr
}
