package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/lunixbochs/go-keychain"
	"log"
	"os"
	"strings"
)

const secretStore = "AuthStore"

type args struct {
	Otp    bool   `arg:"-o,separate" help:"Add Recovery code for generic OTP key"`
	Steam  bool   `arg:"-s,separate" help:"Add Recovery code for Steam OTP key"`
	Print  bool   `arg:"-p,separate" help:"Print the value to screen rather than copy to the clipboard"`
	Add    string `arg:"-a,separate" help:"Name of the secret to add"`
	Del    string `arg:"-d,separate" help:"Name of the secret to delete"`
	Secret string `arg:"positional"`
}

func main() {
	var args args
	parsedArgs := arg.MustParse(&args)

	// Args Handling. Needs improvement
	if args.Add != "" && args.Secret != "" {
		fmt.Printf("Adding secret name: %s with value %s\n", args.Add, args.Secret)
		if args.Steam {
			keychain.Add(secretStore, args.Add, "Steam±"+args.Secret)
		} else if args.Otp {
			keychain.Add(secretStore, args.Add, "TOTP±"+args.Secret)
		} else {
			keychain.Add(secretStore, args.Add, "Password±"+args.Secret)
		}
	} else if args.Del != "" {
		fmt.Printf("Deleting secret name: %s\n", args.Del)
		keychain.Remove(secretStore, args.Del)
	} else if args.Secret != "" {
		password := retrievePassword(args.Secret)
		if args.Print {
			fmt.Printf("Password: %s\n", password)
		} else {
			fmt.Printf("Password has been copied into buffer.\n")
			writeAll(password)
		}
	} else {
		parsedArgs.WriteHelp(os.Stdout)
	}
}

func retrievePassword(secret string) string {
	data, err := keychain.Find(secretStore, secret)
	if err != nil {
		fmt.Printf("Entry cannot be found for secretname: %s\n", secret)
		os.Exit(0)
	}
	// This method of storing metadata needs improvement.
	returneddata := strings.SplitN(data, "±", 2)
	if returneddata[0] == "Password" {
		return returneddata[1]
	} else {
		return decodeOTP(returneddata[1], returneddata[0])
	}
}

func decodeOTP(code string, otptype string) string {
	// Get OTP -1/+1 from currentTS for previous/next
	currentTS, _ := TimeStamp()
	secret := normalizeSecret(code)
	// Options are TOTP or Steam
	otp, e := AuthCode(secret, currentTS, otptype)
	if e != nil {
		log.Fatal(e)
	}
	return otp
}

func (args) Description() string {
	return "OSX Password Keychain tool\n"
}
