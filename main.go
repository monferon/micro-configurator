package main

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	v1 "github.com/monferon/golang-test/controller/http/v1"
	"github.com/monferon/golang-test/pkg/httpserver"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func Run() {
	var it interface{}
	server := initial("config.json", it)
	Watcher(server, it)
}
func Watcher(server *httpserver.Server, it interface{}) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					server.Shutdown()
					server = initial(event.Name, it)
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("config.json")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 64, 64)
	return err == nil
}

func sum(res, b int) int {
	return res + b
}

//func summ(a string, b string) string {
//	aAr = strings.Split(a, "")
//	bAr = strings.Split(a, "")
//	return ""
//}

//func Summary(a []string, b []string) []int {
//	res := make([]int, 0, len(a))
//	temporary := 0
//	if len(a) < len(b) {
//		for i := 0; i < len(a); i++ {
//			number := 0
//			if i > len(b)-1 {
//				number = a[i] + temporary
//			} else if i > len(a)-1 {
//				number = b[i] + temporary
//			} else {
//				number = a[i] + b[i] + temporary
//			}
//
//			temporary = 0
//			if number > 9 {
//				number = number - 10
//				temporary = 1
//			}
//			res = append(res, number)
//		}
//		if temporary != 0 {
//			res = append(res, temporary)
//		}
//	}
//
//	if len(a) > len(b) {
//		for i := 0; i < len(b); i++ {
//			number := 0
//			if i > len(b)-1 {
//				number = a[i] + temporary
//			} else if i > len(a)-1 {
//				number = b[i] + temporary
//			} else {
//				number = a[i] + b[i] + temporary
//			}
//
//			temporary = 0
//			if number > 9 {
//				number = number - 10
//				temporary = 1
//			}
//			res = append(res, number)
//		}
//		if temporary != 0 {
//			res = append(res, temporary)
//		}
//	}
//
//	return res
//
//}

func main() {
	//Read main config
	//Settings for server.
	Run()
	//import (
	//	"bufio"
	//"fmt"
	//"os"
	//)

	//input:
	//var reader = bufio.NewReader(os.Stdin)
	//string_input, _ := reader.ReadString('\n')
	//string_input = strings.Trim(string_input, "\n")
	//s := strings.Split(string_input, " ")
	//
	//aAr := strings.Split(s[0], "")
	//bAr := strings.Split(s[1], "")
	//if len(aAr) > len(bAr) {
	//
	//}
	//
	//temporary := 0
	//if len(aAr) > len(bAr) {
	//
	//	res := make([]int, 0, len(aAr))
	//	for i := len(aAr) - 1; i >= 0; i-- {
	//
	//		//for i := 0; i < len(aAr); i++ {
	//		number := 0
	//		if i > len(bAr)-1 {
	//			fmt.Println("o")
	//			a, _ := strconv.Atoi(aAr[i])
	//			number = a + temporary
	//		} else if i > len(aAr)-1 {
	//			fmt.Println("oo")
	//
	//			b, _ := strconv.Atoi(bAr[i])
	//			number = b + temporary
	//		} else {
	//			fmt.Println("ooo")
	//
	//			a, _ := strconv.Atoi(aAr[i])
	//			b, _ := strconv.Atoi(bAr[i])
	//			number = a + b + temporary
	//		}
	//
	//		temporary = 0
	//		if number > 9 {
	//			number = number - 10
	//			temporary = 1
	//		}
	//		fmt.Println(res)
	//		res = append(res, number)
	//	}
	//	fmt.Println("res", res)
	//} else {
	//	for i := len(bAr) - 1; i >= 0; i-- {
	//
	//	}
	//}
	//
	//for _, v := range s {
	//	for _, char := range v {
	//		if !unicode.IsNumber(char) {
	//			fmt.Println("not digit")
	//		} else {
	//			fmt.Println(string(char), char, "is not a number rune")
	//		}
	//	}
	//
	//	//fmt.Println(v)
	//	//if !isNumeric(v) {
	//	//	log.Fatal("not digit")
	//	//}
	//	//a, _ := strconv.Atoi(v)
	//	//res = sum(res, a)
	//}
	//fmt.Println(s)
	//fmt.Println(len(s))

	//output:
	//fmt.Println(output)
}

func initial(path string, it interface{}) *httpserver.Server {
	m := printer(path, it)
	handler := gin.New()
	v1.NewRouter(handler, m)
	server := httpserver.New(handler, httpserver.Port("8080"))
	return server
}

func printer(name string, it interface{}) map[string]interface{} {

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &it)
	if err != nil {
		log.Fatal(err)
	}
	m, ok := it.(map[string]interface{})
	if !ok {
		log.Fatalf("want type map[string]interface{};  got %T", it)
	}
	//for k, v := range m {
	//	fmt.Println(k, "=>", v)
	//}

	return m
}
