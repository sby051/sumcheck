package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	fileName := flag.String("f", "", "file to check")
	sum := flag.String("s", "", "sum to check")
	algorithm := flag.String("a", "md5", "algorithm to use")
	help := flag.Bool("h", false, "show help")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *fileName == "" || *sum == "" || *algorithm == "" {
		fmt.Println("missing required flags")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(*fileName); os.IsNotExist(err) {
		fmt.Println("file does not exist")
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var hasher hash.Hash

	switch *algorithm {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	default:
		fmt.Println("invalid algorithm")
		os.Exit(1)
	}

	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("file checksum: %x\n", hasher.Sum(nil))
	fmt.Printf("sum to check: %s\n", *sum)

	if fmt.Sprintf("%x", hasher.Sum(nil)) != *sum {
		fmt.Println("checksums do NOT match, file may be corrupt or tampered with. proceed with caution!")
		os.Exit(1)
	}

	fmt.Println("checksums match, file is good to go.")
}
