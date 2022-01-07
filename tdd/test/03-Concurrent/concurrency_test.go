package concurrent_test

import (
	"reflect"
	"testing"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		result[url] = wc(url)
	}

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
