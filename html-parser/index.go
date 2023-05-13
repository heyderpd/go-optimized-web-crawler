package htmlParser

import (
  "log"
  "net/http"
  "golang.org/x/net/html"

  "go-optimized-web-crawler/utils"
  "go-optimized-web-crawler/tail-recursion"
)

func getHtml(link string) *html.Node {
  if !utils.ValidUri(link) {
    return nil
  }
  res, err := http.Get(link)
  defer res.Body.Close()
  if err != nil {
    log.Println(err)
    return nil
  }
  doc, err := html.Parse(res.Body)
  if err != nil {
    log.Println(err)
    return nil
  }
  return doc
}

func Parser(link string, handlerNode utils.UnknownHandler) {
  page := getHtml(link)
  tr := tailRecursion.New()
  tr.Recursion(page, func (i interface{}) {
    node := (i).(*html.Node)
    if node == nil {
      return
    }
    handlerNode(node)
    for node := node.FirstChild; node != nil; node = node.NextSibling {
      tr.AppendTask(node)
    }
  })
}
