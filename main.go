package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/howeyc/gopass"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

const saltLength int = 8

var (
	generate      = flag.Bool("generate", false, "generate a random password (optional)")
	passLength    = flag.Int("length", 14, "(default 14) password length of generated password")
	saltChars     = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	passwordChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&()-_=+,.?;{}[]`~")
)

func newSaltSHA512(length int) (string, error) {
	randomChars, err := randChar(length, saltChars)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("$6$%s", randomChars), nil
}

func newPassword(length int) (string, error) {
	p, err := randChar(length, passwordChars)
	if err != nil {
		return "", err
	}

	return p, nil
}

func randChar(length int, chars []byte) (string, error) {
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			return "", err
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPword), nil
			}
		}
	}
}

func main() {
	flag.Parse()

	// Generate a random salt for SHA512 password
	saltSha512, err := newSaltSHA512(saltLength)
	if err != nil {
		log.Fatal(err)
	}

	var passwordSha512 string
	if *generate == true {
		// Generate a random password for the SHA512 algorithm
		passwordSha512, err = newPassword(*passLength)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Get password from user
		fmt.Printf("Password: ")
		passwordFromUser, err := gopass.GetPasswdMasked()
		if err != nil {
			log.Fatal(err)
		}
		passwordSha512 = string(passwordFromUser)
	}

	// Generate the SHA512 password hash
	d := sha512_crypt.New()
	hashSha512, err := d.Generate([]byte(passwordSha512), []byte(saltSha512))
	if err != nil {
		log.Fatal(err)
	}

	// Print the password, if generated
	if *generate == true {
		fmt.Println(passwordSha512)
	}

	// Print the SHA512 hash of the password
	fmt.Println(hashSha512)
}
