package main

import (
  "os"
  "log"

  "go-optimized-web-crawler/utils"
  "go-optimized-web-crawler/crawler"
)

const (
  routines int = 20
)

func getArgs() string {
  args := os.Args
  if len(args) < 2 {
    panic("empty link")
  }
  if len(args) > 2 {
    panic("pass one link")
  }
  link := args[1]
  valid := utils.ValidUri(link)
  if !valid {
    panic("invalid link")
  }
  return link
}

func main() {
  link := getArgs()
  c := crawler.New(routines)
  links := c.Crawler(link)
  log.Println("\n\n\nlinks found:")
  for _, l := range links {
    log.Println(l)
  }
}
