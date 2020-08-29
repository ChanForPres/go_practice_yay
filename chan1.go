package main

import "fmt"

func main() {
  msg := make(chan string)

  msg <- "true"
  msg <- "false"

  go fmt.Println(<- msg)
  go fmt.Println(<-msg)
  
  // wg to wait for return 
  

}
