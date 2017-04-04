package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "log"
    "text/template"
    "flag"
    "fmt"
)

//srtuct define the  values in values.yaml
//when adding more values in the yaml file add in the struct too
type Config struct {
    NameSpace string `yaml:"NameSpace"`
    AppName string `yaml:"AppName"`
}

//read the yaml values stored in passed filename as arg 1
//Output is retuned in struct Config
func getValuesFromYaml(yamlFilename string) Config {
    log.Println("Reading Yaml file")
    var config Config
    values, err := ioutil.ReadFile(yamlFilename)
    if err != nil {
        log.Fatal(err)
    }
    err = yaml.Unmarshal(values, &config)
    if err != nil {
        log.Fatal(err)
    }
    return config

}

func writeToYaml (tmplFilename string, values Config) {
  log.Println("Generating yaml")
  tmpl, err := template.ParseFiles(tmplFilename)
  if err != nil {
      log.Fatal(err)
  }
  log.Println("=== yaml output ===:")
  err  = tmpl.ExecuteTemplate(os.Stdout, tmplFilename, values)
  if err != nil {
      log.Fatal(err)
  }
}

func main() {
    valuesYamlfile := flag.String("values", "values.yaml", "values yaml file name")
    tmplFile := flag.String("tmplfile", "svc.yaml", "Template yaml file name")
    flag.Parse()
    fmt.Println(*valuesYamlfile)
    fmt.Println(*tmplFile)
    values := getValuesFromYaml(*valuesYamlfile)
    writeToYaml(*tmplFile, values)

}
