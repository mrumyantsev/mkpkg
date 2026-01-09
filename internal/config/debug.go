package config

import (
	"fmt"
	"strings"
)

func (c *Config) showDebug() {
	fmt.Printf("Path: '%s'\n", c.Path)
	fmt.Printf("Filename: '%s'\n", c.Filename)

	eol := c.Eol
	eol = strings.ReplaceAll(eol, "\r", "\\r")
	eol = strings.ReplaceAll(eol, "\n", "\\n")

	fmt.Printf("Eol: '%s'\n", eol)

	fmt.Printf("PackageName: '%s'\n", c.PackageName)
	fmt.Printf("ObjectName: '%s'\n", c.ObjectName)
	fmt.Printf("ObjectType: '%s'\n", c.ObjectType)
	fmt.Printf("ObjectConstructor: '%s'\n", c.ObjectConstructor)
	fmt.Printf("ObjectReceiver: '%s'\n", c.ObjectReceiver)
	fmt.Printf("ObjectMethods: [")

	if len(c.ObjectMethods) != 0 {
		fmt.Printf("\n")

		for i := range c.ObjectMethods {
			fmt.Printf("  '%s',\n", c.ObjectMethods[i])
		}
	}

	fmt.Printf("]\n")
}
