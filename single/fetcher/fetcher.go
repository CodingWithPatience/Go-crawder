package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var rate = time.Tick(500 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//limit the speed of request, 20 times per second
	<- rate
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetch url:%s error StatusCode:%d", url, resp.StatusCode)
	}
	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)
	s := string(bytes)
	s = strings.Replace(s, `\/`, `/`, -1)
	return []byte(s), err
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
