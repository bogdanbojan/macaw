//go:build windows

package main

var (
	user32 = syscall.MustLoadDLL("user32.dll")
)

func main(){
    fmt.Println("We are inside windows")
}
