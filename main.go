package main

import (
  "fmt"
  "os"
  "log"
  "io/ioutil"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Commands = []cli.Command{
    {
      Name:    "add",
      Aliases: []string{"a"},
      Usage:   "add a task to the list",
      Action:  func(c *cli.Context) error {
        file, err := os.OpenFile("/tmp/test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            log.Fatal("Cannot create file", err)
        }
        defer file.Close()

        fmt.Fprintf(file, c.Args().First() + "\n")
        fmt.Println("added task: ", c.Args().First())
        return nil
      },
    },
    {
      Name:    "list",
      Aliases: []string{"l"},
      Usage:   "list all todos",
      Action:  func(c *cli.Context) error  {
        //open file and read contents
        b, err := ioutil.ReadFile("/tmp/test.txt")
        if err != nil {
            fmt.Print(err)
        }

        //print file contents
        str := string(b) // convert content to a 'string'
        fmt.Println(str) // print the content as a 'string'
        return nil
      },
    },
    {
      Name:    "complete",
      Aliases: []string{"c"},
      Usage:   "complete a task on the list",
      Action:  func(c *cli.Context) error {
        fmt.Println("completed task: ", c.Args().First())
        return nil
      },
    },

  }

  app.Run(os.Args)
}
