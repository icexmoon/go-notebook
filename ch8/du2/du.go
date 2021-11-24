package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sync"
	"time"
)

var paramHuman = flag.Bool("h", false, "human show")

func main() {
	start := time.Now()
	flag.Parse()
	dirNames := flag.Args()
	if len(dirNames) == 0 {
		//如果没有指定目录，使用当前目录
		dirNames = []string{"."}
	}
	var allSize int64
	var allFileNums int
	var fileScanWG sync.WaitGroup
	sizeChan := make(chan int64)
	gLimit := make(chan struct{}, 20)
	for _, name := range dirNames {
		fileScanWG.Add(1)
		go getDirSize(name, sizeChan, &fileScanWG, gLimit)
	}
	go func() {
		fileScanWG.Wait()
		close(sizeChan)
	}()
	for size := range sizeChan {
		allSize += size
		allFileNums += 1
	}
	printResult(allSize, allFileNums, *paramHuman)
	end := time.Now()
	usedTimes := end.Sub(start)
	fmt.Printf("used %.2f s.\n", usedTimes.Seconds())
}

func printResult(size int64, fileNums int, human bool) {
	if human {
		humanSize, humanUnit := humanByteSize(size)
		fmt.Printf("total %d %s,%d files.\n", humanSize, humanUnit, fileNums)
		return
	}
	fmt.Printf("total %d byte, %d files.\n", size, fileNums)
}

func humanByteSize(byteSize int64) (humanSize int64, humanUnit string) {
	const kb = 1024
	const mb = 1024 * kb
	const gb = 1024 * mb
	if byteSize > gb {
		humanSize = byteSize / gb
		humanUnit = "G"
	} else if byteSize > mb {
		humanSize = byteSize / mb
		humanUnit = "M"
	} else if byteSize > kb {
		humanSize = byteSize / kb
		humanUnit = "K"
	} else {
		humanSize = byteSize
		humanUnit = "B"
	}
	return
}

func getDirSize(dir string, sizeChan chan int64, fileScanWG *sync.WaitGroup, gLimit chan struct{}) {
	defer fileScanWG.Done()
	gLimit <- struct{}{}
	fis, err := ioutil.ReadDir(dir)
	<-gLimit
	if err != nil {
		log.Println(err)
		return
	}
	for _, fi := range fis {
		if fi.IsDir() {
			fileScanWG.Add(1)
			go getDirSize(path.Join(dir, fi.Name()), sizeChan, fileScanWG, gLimit)
		} else {
			sizeChan <- fi.Size()
		}
	}
}
