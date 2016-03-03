package main

import "time"

type HealthData struct {
    Id           string    `json:"id" gorethink:"id,omitempty"`

    ClusterLeft  int       `json:"clusterLinks"`
    ClusterRight int       `json:"clusterRechts"`
    Migraine     int       `json:"migraene"`
    Mood         int       `json:"stimmung"`
    Tension      int       `json:"anspannung"`

    Timestamp    time.Time `json:"timestamp"`
}
