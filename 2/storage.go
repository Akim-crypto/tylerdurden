package main

import (
    "encoding/json"
    "os"
)

func loadTasks(filename string) ([]Task, error) {
    file, err := os.Open(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return []Task{}, nil // Если файла нет — возвращаем пустой список
        }
        return nil, err // Другая ошибка (например, нет прав доступа)
    }
    defer file.Close()

    var tasks []Task
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

func saveTasks(filename string, tasks []Task) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(tasks)
}
