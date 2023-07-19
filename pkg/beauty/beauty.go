package beauty

import "github.com/pterm/pterm"

func Help() string {
	helpText := ""

	helpText += pterm.DefaultSection.Sprint("Password Generator CLI") + "A simple password generator cli that generates a password based on the length and the type of characters you want to use.\n"
	helpText += pterm.DefaultSection.Sprint("Usage")
	helpText += "* For example, you can generate a simple password in lowercase of length 12 with the following command:\n\n"
	helpText += " $ " + pterm.LightGreen("password-gen --length 12") + "\n\n"
	helpText += "* Generate a password of length 10 with 5 lowercase letters, 2 uppercase letters, 2 numbers and 1 symbol:\n\n"
	helpText += " $ " + pterm.LightGreen("password-gen --length 10 --uppercase 2 --numbers 2 --symbols 1") + "\n\n"
	helpText += "* Specify the number of passwords you want to generate:\n\n"
	helpText += "$ " + pterm.LightGreen("password-gen --length 12 --count 5") + "\n\n"
	helpText += "* Copy the password to the clipboard:\n\n"
	helpText += " $ " + pterm.LightGreen("password-gen --length 12 --clip") + "\n"
	helpText += pterm.LightYellow(" NOTE: On linux, you need to install xclip to use the --clip flag.") + "\n\n"
	helpText += "* Save the password to a file:\n\n"
	helpText += " $ " + pterm.LightGreen("password-gen --length 12 --file /tmp/password.txt") + "\n"

	return helpText
}

func Version() string {

	versionText := "Password Generator CLI v0.0.1"

	return versionText

}
