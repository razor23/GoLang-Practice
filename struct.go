package main
import "fmt"
type Car struct{
  hp int
  color string
  rims string
}

func main(){

  mycar:=Car{536,"red","silver"}
  fmt.Println(mycar)
  p:= &mycar
  fmt.Println(p)
  p.hp= 12312
  fmt.Println(mycar)

}
