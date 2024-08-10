package starter

import (
  "fmt"
  "math"
  "net/http"

)

func SayHello(name string) string {
  return `Olá, ` + name + "!! Seja bem vindo!!"
}

func OddOrEven(num int) string {
  criteria := math.Mod(float64(num), 2)
  if criteria == 1 || criteria == -1 {
    return fmt.Sprintf("%v is an odd number", num)
  }
  return fmt.Sprintf("%v is an even number", num)
}

func CheckHealth(writer http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(writer, "health check passed")
}