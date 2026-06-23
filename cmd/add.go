/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/RaucousDave/personal/devstash/storage"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [library name]",
	Short: "Add a new library or update an existing one",
	Long: `
Add a new library to devstash with its installation command.
If the library already exists, use flags to add commands or update the description.`,
	Example: `
  devstash add drizzle                  # add a new library
  devstash add drizzle --command        # add a command to existing library
  devstash add drizzle --description    # update description`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		addCommand, _ := cmd.Flags().GetBool("command")
		addDesc, _ := cmd.Flags().GetBool("desc")

		data, err := storage.ReadData()
		if err != nil {
			fmt.Println("Error while reading data ", err)
			return
		}

		if addCommand {

			var label string
			var command string

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Command label e.g. 'generate migration'").Value(&label),
				),
				huh.NewGroup(
					huh.NewInput().Title("Command e.g 'npx drizzle migrate'").Value(&command),
				),
			)

			err := form.Run()
			if err != nil {

				if errors.Is(err, huh.ErrUserAborted) {
					fmt.Println("Cancelled. ")
					return
				}
				fmt.Fprintln(os.Stderr, "Error: ", err)
				panic(err)
			}

			existingLibraries := data.Libraries[name]
			existingLibraries.Cmd = append(existingLibraries.Cmd, storage.Command{
				Label: label,
				Cmd:   command,
			})
			data.Libraries[name] = existingLibraries
			err = storage.WriteData(data)
			fmt.Printf("Added command to %s ✓\n", name)

		} else if addDesc {
			var desc string

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Command description e.g 'migrate db'").Value(&desc),
				),
			)

			err := form.Run()
			if err != nil {
				panic(err)
			}
			existingLibrary := data.Libraries[name]
			existingLibrary.Description = desc
			data.Libraries[name] = existingLibrary
			err = storage.WriteData(data)
			if err != nil {
				fmt.Println("Error while saving data", err)
			}
			fmt.Printf("Updated description to %s ✓\n", name)

		} else {
			var installCmd string
			var desc string

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Initialization command").Value(&installCmd),
				),
				huh.NewGroup(
					huh.NewInput().
						Title("Description of tool").Value(&desc),
				),
			)

			err := form.Run()
			if err != nil {
				if errors.Is(err, huh.ErrUserAborted) {
					fmt.Println("Cancelled.")
					return
				}
				fmt.Fprintln(os.Stderr, "Error: ", err)
				panic(err)

			}

			data.Libraries[name] = storage.Library{
				Install:     installCmd,
				Description: desc,
			}

			existingLibrary := data.Libraries[name]

			data.Libraries[name] = existingLibrary
			err = storage.WriteData(data)
			if err != nil {
				fmt.Println("Error while saving data ", err)
			}
			fmt.Printf("Saved %s ✓\n", name)

		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Bool("command", false, "Add a command to an existing library")
	addCmd.Flags().Bool("description", false, "Give a description for an existing command")
}
