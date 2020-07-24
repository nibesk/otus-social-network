// sessions.go
package main

//
//import (
//	"fmt"
//	"math/rand"
//	"net/http"
//
//	"github.com/gorilla/sessions"
//)
//
//var (
//	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
//	key = []byte("super-secret-key")
//	store = sessions.NewCookieStore(key)
//)
//
//func secret(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Check if user is authenticated
//	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
//		http.Error(w, "Forbidden", http.StatusForbidden)
//		return
//	}
//
//	// Print secret message
//	fmt.Fprintln(w, "The cake is a lie!" + fmt.Sprintf("userId = %d", session.Values["userId"].(int)))
//}
//
//func login(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Authentication goes here
//	// ...
//
//	// Set user as authenticated
//	session.Values["authenticated"] = true
//	session.Values["userId"] = rand.Intn(10000)
//	session.Save(r, w)
//
//	// Print secret message
//	fmt.Fprintln(w, fmt.Sprintf("userId = %d", session.Values["userId"].(int)))
//}
//
//func logout(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "cookie-name")
//
//	// Revoke users authentication
//	session.Values["authenticated"] = false
//	session.Save(r, w)
//}
//
//func main() {
//	http.HandleFunc("/secret", secret)
//	http.HandleFunc("/login", login)
//	http.HandleFunc("/logout", logout)
//
//	http.ListenAndServe(":8080", nil)
//}
