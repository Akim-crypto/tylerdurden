package scanner

import (
    "context"
    "io"
    "net/http"
    "sync"
    "time"
)

// Result содержит информацию о выполненном запросе.
type Result struct {
    URL        string
    StatusCode int
    Bytes      int
    Duration   time.Duration
    Error      error
}

// worker читает URL-ы из канала jobs, выполняет HTTP GET с таймаутом и
// отправляет Result в канал results.
func worker(jobs <-chan string, results chan<- Result) {
    for url := range jobs {
        start := time.Now()

        // создаём контекст с таймаутом 2 секунды
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        defer cancel()

        // формируем запрос с контекстом
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
        if err != nil {
            results <- Result{URL: url, Error: err}
            continue
        }

        // выполняем запрос
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            results <- Result{URL: url, Error: err}
            continue
        }
        // гарантированно закрываем тело ответа
        defer resp.Body.Close()

        // читаем весь ответ
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            results <- Result{URL: url, Error: err}
            continue
        }

        // отправляем результат
        results <- Result{
            URL:        url,
            StatusCode: resp.StatusCode,
            Bytes:      len(body),
            Duration:   time.Since(start),
        }
    }
}

// Run запускает numWorkers горутин-воркеров для параллельного сканирования URL-ов.
// Возвращает срез Result той же длины, что и входной список urls.
func Run(urls []string, numWorkers int) []Result {
    jobs := make(chan string, len(urls))
    results := make(chan Result, len(urls))

    var wg sync.WaitGroup

    // 1) Запускаем воркеров
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            worker(jobs, results)
        }()
    }

    // 2) Отправляем URL-ы на обработку
    for _, url := range urls {
        jobs <- url
    }
    close(jobs)

    // 3) Ждём завершения всех воркеров и закрываем канал с результатами
    wg.Wait()
    close(results)

    // 4) Собираем результаты в срез
    var out []Result
    for res := range results {
        out = append(out, res)
    }
    return out
}
