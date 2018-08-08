package middleware

import (
	"context"
	"net/http"
	"os"

	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("Auth Middle Ware.")
	return func(w http.ResponseWriter, r *http.Request) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token."))
			return
		}
		log.Printf("Verifying ID token: %v", token)
		next.ServeHTTP(w, r)
	}
}
