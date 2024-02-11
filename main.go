package main
	
import (
    "fmt"
)
func demo(ab ... int) {

    for _, num:=range ab{
        fmt.Println(num)
    }
}
func main() {
    demo(1,2,3,4)

}
