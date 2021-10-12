package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "/path/to/usbcopy-dir/"
	output := "/path/to/output-dir/"
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if root == path {
				return nil
			}
			start := len(root)
			end := len(path)

			subfolder := path[start:end]

			dist := subfolder[0:7]

			full_dist := strings.Join([]string{output, dist}, "")
			if info.IsDir() {
				if _, err := os.Stat(full_dist); os.IsNotExist(err) {
					_ = os.Mkdir(full_dist, 0777)
				}
			} else {
				fmt.Println(path, strings.Join([]string{full_dist, info.Name()}, "/"))
				err = os.Rename(path, strings.Join([]string{full_dist, info.Name()}, "/"))
				if err != nil {
					log.Fatal(err)
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
