package main

import "fmt"

func main() {

 angka := 8
 listAngka := []int{0, 2, 5, 6, 8, 11, 19}

 fmt.Println(cariAngka(angka, listAngka))

}

func cariAngka(angka int, listAngka []int) int {

 for i := range listAngka {
  if angka == listAngka[i] {
   return i
  }
 }

 return -1
}