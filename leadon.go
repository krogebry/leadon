package main

import (
  "os"
  "io"
  "io/ioutil"
  //"fmt"
  //"html"
  "log"
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"
)

/**
type EmailEndpoint struct {
  Host string
  Mailbox string
}

type Message struct {
  From []EmailEndpoint
  Subject string
}
*/

func register(res http.ResponseWriter, req *http.Request) {
  log.Printf( "Registering user" )
  //res.Header().Set( "Content-Type", "application/json" )

  var user User

  body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := req.Body.Close(); err != nil {
    panic(err)
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")

  if err := json.Unmarshal(body, &user); err != nil {
    res.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder( res ).Encode(err); err != nil {
      panic(err)
    }
  }

  log.Printf( "%s", body )
  log.Printf( "%s", user )

  //t := RepoCreateTodo(todo)
  //log.Printf( "%s", user )

  //if err := dbconn.Find( )
  //result := &User{}
  count := 0
  count, err = dbconn.Find(bson.M{"id": user.Id }).Count()
    //res.WriteHeader(503) 
    //if err := json.NewEncoder( res ).Encode(err); err != nil {
      //panic(err)
    //}
    //panic( "User already exists!" )
  //}

  if count > 0 {
    res.WriteHeader(422) 
    log.Printf( "Found docs: %i", count )
    if err := json.NewEncoder(res).Encode( `{ "Success": "false" }` ); err != nil {
      panic(err)
    }

  }else{
    dbconn.Insert( user )
  //res.Header().Set("Content-Type", "application/json; charset=UTF-8")
    res.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(res).Encode( user ); err != nil {
      panic(err)
    }

  }
}

type User struct {
  Id string
  Name string
}

func users(res http.ResponseWriter, req *http.Request) {
  res.Header().Set( "Content-Type", "application/json" )
  vars := mux.Vars(req)
  userId := vars["userId"]
  log.Printf( "User: %s", vars )

  user := &User{}
  user.Id = userId

  json.NewEncoder( res ).Encode( user )
}

var dbconn *mgo.Collection

func main() {
  log.SetOutput( os.Stdout )
  log.Printf( "Starting" )

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic( "Unable to connect" )
  }
  defer session.Close()
  dbconn = session.DB("leadon").C("users")

  router := mux.NewRouter().StrictSlash(true)

  //router.Methods( "GET" ).Path( "/api/users" ).Name( "Users" ).Handler( users )

  router.HandleFunc("/api/users/{userId}", users)
  router.HandleFunc("/api/register", register)

  log.Fatal(http.ListenAndServe(":9000", router))


  //http.HandleFunc("/api/register", register)
  //http.ListenAndServe(":9000", nil)

  /**
  result := Message{}
  err = c.Find(bson.M{"message_id": "<bcaec51ba4a71e9b0104f19ad8e5@google.com>"}).One(&result)

  if err != nil {
    log.Fatal( "Unable to find row" )
  }

  fmt.Printf("%s says %s\n", result.From[0].Mailbox, result.Subject)
  */
}

