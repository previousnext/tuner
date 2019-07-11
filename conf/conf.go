package conf

import (
	"errors"
)

// Conf interface
type Conf interface {
	TotalMemory(int)
	AvgProc(int)
	MaxProc(int)
	Build() (string, error)
}

// Confs is the map of config objects.
var Confs map[string]Conf

func init() {
	Confs = make(map[string]Conf)
}

// Register builds a list of all the templates.
func Register(name string, balancer Conf) error {
	if _, exists := Confs[name]; exists {
		return errors.New("conf is already defined")
	}
	Confs[name] = balancer

	return nil
}

// Generate initializes a new configuration, sets the config and generates a file.
func Generate(name string, memory, avgProc int, maxProc int) (string, error) {
	if t, exists := Confs[name]; exists {
		t.TotalMemory(memory)
		t.AvgProc(avgProc)
		t.MaxProc(maxProc)
		c, err := t.Build()
		if err != nil {
			return "", err
		}

		return c, nil
	}

	return "", errors.New("could not find the template")
}
