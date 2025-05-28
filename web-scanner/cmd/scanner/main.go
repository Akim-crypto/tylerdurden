package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"

    "example.com/web-scanner/pkg/scanner"
)

func main() {
    // 1. Открываем файл urls.txt
    file, err := os.Open("urls.txt")
    if err != nil {
        log.Fatalf("cannot open urls.txt: %v", err)
    }
    defer file.Close()

    // 2. Читаем URL-ы построчно, пропуская пустые
    var urls []string
    sc := bufio.NewScanner(file)
    for sc.Scan() {
        line := strings.TrimSpace(sc.Text())
        if line == "" {
            continue
        }
        urls = append(urls, line)
    }
    if err := sc.Err(); err != nil {
        log.Fatalf("error reading urls.txt: %v", err)
    }
    if len(urls) == 0 {
        log.Fatal("no URLs to scan in urls.txt")
    }

    // 3. Запускаем пул воркеров
    const numWorkers = 5
    results := scanner.Run(urls, numWorkers)

    // 4. Сортируем результаты по URL для предсказуемого порядка
    sort.Slice(results, func(i, j int) bool {
        return results[i].URL < results[j].URL
    })

    // 5. Печатаем на экран
    for _, r := range results {
        if r.Error != nil {
            fmt.Printf("%-30s ERROR: %v\n", r.URL, r.Error)
        } else {
            fmt.Printf("%-30s %3d   %7d bytes   %s\n",
                r.URL, r.StatusCode, r.Bytes, r.Duration)
        }
    }
}
