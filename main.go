package main

import "os"

func main()  {
    initLogging(os.Stdout, os.Stderr)
    loadConfig()
    initDBConnection()
    defer Rethink.Close()
    initRouter()
}
