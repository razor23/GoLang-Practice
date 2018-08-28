package main
import "fmt"

func main(){
i,j:= 80,90
var s int

p:=&i
fmt.Println(*p)
fmt.Println(p)
fmt.Println(&p)

*p=67

fmt.Println(i)
fmt.Println(j)

k:=&j
s=*p+*k
fmt.Println(s)

}
