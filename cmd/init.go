package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var InitCommand = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new project",
	Action: func(c *cli.Context) error {
		envs := []string{"dev", "stg", "prod"}
		for _, env := range envs {
			envPath := filepath.Join("params", env)
			err := os.MkdirAll(envPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
		fmt.Println("Project initialized successfully")
		return nil
	},
}
