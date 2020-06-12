package main

import (
	cryptorand "crypto/rand"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/uuid"
)

func main() {
	currentDir, _ := os.Getwd()
	filecount := flag.Int("Count", 0, "[Required] Number of files to generate")
	targetdir := flag.String("Dir", currentDir, "Target directory for files. Will be created if not exists.")
	filesize := flag.String("Size", "32", "Filesizes in kilobyte. Single value or comma seperated for random distribution.")
	flag.Parse()

	if flag.NFlag() == 0 || *filecount == 0 {
		fmt.Println("\nYou must specify at least the number of files with the -Count flag.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filesizes := strings.Split(*filesize, ",")
	fmt.Printf("\nGenerating %d files of size %s in directory %s.\n", *filecount, *filesize, *targetdir)

	createDirIfNotExists(*targetdir)
	generateFiles(*filecount, *targetdir, filesizes)
}

func generateFiles(count int, targetdir string, filesizes []string) {
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		filename := fmt.Sprintf("%s\\%s.bin", targetdir, uuid.New())
		randomfilesize, err := strconv.Atoi(filesizes[rand.Intn(len(filesizes))])
		if err != nil {
			fmt.Printf("Cannot parse filesizes parameter: %q", err)
			os.Exit(1)
		}
		createRandomFile(filename, randomfilesize*1024)
		bar.Increment()
	}
	bar.Finish()
}

func createRandomFile(name string, length int) {
	data := make([]byte, length)
	_, err := cryptorand.Read(data)

	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cannot create file: %q", err)
		os.Exit(1)
	}
	file.Write(data)
}

func createDirIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Unable to create directory for files: %q", err)
			os.Exit(1)
		}
	}
}
