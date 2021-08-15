package main

import "fmt"


type Package struct {
  name string
  dep string
}

type Packs []Package
var pack1,pack2 Packs
var results, test []Package


func mai() {
  names := []string{"kitten", "leetmeme", "cyberportal", "camel", "cook", "ice"}
  deps := []string{"", "cyberportal", "ice", "kitten", "leetmeme", ""}
  packages := []Package{}
  for x := 0; x < len(names); x++ {
    packages = append(packages, Package{names[x], deps[x]})
  }

  pack1 = append(pack1, packages...) // think I ran out of memory take remotely...
  fmt.Println(packages, "\n-----------------------------\n")
  packs:= make(chan Package)
  for range packages{
    go sortpackages(packages, packs, results )
    test=append(test, <-packs)
  }
  fmt.Println("Results",test)
}

func sortpackages(packages Packs, pac chan Package, result Packs ){
  fmt.Println(packages)
  for _, pack := range packages{
    fmt.Println(pack.name, "\n", pack.dep)
      if pack.dep=="" {
        //results=append(results, pack)
       // results = append(results[:i], results[i+1:]...)
        pac<-pack
      } else{
        for _, result:= range results{ //First
          if result.name==pack.dep { //z new package v
          //  results=append(results,pack)
          //  results = append(results[:i], results[i+1:]...)
            pac<-pack
          } else{
          //  results=append(results,pack)
          //  results = append(results[:i], results[i+1:]...)
            pac<-pack
            }} }
  }
}

func sortpacks(packages Packs, results Packs)  {

}

func ninspect(packages Packs,pac chan Package ){ //Delete from original then recurse
  for i,v :=range packages{
    fmt.Println(v)
    i := i
    go func(pack Package) {
      fmt.Println(pack)
      if pack.dep=="" {
        results=append(results, pack)
        results = append(results[:i], results[i+1:]...)
        pac<-pack
        return
      } else{
        for i, result:= range results{ //First
          if result.name==pack.dep { //z new package v
            results=append(results,pack)
            results = append(results[:i], results[i+1:]...)
            pac<-pack
            return
          } else{
            results=append(results,pack)
            results = append(results[:i], results[i+1:]...)
            pac<-pack
            return}} }
      fmt.Println(results)
    }(v)
}
  fmt.Println(results)
}

func (p *Packs) deleteElement(){
  fmt.Println("Yee")
}
