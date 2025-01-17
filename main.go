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

    // Обработчик для /go
    http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        html := `
        <!DOCTYPE html>
        <html>
        <head>
            <style>
                body {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                    font-family: Arial, sans-serif;
                    font-size: 24px;
                }
            </style>
        </head>
        <body>
            <div>хуй будешь?</div>
        </body>
        </html>`
        w.Write([]byte(html))
    })

    // Настраиваем обработку статических файлов
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Сервер запущен на http://0.0.0.0:8000")
    if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
        log.Fatal(err)
    }
} 