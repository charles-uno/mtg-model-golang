package mtgbf


import (
    "sort"
    "strconv"
    "strings"
)


func tally_strings(arr []string) string {
    uc := unique_counts(arr)
    keys := []string{}
    for key, _ := range uc { keys = append(keys, key) }
    sort.Strings(keys)
    tallies := []string{}
    for _, key := range keys {
        n := uc[key]
        if n == 1 {
            tallies = append(tallies, key)
        } else {
            tallies = append(tallies, strconv.Itoa(n) + "*" + key)
        }
    }
    return strings.Join(tallies, " ")
}


func unique_strings(arr []string) []string {
    us := []string{}
    for a, _ := range unique_counts(arr) {
        us = append(us, a)
    }
    return us
}


func unique_counts(arr []string) map[string]int {
    uc := make(map[string]int)
    for _, a := range arr { uc[a] += 1 }
    return uc
}
