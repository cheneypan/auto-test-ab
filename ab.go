package main

import (
	"net/http"
	"fmt"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"bufio"
	"io"
)

type ABForm struct {
	Url string
	C   int
	N   int
}

func main() {
	// static files
	http.Handle("/", http.FileServer(http.Dir("./html/")))

	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			fmt.Printf("Only support POST method")
			return
		}
		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()
		fmt.Printf("POST %s\n", result)

		// ab -c 10 -n 100 http://member.dev.cloudtogo.cn/index.html
		var abForm ABForm
		json.Unmarshal([]byte(result), &abForm)

		str := fmt.Sprintf("ab -c %d -n %d %s \n", abForm.C, abForm.N, abForm.Url)
		fmt.Print(str)

		cmd := exec.Command("/bin/bash", "-c", str)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		cmd.Start()
		reader := bufio.NewReader(stdout)
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			//fmt.Println(line)
			fmt.Fprint(w, line)
		}
		cmd.Wait()
	})
	http.ListenAndServe(":80", nil)
}
