package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.Get("name")) // this would be "steve"
	fmt.Println(viper.Get("hobbies"))

}
