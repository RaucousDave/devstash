/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/RaucousDave/personal/devstash/storage"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved libraries",
	Long: `Display all libraries currently saved in your stash.

	Examples:
  devstash list    # show everything in your stash`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadData()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while reading data ", err)
			return
		}

		if len(data.Libraries) == 0 {
			fmt.Fprintln(os.Stderr, "No libraries saved yet")
			return
		}

		var options []huh.Option[string]
		for name, lib := range data.Libraries {
			options = append(options, huh.NewOption(fmt.Sprintf("%s - %s", name, lib.Desc()), name))
		}

		var selected string

		err = huh.NewSelect[string]().Options(options...).Value(&selected).Run()
		if err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				fmt.Println("Cancelled. ")
				return
			}
			fmt.Fprintln(os.Stderr, "Error: ", err)
			return
		}
		lib := data.Libraries[selected]

		var libOptions []huh.Option[string]
		for _, c := range lib.Cmd {

			libOptions = append(libOptions, huh.NewOption(fmt.Sprintf("%s - %s", c.Label, c.Cmd), c.Cmd))
		}

		var selectedCmd string
		err = huh.NewSelect[string]().Options(libOptions...).Value(&selectedCmd).Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			if errors.Is(err, huh.ErrUserAborted) {
				fmt.Println("Cancelled.")
				return
			}
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}

		var shellCmd *exec.Cmd

		if runtime.GOOS == "windows" {
			shellCmd = exec.Command("cmd", "/C", selectedCmd)
		} else {
			shellCmd = exec.Command("sh", "-c", selectedCmd)
		}
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		shellCmd.Stdin = os.Stdin

		if err := shellCmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
