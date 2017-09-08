package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
)

func main()  {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("add a new note: ")
  newNote, err :=reader.ReadString('\n')
  if err != nil {
    fmt.Println(err)
  } else {
    //write contents of newNote to file
    file, err := os.OpenFile("/tmp/output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    fmt.Fprintf(file, newNote)

    //check output
    fmt.Println(newNote)
  }

}
