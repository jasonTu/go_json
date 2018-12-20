package main

import "os"

func tempDirs() []string {
	folders := []string{"/tmp/tfolder1", "/tmp/tfolder2"}
	return folders
}

// Version 1
/*
func main() {
	var rmdirs []func()

	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}
*/

// Version 2
func main() {
	var rmdirs []func()

	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	// Will only remove one folder
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
