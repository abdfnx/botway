package token

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os/user"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/charmbracelet/lipgloss"
)

var (
	FocusedStyle  = lipgloss.NewStyle().Foreground(constants.PRIMARY_COLOR)
	BlurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	BoldStyle     = lipgloss.NewStyle().Bold(true)
	CursorStyle   = FocusedStyle.Copy()
	NoStyle       = lipgloss.NewStyle()
	FocusedButton = FocusedStyle.Copy().Render("[ Done ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Done"))
)

func Generator(secret string) string {
	rand.Seed(time.Now().Unix())
	charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP1234567890" + secret)

	var output strings.Builder

	for i := 0; i < 16; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}

	return output.String()
}

func EncryptTokens() (string, string) {
	username, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	var encryptAES = func(data string) string {
		text := []byte(data)
		key := []byte(Generator(username.Username))

		c, err := aes.NewCipher(key)

		if err != nil {
			fmt.Println(err)
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}

		nonce := make([]byte, gcm.NonceSize())

		if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
			fmt.Println(err)
		}

		return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, text, nil))
	}

	return encryptAES("Access Token for " + username.Username), encryptAES("Refresh Token for " + username.Username)
}

func CreateRSATokens() (string, string) {
	bitSize := 4096

	// Generate RSA key.
	key, err := rsa.GenerateKey(crand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	return string(pubPEM), string(keyPEM)
}
