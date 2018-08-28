package main
import "fmt"
func main(){

hieghts:= [10]int{1,2,3,4,1}
var s[]int=hieghts[3:7]

fmt.Println(hieghts)

j:=hieghts[1:5]
j[0]=12321
fmt.Println(hieghts)
fmt.Println(s)



}
