## OSX Keychain Tool

A tool that allows you to create and delete secrets inside the OSX keychain.
Secrets can currently be either

   * Plain text
   * Google OTP
   * Steam OTP

## To install

Check out the repo and run

```
make install
```

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

### To delete a secret

```
password --del <name>
```

### To get help
Either run the command with no options or use

```
$ password --help
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

## To Build without installing

Check out the Repo within your GOPATH and run

```
make
```

### Pre-requisites

   * This tool will only work on OSX
   * This tool requires Golang to be installed

## FAQ

   * I get asked to unlock my keychain
      * This is required to gain access to the secrets and is part of securing the data
   * Is this secure?
      * Its as secure as the OSX keychain. Locking your machine will render the keychain inaccessible (assuming Apple have no bugs).
   * Should I use this to store my password AND two factor for the same use
      * Its up to you but its worth remembering that the primary point of two factor is to require two disparate sources in order to gain access. This app DOES negate that but does draw upon OSX encryption to reduce the risk slightly.
   * I have a better idea!
      * Create an issue and i'll take a look.

## Help

Any help would be appreciated. Please create PR's.

Items to tackle are in the issues

## Thanks

To all the examples online which made this code possible. I'll credit each snippet when I find them again.

