//go:generate go run generate.go

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"

    //"sort"
    "gopkg.in/yaml.v2"
)

type Config struct {
	EntryPoints map[string]string `yaml:"service_entrypoints"`
}


func main() {
    // YAMLファイルの読み込み
    data, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Fatalf("Failed to read YAML file: %v", err)
    }

    // YAMLデータを構造体に変換
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Failed to unmarshal YAML: %v", err)
    }
    

    // const型の変数として出力
    output := fmt.Sprintf("package main\n\nconst (\n")
    for key,value := range config.EntryPoints {
        output += fmt.Sprintf("  %s = \"%s\"\n", key, value)
    }
    output += fmt.Sprintf(")\n")
    err = ioutil.WriteFile("constants.go", []byte(output), os.ModePerm)
    if err != nil {
        log.Fatalf("Failed to write constants file: %v", err)
    }
}

