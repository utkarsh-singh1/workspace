package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {


	// Directory creation
	err := os.Mkdir("MyDir", 0744)

	check(err)
	
	// Delete directory at the end (Similiar to rm -rf)
	// defer os.RemoveAll("MyDir")

	// Create a helper function to create a file
	createmptyFile := func(name string) {
		d := []byte("Haha")
		check(os.WriteFile(name, d, 0744))
	}

	createmptyFile("MyDir/file1")

	// We can create a hierarchy of directories, including parents with MkdirAll. This is Similiar to mkdir -p
	err = os.MkdirAll("MyDir/parent/child", 0744)
	check(err)

	createmptyFile("MyDir/parent/file1")
	createmptyFile("MyDir/parent/file2")
	createmptyFile("MyDir/parent/child/file4")


	// ReadDir can lists content of dir, returning a slice of os.DirEntry objects

	c, err := os.ReadDir("MyDir/parent")
	check(err)


	for _, value := range c {
		fmt.Println("",value.Name(),value.IsDir())
	}

	// Chdir let us change the directory

	err = os.Chdir("MyDir/parent/child")
	check(err)

	c,err = os.ReadDir(".")

	for _, value := range c {
		fmt.Println("",value.Name(),value.IsDir())
	}

	// Back to where we started
	err = os.Chdir("../../../")
	check(err)

	// we can also recursively visit directory, including its subdirectory using WalkDir, it accepts a callback function to handle every file or directory visited.

	err = filepath.WalkDir("MyDir", visit)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func visit(name string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	fmt.Println("",name,d.IsDir())
	return nil
}
