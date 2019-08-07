## OSX Keychain Tool

A tool that allows you to create and delete secrets inside the OSX keychain.
Secrets can currently be either

   * Plain text
   * Google OTP
   * Steam OTP

## Usage

### To retrieve a secret

```
password <name>
```

By default secrets are copied into the copy and paste buffer ready to be pasted.

The ```-p``` option will print the secret to the screen.

There is currently no way to view or dump all of your stored secrets other than by looking at the OSX Keychain.

### To add a plain text secret 

```
password --add <name> <secret>
```

### To add a google OTP

```
password --add <name> -otp <recovery code>
```

### To add a steam OTP

```
password --add <name> -steam <recovery code>
```

### To get help
Either run the command with no options or use

```
password --help
OSX Password Keychain tool

Usage: password [--otp] [--steam] [--print] [--add ADD] [--del DEL] SECRET

Positional arguments:
  SECRET

Options:
  --otp, -o              Add Recovery code for generic OTP key
  --steam, -s            Add Recovery code for Steam OTP key
  --print, -p            Print the value to screen rather than copy to the clipboard
  --add ADD, -a ADD      Name of the secret to add
  --del DEL, -d DEL      Name of the secret to delete
  --help, -h             display this help and exit
```

## To Build

Check out the Repo within your GOPATH and run

```
make
```

### Pre-requisites

   * This tool will only work on OSX
   * This tool requires Golang to be installed

## Thanks

To all the examples online which made this code possible. I'll credit each snippet when ic can find them again.

