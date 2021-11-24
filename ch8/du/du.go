package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
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
	for _, name := range dirNames {
		size, fileNums := getDirSize(name)
		allSize += size
		allFileNums += fileNums
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

func getDirSize(dir string) (size int64, fileNums int) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}
	for _, fi := range fis {
		if fi.IsDir() {
			subDirSize, subDirFn := getDirSize(path.Join(dir, fi.Name()))
			size += subDirSize
			fileNums += subDirFn
		} else {
			size += fi.Size()
			fileNums += 1
		}
	}
	return
}
