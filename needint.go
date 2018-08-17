package main
import "fmt"

const(

small=1<<100
Big = small>>99

)

func needint(x int) int{ return x*10+1}
func needfloat(x float64) float64 {return x*.1 }

func main(){
z := float64(small)
  fmt.Println(z,Big)
  fmt.Println(needint(Big))
	fmt.Println(needfloat(small))
	fmt.Println(needfloat(Big))
}
