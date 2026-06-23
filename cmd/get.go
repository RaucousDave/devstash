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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [tool name]",
	Short: "Get commands for a specific library",
	Long: `Retrieve and browse saved commands for a specific library.

Examples:
  devstash get drizzle        # browse drizzle commands
  devstash get better-auth    # browse better-auth commands
  devstash get gorilla/mux    # browse gorilla/mux commands`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		data, err := storage.ReadData()
		if err != nil {
			fmt.Println("Error reading data", err)
			return
		}
		lib, exists := data.Libraries[name]
		if !exists {
			fmt.Printf("%s not found \n", name)
			return
		}

		if len(lib.Cmd) == 0 {
			fmt.Printf("No command saved for %s\n", name)
			return
		}

		var options []huh.Option[string]
		for _, c := range lib.Cmd {
			options = append(options, huh.NewOption(fmt.Sprintf("%s - %s", c.Label, c.Cmd), c.Cmd))
		}

		var selected string

		err = huh.NewSelect[string]().Options(options...).Value(&selected).Run()
		if err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				fmt.Println("Cancelled. ")
				return
			}
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}

		var shellCmd *exec.Cmd

		if runtime.GOOS == "windows" {
			shellCmd = exec.Command("cmd", "/C", selected)
		} else {
			shellCmd = exec.Command("sh", "-c", selected)
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
	rootCmd.AddCommand(getCmd)
}
