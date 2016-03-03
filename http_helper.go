package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func ParsePostBody(res http.ResponseWriter, req *http.Request) []byte {
    // Read body
    data, err := ioutil.ReadAll(req.Body)
    if err != nil {
        BadRequest(&res, err.Error())
        return nil
    }
    return data
}

func WriteJSON(res http.ResponseWriter, element interface{}) {
    jsonElement, err := json.Marshal(&element)
    if err != nil {
        HttpError(&res, "Error converting to JSON", 500, err.Error())
        return
    }
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(jsonElement)
}

func HttpError(res *http.ResponseWriter, errMsg string, errCode int, logMsg string) {
    Error.Println(errMsg, logMsg)
    http.Error(*res, errMsg, errCode)
}

func BadRequest(res *http.ResponseWriter, errMsg string) {
    HttpError(res, errMsg, 400, "")
}

func EmptyResponse(res *http.ResponseWriter, errMsg string) {
    HttpError(res, "", 204, errMsg)
}

func NotFound(res *http.ResponseWriter, errMsg string) {
    HttpError(res, errMsg, 404, "")
}

func InternalServerError(res *http.ResponseWriter, errMsg string) {
    HttpError(res, errMsg, 500, "")
}
