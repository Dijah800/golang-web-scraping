package main
func home(w http.ResponseWriter, r *http.Request) {
 msg := "Hello, welcome to your app. Use the following suffix's on the URL to show the different results.\n1)'/scrape' to show results of web scraping.\n2)'/crawl' to show results of a web crawler"
 logr.Info("Received request for the home page")
  w.Write([]byte(fmt.Sprintf(msg)))
}
import (
     "fmt"
     "net/http"
     logr "github.com/sirupsen/logrus""net/http"
)
http.HandleFunc("/", home)
logr.Info("Starting up on 8080")
http.ListenAndServe(":8080", nil)
