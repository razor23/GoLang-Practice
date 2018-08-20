package main
import "fmt"
import "math"

func power(x,n,lim float64) float64{
  if v:=math.Pow(x,n);v<lim{
    return v
  }
  return lim
}

func main(){
  fmt.Println(power(3,2,3))
  fmt.Println(power(3,3,41))
}
