package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/Tom-Kail/gocrawl"
)

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	fmt.Printf("Visit: %s\n", ctx.URL())
	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}
	//	if ctx.URL().Host == "github.com" || ctx.URL().Host == "golang.org" || ctx.URL().Host == "0value.com" {
	//		return true
	//	}
	return true
}

// End is a no-op.
func (e *Ext) End(err error) {
	upload := make([]*gocrawl.URLContext, 0)
	for {
		if len(e.EnqueueChan) == 0 {
			break
		}
		select {
		case ctx := <-e.EnqueueChan:
			upload = append(upload, ctx.(*gocrawl.URLContext))
			fmt.Println("Url left in EnqueueChan:", ctx.(*gocrawl.URLContext).URL().String())
		}
	}
	fmt.Println("***************************************")
	fmt.Println("upload url number:", len(upload))
	fmt.Println("***************************************")
}

func main() {

	//	x := make(chan interface{}, 1000)
	now := time.Now()
	ext := &Ext{&gocrawl.DefaultExtender{}}

	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 10 * time.Millisecond
	opts.LogFlags = gocrawl.LogError
	opts.EnqueueChanBuffer = 1000
	opts.WokerPoolSize = 3000
	opts.SameHostOnly = true
	opts.MaxVisits = 10000

	c := gocrawl.NewCrawlerWithOptions(opts)
	err := c.Run("http://www.gov.cn")
	//	err := c.Run("http://192.168.133.134")
	fmt.Println(err)
	fmt.Println(time.Since(now))
}
