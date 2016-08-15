package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/previousnext/tuner/conf"
	_ "github.com/previousnext/tuner/conf/apache"
	_ "github.com/previousnext/tuner/conf/php"
	_ "github.com/previousnext/tuner/conf/passenger"
)

var (
	cliConf       = kingpin.Flag("conf", "The type of configuration file to return").Default("apache").OverrideDefaultFromEnvar("TUNER_CONF").String()
	cliMax        = kingpin.Flag("max", "Max memory allocated to this instance").Default("512").OverrideDefaultFromEnvar("TUNER_MAX").Int()
	cliProc       = kingpin.Flag("proc", "The size of the PHP proccess.").Default("128").OverrideDefaultFromEnvar("TUNER_PROC").Int()
	cliMultiplier = kingpin.Flag("multiplier", "The multiplier for calculating apache max clients").Default("2").OverrideDefaultFromEnvar("TUNER_MULTIPLIER").Int()
)

func main() {
	kingpin.Parse()

	// Apply the mupltier here, the configurations should not have a account for this.
	multiMax := *cliMax * *cliMultiplier

	// Get the configuration object.
	c, err := conf.Generate(*cliConf, multiMax, *cliProc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}