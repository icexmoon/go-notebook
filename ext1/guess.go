package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	rand.Seed(time.Now().Unix())
	targetNum := rand.Intn(100) + 1 //产生的随机数是0-99，这里+1
	var success bool
	const TOTAL_TIMES = 10 //总的尝试次数
	for i := 0; i < TOTAL_TIMES; i++ {
		var leftTimes int = TOTAL_TIMES - i
		fmt.Printf("Please enter a number(%d times left):", leftTimes)
		inputNum := inputInt()
		if inputNum < targetNum {
			fmt.Println("low")
		} else if inputNum > targetNum {
			fmt.Println("hight")
		} else {
			fmt.Println("Success! You win!")
			success = true
			break
		}
	}
	if !success {
		fmt.Printf("You failed, the target number is %d.Please try again", targetNum)
	}
}

func inputInt() int {
	strNum, err := reader.ReadString('\n') //从命令行读取数据，换行为止
	if err != nil {
		log.Fatal(err)
	}
	strNum = strings.TrimSpace(strNum) //去除结尾的换行符
	num, err := strconv.Atoi(strNum)   //将字符串转换为int
	if err != nil {
		log.Fatal(err)
	}
	return num
}
