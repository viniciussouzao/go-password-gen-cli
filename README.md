# GO Password Generator CLI
A simple command-line interface (CLI) tool for generating passwords written in Go.

## Installation
To install the Password Generator CLI, you can download the latest binary from the releases page. The binary is available for Windows, macOS, and Linux and the current version is 0.0.1.

Alternatively, you can build the binary from source.
```
# First, clone the repository:
$ git clone git@github.com:viniciussouzao/go-password-gen-cli.git

# Then, build the binary:
$ cd go-password-generator-cli

$ go build -o password-gen main.go

```

After the download or build is complete, you can move the binary to a directory in your PATH to make it easier to use. For example, on macOS and linux, you can move the binary to `/usr/local/bin`:
```

$ mv password-gen /usr/local/bin

```

## Usage
To generate a password, run the following command:

By default, this will generate a password with a length of 12 characters, containing only lowercase letters.

You can customize the password by using the following flags:
```
--length: The length of the password (default: 12)
--uppercase: The number of uppercase letters in the password
--numbers: The number of numbers in the password
--symbols: The number of symbols in the password
--count: The number of passwords to generate (default: 1)
--clip: Copy the password to the clipboard
--file: Save the password to a file
```

## Examples
To generate a password with a length of 16 characters, containing 4 uppercase letters, 2 numbers, and 2 symbols, you can run the following command:
```
$ password-gen --length 16 --uppercase 4 --numbers 2 --symbols 2
```

To generate 5 passwords with a length of 8 characters, containing 2 uppercase letters, 1 numbers, and 3 symbols, you can run the following command:
```
$ password-gen --length 8 --uppercase 2 --numbers 1 --symbols 3 --count 5
```

To generate a password with a length of 20 characters, containing 5 uppercase letters, 5 numbers, and 5 symbols, and copy it to the clipboard, you can run the following command:
```
$ password-gen --length 20 --uppercase 5 --numbers 5 --symbols 5 --clip
```

To generate a password with a length of 32 characters, containing 8 uppercase letters, 8 numbers, and 8 symbols, and save it to a file, you can run the following command:
```
$ password-gen --length 32 --uppercase 8 --numbers 8 --symbols 8 --file /tmp/password.txt
```


## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub. If you want to contribute code, please fork the repository and submit a pull request.

---
## License
This project is licensed under the MIT License - see the LICENSE file for details.
