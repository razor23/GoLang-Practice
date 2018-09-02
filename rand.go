package main
import(
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := random(1, 6)
    fmt.Println(myrand)
}
