package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "unavailable.site" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{"http://google.com", "https://gnu.org", "unavailable.site"}

	got := CheckWebsites(mockWebsiteChecker, urls)
	expected := map[string]bool{"http://google.com": true, "https://gnu.org": true, "unavailable.site": false}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v \nexpected %v", got, expected)
	}
}

func slowMockWebsiteChecker(url string) bool {
	time.Sleep(200 * time.Millisecond)
	return mockWebsiteChecker(url)
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "rand url"
	}

	for b.Loop() {
		CheckWebsites(slowMockWebsiteChecker, urls)
	}
}
