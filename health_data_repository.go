package main

import (
    r "github.com/dancannon/gorethink"
    "strings"
)

var healthDataTable = r.Table("healthData")

func saveHealthData(healthData *HealthData) (string, error) {
    Info.Println("Writing healthData to DB")
    resp, err := healthDataTable.Insert(&healthData).RunWrite(Rethink)
    if err != nil {
        return "", err
    }
    Info.Printf("%d row inserted with id %v\n", resp.Inserted, resp.GeneratedKeys)
    return strings.Join(resp.GeneratedKeys, ""), nil
}

func findHealthData() ([]HealthData, error) {
    Info.Printf("Searching for healthData")
    resp, err := healthDataTable.Run(Rethink)
    if err != nil {
        return nil, err
    }
    defer resp.Close()
    if resp.IsNil() {
        return nil, NotFoundError{}
    }
    var healthData []HealthData
    err = resp.All(&healthData)
    if err != nil {
        return nil, err
    }
    Info.Printf("Found %d healthData", len(healthData))
    return healthData, nil
}

func removeHealthData(id string) error {
    Info.Println("Deleting healthData "+ id)
    resp, err := healthDataTable.Get(id).Delete().RunWrite(Rethink)
    if err != nil {
        return err
    }
    if resp.Deleted == 0 {
        return NotFoundError{}
    }
    Info.Println("Removed healthData "+ id)
    return nil
}
