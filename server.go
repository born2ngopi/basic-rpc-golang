package main

import (
  "net"
  "net/http"
  "net/rpc"
  "log"
)

type Item struct{
  Title string
  Body string
}

type API int

var database []Item

func (a *API) GetDB(title string, reply *[]Item) error{
  *reply = database
  return nil
}

func (a *API) GetByName(title string, reply *Item) error{
  var getItem Item

  for _, val := range database{
    if val.Title == title{
      getItem = val
    }
  }

  *reply = getItem

  return nil
}

func (a *API) AddItem(item Item,reply *Item)error{
  database = append(database,item)
  *reply = item
  return nil
}

func (a *API) EditItem(edit Item, reply *Item)error{
  var changed Item

  for idx, val := range database{
    if val.Title == edit.Title{
      database[idx]=Item{edit.Title,edit.Body}
      changed = database[idx]
    }
  }
  *reply = changed
  return nil
}

func (a *API) DeleteItem(item Item, reply *Item)error{
  var del Item

  for idx, val := range database{
    if val.Title == item.Title && val.Body == item.Body{
      database = append(database[:idx], database[idx+1:]...)
      del = item
      break
    }
  }
  *reply = del
  return nil
}

func main(){
  var api = new(API)
  err := rpc.Register(api)

  if err != nil{
    log.Fatal("error registering API",err)
  }

  rpc.HandleHTTP()

  listener, err := net.Listen("tcp",":4040")
  if err!=nil{
    log.Fatal("listener error ",err)
  }

  log.Printf("serving rpc on port 4040")
  err = http.Serve(listener,nil)
  if err!=nil{
    log.Fatal("error serving :",err)
  }




  //fmt.Println("initial database: ",database)
  //a := Item{"first","a test item"}
 // b := Item{"second","a second item"}
  //c := Item{"third", "a third item"}

  //AddItem(a)
  //AddItem(b)
  //AddItem(c)
  //fmt.Println("second database : ", database)

  //DeleteItem(b)
  //fmt.Println("third database ",database)

  //EditItem("third",Item{"fourth","a new item"})
  //fmt.Println("fourth database ",database)

  //x := GetByName("fourth")
  //y := GetByName("first")
  //fmt.Println(x,y)
}
