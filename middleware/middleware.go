package middleware

import (
    "encoding/json"
    "log"
    "net/http"
)

func Recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err) // Usar log en lugar de fmt.Println

                jsonBody, _ := json.Marshal(map[string]string{
                    "error": "There was an internal server error",
                })

                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                w.Write(jsonBody)
            }
        }()

        next.ServeHTTP(w, r)
    })
}