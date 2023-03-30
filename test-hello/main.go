package main

import (
	cr "crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
	"unsafe"
)

// 给Int扩展个方法ToString
type Int int

func (i Int) ToString() string {
	return strconv.Itoa(int(i))
}

// 取随机整数
func randomInt(min, max int) int {
	res, _ := cr.Int(cr.Reader, big.NewInt(int64(max-min)))
	return min + int(res.Int64())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bytes, err := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("whil reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, bytes, url)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("命令行参数是：", os.Args[1:])
	}

	var i Int = 32
	fmt.Printf("%T %s\n", i.ToString(), i.ToString())

	fmt.Println("计算占用内存 int size = ", unsafe.Sizeof(i), "string size = ", i.ToString())

	fmt.Println("随机整数：", randomInt(50, 100))

	fmt.Println("Hello World")

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
