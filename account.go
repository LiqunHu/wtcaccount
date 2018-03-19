package main

import (
	"fmt"
	"os"

	"github.com/LiqunHu/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const hardened = 0x80000000

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	if len(os.Args) > 1 {
		mnemonic = os.Args[1]
	}

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	// fmt.Println("seed: ", seed)
	masterKey, _ := bip32.NewMasterKey(seed)
	account, _ := masterKey.DerivePath("m/44'/60'/0'/0/0")

	// Display mnemonic and keys
	if len(os.Args) == 1 {
		fmt.Println(mnemonic)
	}
	fmt.Printf("%x\n", account.Key)
}
