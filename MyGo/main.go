package main
import "fmt"

type User struct {
    name string
}

func (b User) cal(weight,height float64) (result float64) {
         result = weight / height / height * 10000
         return
}

func main(){
    s := User{name: "a"}
    fmt.Println(s.name , s.cal(75, 170))
}