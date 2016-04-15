package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	memory     = kingpin.Flag("memory", "Max memory allocated to this instance").Default("1024").OverrideDefaultFromEnvar("TUNER_MEMORY").Int()
	apacheConf = kingpin.Flag("apache-conf", "Apache configuration file for max procs").Default("/etc/apache2/mods-enabled/mpm_prefork.conf").OverrideDefaultFromEnvar("TUNER_APACHE_CONF").String()
	phpConf    = kingpin.Flag("php-conf", "PHP configuration file").Default("/usr/local/etc/php/php.ini").OverrideDefaultFromEnvar("TUNER_PHP_CONF").String()
	phpMem     = kingpin.Flag("php-memory", "The size of the PHP proccess.").Default("128").OverrideDefaultFromEnvar("TUNER_PHP_MEMORY").Int()
	multiplier = kingpin.Flag("multiplier", "The multiplier for calculating apache max clients").Default("2").OverrideDefaultFromEnvar("TUNER_MULTIPLIER").Int()
)

func main() {
	kingpin.Parse()

	// Compute the total procs.
	procs, err := ApacheProcs(*memory, *phpMem, *multiplier)
	if err != nil {
		Exit(err)
	}

	// Tell the user so we can debug at a later date if required.
	fmt.Printf("Tuning system to: Memory = %v, PHP = %v, Procs = %v\n", *memory, *phpMem, procs)

	// Update Apache configuration to use the required procs.
	err = Write("apache", apacheTpl, procs, *apacheConf)
	if err != nil {
		Exit(err)
	}

	// Update PHP configuration to use the required memory per proc.
	err = Write("php", phpTpl, *phpMem, *phpConf)
	if err != nil {
		Exit(err)
	}
}

func ApacheProcs(memory, phpMem, multiplier int) (int, error) {
	return memory / phpMem * multiplier , nil
}

func Write(name, tpl string, val int, file string) error {
	fm := template.FuncMap{"divide": func(a, b int) int {
		return a / b
	}}

	t := template.Must(template.New(tpl).Funcs(fm).Parse(tpl))

	// Write the contents to memory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, val)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, buf.Bytes(), 0644)
}

func Exit(e error) {
	fmt.Println(e)
	os.Exit(1)
}
