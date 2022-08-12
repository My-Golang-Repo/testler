package main

import(
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)


func main() {
	var configFile *string = flag.String("c", "myConfig", "Settng the conf file")
	flag.Parse()
	_,err:=os.Stat(*configFile)
	if err==nil {
		fmt.Println("Using user spesific file")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("./tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf(`%v\n`, err)
		return
	}
	fmt.Printf("Istifade olunan congfig %s \n", viper.ConfigFileUsed())

	if viper.IsSet("item1.k1"){
		fmt.Println("item bitrin deyeri: ",viper.Get("item1.k2"))
	}



}