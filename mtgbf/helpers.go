package mtgbf


import (
    "errors"
    "sort"
    "strconv"
    "strings"
)


func contains(arr []string, elt string) bool {
    for _, a := range arr {
        if a == elt { return true }
    }
    return false
}


func discard(arr_old []string, elt string) []string {
    if ! contains(arr_old, elt) {
        panic(errors.New(elt + " not in " + tally_strings(arr_old)))
    }
    arr_new := arr_old[:len(arr_old) - 1]
    for i, a := range arr_new {
        if a == elt {
            arr_new[i] = arr_old[len(arr_old) - 1]
            break
        }
    }
    return arr_new
}


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


func count_strings(arr []string, elt string) int {
    n := 0
    for _, a := range arr {
        if a == elt { n += 1 }
    }
    return n
}


func count_runes(s string, r rune) int {
    count := 0
    for _, char := range s {
        if char == r { count += 1 }
    }
    return count
}


func unique_counts(arr []string) map[string]int {
    uc := make(map[string]int)
    for _, a := range arr { uc[a] += 1 }
    return uc
}
