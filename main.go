package main

import (
	"fmt"
)

func main() {
  var day string
	fmt.Println("Дни недели.\nВведите будний день недели: пн, вт, ср, чт, пт:")
fmt.Scan(&day)
  switch day{
    case "пн":
      fmt.Println("понедельник")
      day = "вт"
      fallthrough
    case "вт":
      fmt.Println("вторник")
      day = "ср"
      fallthrough
     case "ср":
      fmt.Println("среда")
      day = "чт"
      fallthrough
     case "чт":
      fmt.Println("четверг")
      day = "пт"
      fallthrough
     case "пт":
      fmt.Println("пятница")
    default:
      fmt.Println("Введенное значение не известно")
  }
}
