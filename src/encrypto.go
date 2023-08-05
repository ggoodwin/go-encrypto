package goencrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func encryptFile(key []byte, inputFile, outputFile string) error {
	// Read the input file
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	// Create the output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	// Generate a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the output file first
	if _, err := output.Write(iv); err != nil {
		return err
	}

	// Create the AES cipher mode for encryption
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt and write data to the output file
	buffer := make([]byte, 4096)
	for {
		n, err := input.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		if _, err := output.Write(buffer[:n]); err != nil {
			return err
		}
	}

	return nil
}

func decryptFile(key []byte, inputFile, outputFile string) error {
	// Read the input file
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	// Create the output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	// Read the IV (Initialization Vector) from the input file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(input, iv); err != nil {
		return err
	}

	// Generate a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create the AES cipher mode for decryption
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt and write data to the output file
	buffer := make([]byte, 4096)
	for {
		n, err := input.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		if _, err := output.Write(buffer[:n]); err != nil {
			return err
		}
	}

	return nil
}