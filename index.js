
var packages = [
  {
    name: "kitten",  //1
    dep: ""
  },
  {
    name: "leetmeme",
    dep: "cyberportal" //
  },
  {
    name: "cyberportal",
    dep: "ice"       // 2
  },
  {
    name: "camel",
    dep: "kitten"    //2
  },
  {
    name: "cook",
    dep: "leetmeme"
  },
  {
    name: "ice",  //1
    dep: ""
  }
]
var result = []
var pac2=[...packages]
function getInstallOrder(packages) {
  console.log(packages)
  inspect(packages)
  return results
  //return ["kitten", "ice", "cyberportal", "leetmeme", "camel", "cook"]
}

function inspect(...packages){
  var results
  for (var i in packages){
    if (i.dep==""){
      result.push(i)
      let index=packages.indexOf(i)
      if (index>-1){
        pac2.splice(index,1)}
      break} else{
      for (var c in results){
        if(c.name==i.dep){
          results.push(i)
          let index=packages.indexOf(i)
          if (index>-1){
            pac2.splice(index,1)
            break
          } else{
            results.push(i)
            let index=packages.indexOf(i)
            if (index>-1){
              pac2.splice(index,1)
              break
            }
          }
        }
      }
    }}
    inspect(pac2)
  }
  console.log(getInstallOrder(packages))
