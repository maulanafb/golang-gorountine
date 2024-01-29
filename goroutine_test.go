package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(){
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}