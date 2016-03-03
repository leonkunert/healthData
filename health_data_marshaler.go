package main

import (
    "encoding/json"
    "time"
)

func marshalHealthData(data []byte) (*HealthData, error) {
    hd := new(HealthData)
    err := json.Unmarshal(data, hd)
    if err != nil {
        return nil, JsonInvalidError{}
    }
    hd.Timestamp = time.Now()
    return hd, nil
}
