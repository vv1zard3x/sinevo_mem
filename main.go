package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    // Загружаем HTML шаблон
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal(err)
    }

    // Обработчик для корневого маршрута
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        err := tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    // Настраиваем обработку статических файлов
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Сервер запущен на http://0.0.0.0:8000")
    if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
        log.Fatal(err)
    }
} 