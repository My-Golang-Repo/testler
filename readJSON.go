/* package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./myJSONconfig.json")
	fmt.Printf("istifade olunan config: %s \n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("item1.key1"){
		fmt.Println(viper.Get("item1.key1"))
	} else {
		fmt.Println("ITEM YOXDUR")
	}

	if viper.IsSet("item3.key1"){
		fmt.Println(viper.Get("item1.key1"))
	} else {
		fmt.Println("ITEM YOXDUR")
	}


} */