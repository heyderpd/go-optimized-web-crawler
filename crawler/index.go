package crawler

import (
  "log"
  "context"
  "sync"
  "time"
  "golang.org/x/net/html"

  "go-optimized-web-crawler/set"
  "go-optimized-web-crawler/utils"
  "go-optimized-web-crawler/html-parser"
  "go-optimized-web-crawler/round-robin"
)

type Crawler struct {
  routines int
  domain string
  linkSet *set.Set
  links chan string
}

func New(routines int) *Crawler {
  s := Crawler{
    routines: routines,
    linkSet: set.New(),
    links: make(chan string),
  }
  return &s
}

func (s *Crawler) isValidDomain(link string) bool {
  return s.domain == utils.ParseDomain(link)
}

func (s *Crawler) isElementA(node *html.Node) bool {
  return node.Type == html.ElementNode && node.Data == "a"
}

func (s *Crawler) getHrefAttr(node *html.Node) string {
  for _, a := range node.Attr {
    if a.Key == "href" {
      return a.Val
    }
  }
  return ""
}

func (s *Crawler) pushLink(originalLink string) {
  link := utils.ParseLink(originalLink)
  if s.isValidDomain(originalLink) && !s.linkSet.Has(link) {
    s.linkSet.Add(link)
    s.links <- originalLink
  }
}

func (s *Crawler) findLinks(i interface{}) {
  node := (i).(*html.Node)
  if s.isElementA(node) {
    link := s.getHrefAttr(node)
    s.pushLink(link)
  }
}

func (s *Crawler) work(link string) {
  log.Println("open", link)
  htmlParser.Parser(link, s.findLinks)
}

func (s *Crawler) Crawler(link string) []interface{} {
  s.domain = utils.ParseDomain(link)
  ctx, cancelCtx := context.WithCancel(context.Background())
  waitGroup := sync.WaitGroup{}
  rr := roundRobin.New(s.routines)
  go func() {
    s.pushLink(link)
    time.Sleep(100 * time.Millisecond)
    waitGroup.Wait()
    cancelCtx()
  }()
  for {
    select {
    case link := <- s.links:
      waitGroup.Add(1)
      rr.Promise(func (i interface{}) {
        s.work(link)
        waitGroup.Done()
      })
    case <- ctx.Done():
      return s.linkSet.ToList()
    }
  }
}
