package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	var fsCloseChan = make(chan struct{}) //file scan close chan
	sizeChan := make(chan int64)
	gLimit := make(chan struct{}, 20)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(fsCloseChan)
	}()
	for _, name := range dirNames {
		fileScanWG.Add(1)
		go getDirSize(name, sizeChan, &fileScanWG, gLimit, fsCloseChan)
	}
	go func() {
		fileScanWG.Wait()
		close(sizeChan)
	}()
loop:
	for {
		select {
		case size, ok := <-sizeChan:
			if !ok {
				break loop
			}
			allSize += size
			allFileNums += 1
		case <-fsCloseChan:
			//消耗正在等待返回结果的goroutine以正常终止
			for range sizeChan {
			}
			fmt.Println("program is closed.")
			return
		}
	}
	printResult(allSize, allFileNums, *paramHuman)
	end := time.Now()
	usedTimes := end.Sub(start)
	fmt.Printf("used %.2f s.\n", usedTimes.Seconds())
}

func isFileScanClosed(fsCloseChan chan struct{}) bool {
	select {
	case <-fsCloseChan:
		//通道关闭，表示应当关闭所有并发任务
		return true
	default:
		//通道没有关闭
		return false
	}
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

func getDirSize(dir string, sizeChan chan int64, fileScanWG *sync.WaitGroup, gLimit chan struct{}, fsCloseChan chan struct{}) {
	defer fileScanWG.Done()
	if isFileScanClosed(fsCloseChan) {
		return
	}
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
			go getDirSize(path.Join(dir, fi.Name()), sizeChan, fileScanWG, gLimit, fsCloseChan)
		} else {
			sizeChan <- fi.Size()
		}
	}
}
