package utils

import (
  "net/url"
)

func ValidUri(link string) bool {
  _, err := url.ParseRequestURI(link)
  return err == nil
}

func parseUri(link string) (string, string) {
  uri, err := url.Parse(link)
  if err != nil {
    return "", ""
  }
  if uri.Path == "/" {
    return uri.Host, uri.Host
  }
  return uri.Host, uri.Host + uri.Path
}

func ParseDomain(originalLink string) string {
  domain, _ := parseUri(originalLink)
  return domain
}

func ParseLink(originalLink string) string {
  _, link := parseUri(originalLink)
  return link
}
