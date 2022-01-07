package concurrent_test

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			result[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return result
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResult := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResult)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(expectedResults, actualResult) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResult)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
