package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 內容，並且將在這個頁面上找到的 URL 放到一個 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var (
	//  map  存放爬取的url
	m = make(map[string]int)
	// 互斥鎖
	l sync.Mutex
	//   群組等待     當新增的任務沒有完成時（done（））， wait（） 會一直等待    三個方法   Add()  Done()  Wait()
	i sync.WaitGroup
)

func main() {
	i.Add(1)
	Crawl("http://golang.org/", 4, fetcher)

	i.Wait() // 會一直等待直到子執行緒任務結束

	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("over")
}

// Crawl 使用 fetcher 從某個 URL 開始遞迴的爬取頁面，直到達到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {

	defer i.Done() //  和add相對應
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		//  fmt.Println(err)
		return
	}
	// 存入資料  需要同步鎖  因為這是在子執行緒中
	l.Lock()
	if m[url] == 0 { //  還未爬取過
		m[url]++ // 存入爬取的url  改變對應的標示
		depth--
		//      fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			i.Add(1)
			go Crawl(u, depth, fetcher) // 繼續爬取
		}
	}
	l.Unlock()

}

// fakeFetcher 是返回若干結果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充後的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
