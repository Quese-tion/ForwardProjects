package main

import (
  "fmt"
  "testing"
)

func mainTest(t *testing.T){

  t.Log("Hello")
}

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
    if (i==0){
      if (v.dep==""){
        results=append(results, v)
        pac2=append(pac2[:1],pac2[1:]...)
      } else {
        oldone:=v
        pac1=append(pac1[:1],pac1[:1]... )
        pac1=append(pac1, oldone)
      }
      //need to delete element from struct slice
    } else if searchresults(mapper(results),v) { //Located in results
      results=append(results, v)
      pac2=append(pac2[i:],pac2[:i+1]...)
      //need to delete element from struct slice
    } else{
      for j, r:= range results{ //First
        if(r.name==v.dep){//z new package v
          results=append(results,v)
          pac2=append(pac2[:j],pac2[j+1:]... )
        } else{
          results=append(results,v)
      } }}
  }
  fmt.Println("Results: ", results)
  inspect(pac1,pac2)
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
  return err
}
