package main

import (
  "fmt"
  "log"
)
type Package struct {
  name string
  dep string
}

type Packs []Package
type Mapper map[string]Package

var pack1,pack2 Packs
var counter int
var innerpack []Package
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
  packages:=newpack()
 chOut:=inspect(packages,0)
  fmt.Println(chOut)
}

func inspect(packets Packs, level int) Packs {
  if level!=3{
    level++
    switch level {
    case 1:
      packets = checknone(packets)
      break
    case 2:
      packets = checkdependency(packets)
      break
    case 3:
      pack1 =append([]Package{}, checkOut(packets)...)
      return pack1
    default:
      log.Fatal("Something happened....")
    }} else { return packets}

  inspect(packets, level)
  return pack1
} //Recursive

func checknone(first Packs)(Packs){
  var update []Package
  fmt.Println("SIZE: ",len(first), first,"\nLets put blank dependencies at top of list:")
  _, err := fmt.Scanf("\n")
  CheckErr(err)

  for i, v:= range first{
    if i==0 && v.dep=="" {
      RR=append(RR, v)
      update=append([]Package{},first[i+1:]...)
      fmt.Println("Added: ", v)
    } else if v.dep=="" && i!=len(first)-1{
      RR=append(RR, v)
      update=append([]Package{},update[i:]...)
      fmt.Println("Added: ", v)
    }else if i==len(first)-1 && v.dep==""{
      RR=append(RR, v)
      update=append([]Package{},update[:i-1]...)
      fmt.Println("Added: ", v)
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
  fmt.Println("SIZE: ",length, update,"\nLets clear dependencies:")
  _, err := fmt.Scanf("\n")
  CheckErr(err)

  for i, v:= range second{
    search:=searchresults(mapper(RR), v.dep)
    if i==0 && search==false{ //SKIP
      last:=second[0]
      update=append([]Package{},update[i+1:]...)
      update=append(update, last)
    } else if i==0 && search==true {
      RR=append(RR, v)
      update=append([]Package{},update[i+1:]...)
      fmt.Println("Added: ", v)
    } else if i<(length-1) && search==true {
      RR=append(RR, v)
      update=append([]Package{},update[1:]...)
      fmt.Println("Added: ", v)
    }else if i==(length-1) && search==true{
      RR=append(RR, v)
      update=append(update[:i],update[i+1:]...)
      fmt.Println("Added: ", v )
    } else {}}

  return update
}
func checkOut(third Packs) (Packs){
  result:=append(RR, third...)
  fmt.Println( "\nFully Sorted: ")
  _, err := fmt.Scanf("\n")
  CheckErr(err)
  return result
}
func CheckErr(err error){
  if err != nil {log.Fatal(err)}
}
