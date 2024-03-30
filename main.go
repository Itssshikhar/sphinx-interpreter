package main

import (
  "fmt"
  "os"
  "os/user"
  "sphinx/repl"
)

func main() {
  const INTRO_LOGO = `           _     _            
          | |   (_)           
 ___ _ __ | |__  _ _ __ __  __
/ __| '_ \| '_ \| | '_ \\ \/ /
\__ \ |_) | | | | | | | |>  < 
|___/ .__/|_| |_|_|_| |_/_/\_\
    | |                       
    |_|                       
` 
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Print(INTRO_LOGO)
  fmt.Printf("Aa gya Manhoos %s! This is the Sphinx programming language!\n", user.Username)
  fmt.Printf("Feel free to type in commands\n")
  repl.Start(os.Stdin, os.Stdout)
}
