package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/viniciussouzao/go-password-gen-cli/pkg/beauty"
	"github.com/viniciussouzao/go-password-gen-cli/pkg/generate"
	"github.com/viniciussouzao/go-password-gen-cli/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "password-gen",
	Short: "a simple password generator cli",
	Long:  beauty.Help(),
	Run:   exec,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Help message")

	// Length flag
	rootCmd.Flags().IntP("length", "l", 12, "The length of the password")

	// Custom password flags
	rootCmd.Flags().IntP("uppercase", "U", 0, "The number of uppercase letters")
	rootCmd.Flags().IntP("numbers", "N", 0, "The number of numbers")
	rootCmd.Flags().IntP("symbols", "S", 0, "The number of symbols")
	rootCmd.Flags().IntP("count", "c", 1, "The number of passwords to generate")

	// Feat flags
	rootCmd.Flags().Bool("clip", false, "Copy the password to the clipboard")
	rootCmd.Flags().StringP("file", "f", "", "Save the password to a file")

	// Version flag
	rootCmd.Flags().BoolP("version", "v", false, "Show the version of the cli")

}

func exec(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")

	uppercase, _ := cmd.Flags().GetInt("uppercase")
	numbers, _ := cmd.Flags().GetInt("numbers")
	symbols, _ := cmd.Flags().GetInt("symbols")
	count, _ := cmd.Flags().GetInt("count")

	clip, _ := cmd.Flags().GetBool("clip")
	file, _ := cmd.Flags().GetString("file")

	version, _ := cmd.Flags().GetBool("version")

	if version {
		color.White("Password Generator CLI v0.0.1")
		return
	}

	if length < 1 {
		color.Red("The length of the password must be greater than 0")
		return
	}

	if numbers+uppercase+symbols > length {
		color.Red("The sum of the number of characters must be equal to the length of the password")
		return
	}

	if numbers+uppercase+symbols == 0 {
		if count > 1 {
			var generatedPass string
			for i := 0; i < count; i++ {

				pass, err := generate.Password(length, false, false, false)
				if err != nil {
					color.Red("Error generating password: %v", err)
					return
				}

				generatedPass += pass + "\n"
			}

			if clip {
				err := utils.CopyToClipboard(generatedPass)
				if err != nil {
					color.Yellow("Error copying to clipboard: %v", err)
				}

				return
			}

			if file != "" {
				err := utils.SaveToFile(generatedPass, file)
				if err != nil {
					color.Yellow("Error saving to file: %v", err)
				}
				return
			}

			color.White(generatedPass)
			return

		} else {
			pass, err := generate.Password(length, false, false, false)
			if err != nil {
				color.Red("Error generating password: %v", err)
				return
			}

			if clip {
				err := utils.CopyToClipboard(pass)
				if err != nil {
					color.Yellow("Error copying to clipboard: %v", err)
				}

				return
			}

			if file != "" {
				err := utils.SaveToFile(pass, file)
				if err != nil {
					color.Yellow("Error saving to file: %v", err)
				}
				return
			}

			color.White(pass)
			return
		}
	}

	if numbers != 0 || uppercase != 0 || symbols != 0 {
		if count > 1 {
			var generatedPass string
			for i := 0; i < count; i++ {
				pass, err := generate.CustomPassword(length, numbers, uppercase, symbols)
				if err != nil {
					color.Red("Error generating password: %v", err)
					return
				}

				generatedPass += pass + "\n"
			}

			if clip {
				err := utils.CopyToClipboard(generatedPass)
				if err != nil {
					color.Yellow("Error copying to clipboard: %v", err)
				}

				return
			}

			if file != "" {
				err := utils.SaveToFile(generatedPass, file)
				if err != nil {
					color.Yellow("Error saving to file: %v", err)
				}
				return
			}

			color.White(generatedPass)

			return
		} else {
			pass, err := generate.CustomPassword(length, numbers, uppercase, symbols)
			if err != nil {
				color.Red("Error generating password: %v", err)
				return
			}

			color.White(pass)
			return
		}
	}

}
