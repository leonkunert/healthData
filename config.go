package main

import (
    "github.com/olebedev/config"
    "flag"
)

type Configuration interface {
    ParseYamlFile(filename string) (*config.Config, error)
}

/*
Initialize Configuration
Defines some configuration options in a global variable
 */
var CFG *config.Config

func loadConfig() {
    var configPath string
    flag.StringVar(&configPath, "config", "", "The path to configuration")
    flag.Parse()
    if configPath == "" {
        Error.Panic("No config specified")
    }
    config, err := config.ParseYamlFile(configPath)
    if err != nil {
        Error.Println(err.Error())
        return
    }
    Info.Println("config loaded from " + configPath)
    CFG = config.Flag()
}
