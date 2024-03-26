package main

import "fmt"

type Page struct {
	Prev       *Page
	Next       *Page
	CurrentURL string
}

type BrowserHistory struct {
	Homepage *Page
}

func Constructor(homepage string) BrowserHistory {
	page := &Page{CurrentURL: homepage}

	return BrowserHistory{Homepage: page}
}

func (this *BrowserHistory) Visit(url string) {
	page := &Page{
		CurrentURL: url,
		Prev:       this.Homepage,
	}

	this.Homepage.Next = page
	this.Homepage = page
}

func (this *BrowserHistory) Back(steps int) string {
	for this.Homepage.Prev != nil {
		this.Homepage = this.Homepage.Prev
		if steps--; steps == 0 {
			break
		}
	}

	return this.Homepage.CurrentURL
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.Homepage.Next != nil {
		this.Homepage = this.Homepage.Next
		if steps--; steps == 0 {
			break
		}
	}

	return this.Homepage.CurrentURL
}

func main() {
	browser := Constructor("google.com")
	browser.Visit("leetcode.com")
	browser.Visit("facebook.com")
	browser.Visit("youtube.com")
	fmt.Println(browser.Homepage)
	fmt.Println(browser.Back(2))
}
