package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func addHealthData(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    // Read body
    data := ParsePostBody(res, req)
    // unmarshal json
    hd, err := marshalHealthData(data)
    if err != nil {
        HttpError(&res, err.Error(), 400, err.Error())
        return
    }
    // save to db
    id, err := saveHealthData(hd)
    if err != nil {
        HttpError(&res, err.Error(), 500, err.Error())
        return
    }
    res.Header().Set("Location", "/api/v0/data/"+id)
    res.WriteHeader(http.StatusCreated)
}

func getHealthData(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    healthData, err := findHealthData()
    if err != nil {
        switch err.(type) {
        case NotFoundError:
            EmptyResponse(&res, "")
            return
        default:
            InternalServerError(&res, "Internal Server Error")
            return
        }
    }
    WriteJSON(res, healthData)
}

func deleteHealthData(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    healthDataId := params.ByName("HealthDataId")
    if err := removeHealthData(healthDataId); err != nil {
        switch err.(type) {
        case NotFoundError:
            NotFound(&res, "HealthData with id "+healthDataId+" not found")
            return
        default:
            InternalServerError(&res, "Internal Server Error")
            return
        }
    }
    res.WriteHeader(http.StatusNoContent)
}
