package main

import (
  //"net/url"
  "io"
  "io/ioutil"
  //"fmt"
  "log"
  //"github.com/bndr/gopencils"
  "strings"
  "net/http"
  "encoding/json"
  //"gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

type User struct {
  Id string
  Name string
}

type UserCreateResponse struct {
  Success string
}

func main() {
  //resp, err := http.PostForm("http://localhost:9000/api/register", url.Values{"key": {"id"}, "id": { 321 }})

  //user := &User{ Id: "321" }

  client := &http.Client{}

  //resp, err := client.PostForm("http://localhost:9000/api/register", url.Values{"key": {"id"}, "id": { "321" }})
  //json_str,err := json.Marshal( user ) 

  json_str := `{"id":"123"}`

  resp, err := client.Post("http://localhost:9000/api/register", "application/json", strings.NewReader( json_str ))
  if err != nil {
    panic( err )
  }

  log.Printf( "Response: %i", resp.StatusCode )

  body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := resp.Body.Close(); err != nil {
    panic(err)
  }
  log.Printf( "%s\n", body )

  //var f interface{}
  f := UserCreateResponse{}

  if err := json.Unmarshal(body, &f); err != nil {
    panic(err)
  }

  //log.Printf( "%s", f )

  log.Printf( "Success: %s", f.Success )

  resp.Body.Close()
}
