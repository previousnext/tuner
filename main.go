package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/previousnext/tuner/conf"
	_ "github.com/previousnext/tuner/conf/apache"
	_ "github.com/previousnext/tuner/conf/passenger"
	_ "github.com/previousnext/tuner/conf/php"
)

var (
	cliConf    = kingpin.Flag("conf", "The type of configuration file to return").Default("apache").Envar("TUNER_CONF").String()
	cliMemory  = kingpin.Flag("memory", "Total available memory.").Default("512").Envar("TUNER_MEMORY").Int()
	cliAvgProc = kingpin.Flag("avg-proc", "The average memory size of a process.").Default("64").Envar("TUNER_AVG_PROC").Int()
	cliMaxProc = kingpin.Flag("max-proc", "The maximum allowed memory size of a process.").Default("128").Envar("TUNER_MAX_PROC").Int()
)

func main() {
	kingpin.Parse()


	// Get the configuration object.
	c, err := conf.Generate(*cliConf, *cliMemory, *cliAvgProc, *cliMaxProc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}
