package scanner

import "testing"

func TestRun(t *testing.T) {
    urls := []string{
        "https://golang.org",
        "https://example.com",
    }
    results := Run(urls, 2)
    if len(results) != len(urls) {
        t.Fatalf("expected %d results, got %d", len(urls), len(results))
    }
    for _, r := range results {
        if r.URL == "" {
            t.Error("empty URL in result")
        }
        // допускаем как успешные коды, так и ошибки (invalid URL)
    }
}
