package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/appleboy/easyssh-proxy"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "appleboy",
		Server: "example.com",
		// Optional key or Password without either we try to contact your agent SOCKET
		// Password: "password",
		// Paste your source content of private key
		// Key: `-----BEGIN RSA PRIVATE KEY-----
		// MIIEpAIBAAKCAQEA4e2D/qPN08pzTac+a8ZmlP1ziJOXk45CynMPtva0rtK/RB26
		// 7XC9wlRna4b3Ln8ew3q1ZcBjXwD4ppbTlmwAfQIaZTGJUgQbdsO9YA==
		// -----END RSA PRIVATE KEY-----
		// `,
		KeyPath: "/Users/username/.ssh/id_rsa",
		Port:    "22",
		Timeout: 60 * time.Second,

		// Parse PrivateKey With Passphrase
		Passphrase: "1234",

		// Optional fingerprint SHA256 verification
		// Get Fingerprint: ssh.FingerprintSHA256(key)
		// Fingerprint: "SHA256:mVPwvezndPv/ARoIadVY98vAC0g+P/5633yTC4d/wXE"
	}

	// Example 1: Upload a file using SFTP
	fmt.Println("=== SFTP Upload Example ===")
	err := ssh.SftpUpload("/local/path/source.txt", "/remote/path/target.txt")
	if err != nil {
		log.Printf("SFTP Upload failed: %v", err)
	} else {
		fmt.Println("SFTP Upload successful!")
	}

	// Example 2: Download a file using SFTP
	fmt.Println("\n=== SFTP Download Example ===")
	err = ssh.SftpDownload("/remote/path/source.txt", "/local/path/downloaded.txt")
	if err != nil {
		log.Printf("SFTP Download failed: %v", err)
	} else {
		fmt.Println("SFTP Download successful!")
	}

	// Example 3: List directory contents
	fmt.Println("\n=== SFTP List Directory Example ===")
	fileInfos, err := ssh.SftpList("/remote/directory")
	if err != nil {
		log.Printf("SFTP List failed: %v", err)
	} else {
		fmt.Printf("Found %d items in directory:\n", len(fileInfos))
		for _, info := range fileInfos {
			fmt.Printf("  %s (size: %d, mode: %s)\n",
				info.Name(), info.Size(), info.Mode())
		}
	}

	// Example 4: Create a directory
	fmt.Println("\n=== SFTP Create Directory Example ===")
	err = ssh.SftpMkdir("/remote/new_directory")
	if err != nil {
		log.Printf("SFTP Mkdir failed: %v", err)
	} else {
		fmt.Println("SFTP Directory created successfully!")
	}

	// Example 5: Create directories recursively
	fmt.Println("\n=== SFTP Create Directory Recursively Example ===")
	err = ssh.SftpMkdirAll("/remote/path/to/nested/directory")
	if err != nil {
		log.Printf("SFTP MkdirAll failed: %v", err)
	} else {
		fmt.Println("SFTP Nested directories created successfully!")
	}

	// Example 6: Get file information
	fmt.Println("\n=== SFTP File Stat Example ===")
	fileInfo, err := ssh.SftpStat("/remote/path/file.txt")
	if err != nil {
		log.Printf("SFTP Stat failed: %v", err)
	} else {
		fmt.Printf("File info: %s (size: %d, mode: %s, modified: %s)\n",
			fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.ModTime())
	}

	// Example 7: Change file permissions
	fmt.Println("\n=== SFTP Change Permissions Example ===")
	err = ssh.SftpChmod("/remote/path/file.txt", 0644)
	if err != nil {
		log.Printf("SFTP Chmod failed: %v", err)
	} else {
		fmt.Println("SFTP File permissions changed successfully!")
	}

	// Example 8: Remove a file
	fmt.Println("\n=== SFTP Remove File Example ===")
	err = ssh.SftpRemove("/remote/path/file_to_delete.txt")
	if err != nil {
		log.Printf("SFTP Remove failed: %v", err)
	} else {
		fmt.Println("SFTP File removed successfully!")
	}

	// Example 9: Working with SFTP client directly for advanced operations
	fmt.Println("\n=== SFTP Direct Client Example ===")
	sftpClient, client, err := ssh.SftpClient()
	if err != nil {
		log.Printf("SFTP Client creation failed: %v", err)
		return
	}
	defer client.Close()
	defer sftpClient.Close()

	// You can now use sftpClient for any advanced SFTP operations
	// For example, opening a file for reading/writing
	file, err := sftpClient.OpenFile("/remote/path/file.txt", os.O_RDWR|os.O_CREATE)
	if err != nil {
		log.Printf("SFTP OpenFile failed: %v", err)
	} else {
		defer file.Close()
		fmt.Println("SFTP File opened successfully for advanced operations!")

		// Write some data
		_, err = file.Write([]byte("Hello from SFTP!\n"))
		if err != nil {
			log.Printf("SFTP Write failed: %v", err)
		} else {
			fmt.Println("Data written to file successfully!")
		}
	}

	fmt.Println("\n=== All SFTP examples completed ===")
}
