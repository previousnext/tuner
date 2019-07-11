package conf

import (
	"errors"
)

// Conf interface
type Conf interface {
	Max(int)
	Proc(int)
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
func Generate(name string, max, proc int) (string, error) {
	if t, exists := Confs[name]; exists {
		t.Max(max)
		t.Proc(proc)
		c, err := t.Build()
		if err != nil {
			return "", err
		}

		return c, nil
	}

	return "", errors.New("could not find the tpl")
}
