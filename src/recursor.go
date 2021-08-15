package main

import (
  "fmt"
)
type Mapper map[string]Package
func main() {
  names := []string{"kitten", "leetmeme", "cyberportal", "camel", "cook", "ice"}
  deps := []string{"", "cyberportal", "ice", "kitten", "leetmeme", ""}
  packages := []Package{}
  for x := 0; x < len(names); x++ {
    packages = append(packages, Package{names[x], deps[x]})
  }

  pack1 = append(pack1, packages...) // think I ran out of memory take remotely...
  fmt.Println(packages, "\n-----------------------------\n")
  inspect(packages, pack1)
}

func inspect(pac1 Packs, pac2 Packs){ //Delete from original then recurse
  for i,v :=range pac1{
    fmt.Println("Interation ", i)
    if (i==0){
      if (v.dep==""){
        fmt.Println("Hello ", pac2)
        results=append(results, v)
        pac2=append(pac2[1:],pac2[1:]...)
        break
      } else {
        oldone:=v
        pac1=append(pac1[:1],pac1[:1]... )
        pac1=append(pac1, oldone)
      }
      //need to delete element from struct slice
    } else if i>0{
      if v.dep==""{
        results=append(results, v)
        pac2=append(pac2[:1],pac2[1:]...)
      }else
      if searchresults(mapper(results),v) { //Located in results
        results=append(results, v)
        pac2=append(pac2[i:],pac2[:i+1]...)
        //need to delete element from struct slice
      }else{
          results=append(results,v)
          fmt.Println("Results: ", results, "\nLeft: ", pac2)
          return
      } }
  }
  fmt.Println("Results: ", results, "New list : ", pac2)
 //inspect(pac1,pac2)
}

func mapper(packs []Package) map[string]Package{
  mapout:=make(map[string]Package)
  for _,y:= range packs{
    mapout[y.name]=y
  }
  return mapout
}
func searchresults(m Mapper, p Package)(bool){
  _,err:=m[p.name]
  fmt.Println("Searchresults for : ", m, "| Checking for: ",p, "| Showed: ", err)
  return err
}
