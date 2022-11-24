package token

import (
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/tran/dfs"
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

func Generator() string {
	rand.Seed(time.Now().Unix())
	charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP1234567890")

	var output strings.Builder

	for i := 0; i < 16; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}

	return output.String()
}

func CreateRSATokens() {
	privatekey, err := rsa.GenerateKey(crand.Reader, 1024)

	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}

	keysDir := dfs.CreateDirectory(filepath.Join(constants.BotwayDirPath, "keys"))

	if keysDir != nil {
		log.Fatal(keysDir)
	}

	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePem, err := os.Create(filepath.Join(constants.BotwayDirPath, "keys", "bwrsa.rsa"))

	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}

	err = pem.Encode(privatePem, privateKeyBlock)

	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)

	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPem, err := os.Create(filepath.Join(constants.BotwayDirPath, "keys", "bwrsa-pub.rsa"))

	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}

	err = pem.Encode(publicPem, publicKeyBlock)

	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
}
