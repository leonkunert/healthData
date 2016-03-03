package main

import r "github.com/dancannon/gorethink"

var (
    Rethink *r.Session
    err error
)

func initDBConnection() {
    Rethink, err = r.Connect(r.ConnectOpts{
        Address: CFG.UString("db.rethink.host"),
        Database: CFG.UString("db.rethink.database"),
        MaxIdle: CFG.UInt("db.rethink.maxIdle"),
        MaxOpen: CFG.UInt("db.rethink.maxOpen"),
    })
    if err != nil {
        Error.Println(err.Error())
    }
    Info.Println("Connected to RethinkDB")
}
