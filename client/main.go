package main

import (
  "fmt"
  "log"
  "net/rpc"
)

type Item struct{
  Title string
  Body string
}

func main(){
  var reply Item
  var db []Item

  client, err := rpc.DialHTTP("tcp","localhost:4040")

  if err!=nil{
    log.Fatal("connection error: ",err)
  }

  a := Item{"first","A first item"}
  b := Item{"second","A second item"}
  c := Item{"third", "A third item"}

  client.Call("API.AddItem",a,&reply)
  client.Call("API.AddItem",b,&reply)
  client.Call("API.AddItem",c,&reply)
  client.Call("API.GetDB","",&db)

  fmt.Println("database : ",db)

}
