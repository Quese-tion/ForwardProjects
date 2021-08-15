package test

import (
  "fmt"
  "testing"
)

var err error
func TestLoop(m *testing.T) {

  err =nil
  if err !=nil {m.Fatal(err)}
}
func TestInspect(){
  packages:=newpack()
  pack1 = newpack()
  fmt.Println(packages, "\n-----------------------------\n")
  chNone:=checknone(pack1)
  chDepend:=checkdependency(chNone)
  chOut:=checkOut(chDepend) //chOut
  fmt.Println(  chOut)
} //Test-seperately

