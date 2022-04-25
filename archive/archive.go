package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("creating zip archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create("csv/test.csv")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	fmt.Println("opening second file")
	f2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	fmt.Println("writing second file to archive...")
	w2, err := zipWriter.Create("txt/test.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w2, f2); err != nil {
		panic(err)
	}
	fmt.Println("closing zip archive...")
	zipWriter.Close()
}
