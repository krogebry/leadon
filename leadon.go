package main

import (
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type EmailEndpoint struct {
  Host string
  Mailbox string
}

type Message struct {
  From []EmailEndpoint
  Subject string
}

func main() {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic( "Unable to connect" )
  }
  defer session.Close()

  c := session.DB("email").C("raw")

  result := Message{}
  err = c.Find(bson.M{"message_id": "<bcaec51ba4a71e9b0104f19ad8e5@google.com>"}).One(&result)

  if err != nil {
    log.Fatal( "Unable to find row" )
  }

  fmt.Printf("%s says %s\n", result.From[0].Mailbox, result.Subject)
}

