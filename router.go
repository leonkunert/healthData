package main

import (
    "github.com/julienschmidt/httprouter"
    "os"
    "os/signal"
    "syscall"
    "net/http"
)

var router *httprouter.Router

type Status struct {
    status string `json:"status"`
}

func initRouter()  {
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

    // start the router
    go func() {
        router = httprouter.New()
        router.HandleMethodNotAllowed = true
        router.RedirectFixedPath = true
        router.HandleOPTIONS = true

        /******* /menus/ *******/
        // Get all data
        router.GET("/api/v0/data/", getHealthData)
        // Add new data
        router.POST("/api/v0/data/", addHealthData)
        // Delete a data entry
        router.DELETE("/api/v0/data/:HealthDataId/", deleteHealthData)

        /******* Status Handler *******/
        router.GET("/api/v0/status/", func (res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
            Info.Println("Sending Status Report")
            WriteJSON(res, Status{"Ok"})
        })

        port := CFG.UString("httpPort")
        Info.Println("Started Server at Port "+port)
        if err := http.ListenAndServe(":"+port, router); err != nil {
            Error.Println("ListenAndServer", err.Error())
            panic("Address Already In Use")
        }
    }()

    <-sigchan
}
