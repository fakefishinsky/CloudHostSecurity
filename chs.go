package main

/*
  Cloud Host Security
*/

import (
    "fmt"
    comp "modules/compliance"
)

func main() {
    policies := &comp.Policy{}
    err := policies.LoadPolicy("policy.json")
    if err != nil {
        fmt.Println("Error: ", err)
    }
    policies.Audit()
}