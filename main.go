package main

import (
	"fmt"
	"log"
	"os"

	"go-torrent/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]
	fmt.Println(inPath, outPath)

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
