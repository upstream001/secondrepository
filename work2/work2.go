package main
import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://blog.lenconda.top/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	f,err:=os.Create("C:/Users/22391/Desktop/work2/数据.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	doc,err:=goquery.NewDocumentFromReader(res.Body)
	doc.Find("h1").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())

	})
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())
	})
	doc.Find("time").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())
	})
	doc.Find("div[box post-box]:nth-child(3)").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())
	})
	doc.Find("div[box post-box]:nth-child(4)").Each(func(i int, selection *goquery.Selection) {
		f.WriteString(selection.Text())
	})

}