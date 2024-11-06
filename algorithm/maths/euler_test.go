package maths

import (
    "math/rand"
    "testing"
)


func TestEuler(t *testing.T) {
    maxn := 1000000
    n := rand.Intn(maxn) + 1
    t.Logf("TestEuler(%d)", n)
    a, c := Euler(n)
    b := BruteForce(n)
    for i:=2; i<=n; i++ {
        if a[i] != b[i] {
           t.Fatal("a[i] != b[i]") 
        }
    }
    for _, x := range c {
        if b[x] {
            t.Fatal("b[x] is NOT a prime number")
        }
    }
}
