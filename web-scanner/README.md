# Web-Scanner

Параллельный сканер URL:  
читаем список из `urls.txt`, делаем HTTP GET с таймаутом 2s, выводим:
- URL
- Статус-код
- Размер тела (bytes)
- Время отклика

## Структура

- `go.mod` — модуль проекта
- `urls.txt` — входные URL
- `cmd/scanner/main.go` — CLI для запуска
- `pkg/scanner/scanner.go` — пул воркеров + HTTP-логика
- `pkg/scanner/scanner_test.go` — простые unit-тесты

## Запуск

```bash
go run cmd/scanner/main.go
