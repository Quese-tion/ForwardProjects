package main

import (
  "fmt"
)
type Package struct {
  name string
  dep string
}

type Packs []Package
type Mapper map[string]Package

var pack1,pack2 Packs
var RR []Package

//Helper Classes
func newpack() (Packs) {
  names := []string{"kitten", "leetmeme", "cyberportal", "camel", "cook", "ice"}
  deps := []string{"", "cyberportal", "ice", "kitten", "leetmeme", ""}
  packages := []Package{}
  for x := 0; x < len(names); x++ {
    packages = append(packages, Package{names[x], deps[x]})
  }
  return packages
} //Creates Dataset
func mapper(packs []Package) map[string]Package{
  mapout:=make(map[string]Package)
  for _,y:= range packs{
    mapout[y.name]=y
  }
  return mapout
} //Convert Slice to Map
func searchresults(m Mapper, p string)(bool){
  _,err:=m[p]
  //fmt.Println("Searchresults for : ", m, "| Checking for: ",p, "| Showed: ", err)
  return err
}    //Searches map for value

func main() {
  TestInspect()
}

func TestInspect(){
  packages:=newpack()
  pack1 = newpack()
  fmt.Println(packages, "\n-----------------------------\n")
  chNone:=checknone(pack1)
  chDepend:=checkdependency(chNone)
  chOut:=checkOut(chDepend)
  fmt.Println("Old List: ",chNone, "\nNew List: ", chDepend, "\nResults: ", chOut)
}

func checknone(first Packs)(Packs){
  var update []Package
  fmt.Println("SIZE: ",len(first), first)
  for i, v:= range first{
    if i==0 && v.dep=="" {
      RR=append(RR, v)
      update=append([]Package{},first[i+1:]...)
      fmt.Println("Added: ", v, "In First 0: ", update )
    } else if v.dep=="" && i!=len(first)-1{
      RR=append(RR, v)
      update=append([]Package{},update[i:]...)
      fmt.Println("Added: ", v, "In First not end: ", update )
    }else if i==len(first)-1 && v.dep==""{
      RR=append(RR, v)
      update=append([]Package{},update[:i-1]...)
      fmt.Println("Added: ", v, "First that is end: ", update )
    } else {
      last:=first[0]
      update=append(update,first[1:]...)
      update=append(update, last)}}

  return update
}
func checkdependency(second Packs)(Packs){
  var update []Package
  update=append(update,second...)
  length:=len(second)
  fmt.Println("SIZE: ",length, update)
  for i, v:= range second{
    search:=searchresults(mapper(RR), v.dep)
    if i==0 && search==false{ //SKIP
      last:=second[0]
      update=append([]Package{},update[i+1:]...)
      update=append(update, last)
    } else if i==0 && search==true {
      RR=append(RR, v)
      update=append([]Package{},update[i+1:]...)
      fmt.Println("Added: ", v, "In First 0: ", update )
    } else if i<(length-1) && search==true {
      fmt.Println("INDEX: ",i," Bring: ", update)
      RR=append(RR, v)
      update=append([]Package{},update[1:]...)
    }else if i==(length-1) && search==true{
      RR=append(RR, v)
      update=append(update[:i],update[i+1:]...)
      fmt.Println("Added: ", v, "First that is end: ", update )
    } else {}
    fmt.Println("Interation ", i, "Update: ", update)
  }

  return update
}
func checkOut(third Packs) (Packs){
  result:=append(RR, third...)
  return result
}
