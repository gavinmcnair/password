## OSX Keychain Tool

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
password --add <name> -otp <secret>
```

### To add a steam OTP

```
password --add <name> -steam <secret>
```

## To Build

Check out the Repo within your GOPATH and run

```
make
```

### Pre-requisites

   * This tool will only work on OSX
   * This tool requires Golang to be installed 

