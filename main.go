package main

import s "github.com/starptech/go-web/server"

func main() {
	s := s.New()
	s.Logger.Fatal(s.Start(":8080"))
}
