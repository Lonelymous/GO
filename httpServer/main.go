package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Session struct {
	UserId   int
	Language string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var Sessions map[string]*Session //map[token]

func NewSession() (token string) {
	for {
		token = RandStringRunes(16)
		if _, found := Sessions[token]; !found {
			break
		}
	}
	Sessions[token] = &Session{
		UserId:   0,
		Language: "hu",
	}
	SaveSessions()
	return token
}

func ProcessSession(w http.ResponseWriter, r *http.Request) *Session {
	token := r.Header.Get("Token")
	if token == "" {
		token = NewSession()
	} else {
		if _, found := Sessions[token]; !found {
			token = NewSession()
		}
		fmt.Println("Hello újra: ", token)
	}
	w.Header().Set("Token", token)
	return Sessions[token]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	session := ProcessSession(w, r)
	fmt.Println("Handler")
	switch session.Language {
	case "hu":
		w.Write([]byte("Szia Lajos"))
	case "en":
		w.Write([]byte("Hi Sir"))
	default:
		w.Write([]byte("Invalid language"))
	}
}

func HandlerChange(w http.ResponseWriter, r *http.Request) {
	session := ProcessSession(w, r)
	fmt.Println(r.Body)
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Can't read body ("+err.Error()+")", http.StatusInternalServerError)
		return
	}
	bodyStr := string(body)
	fmt.Println(bodyStr)
	switch bodyStr {
	case "hu", "en":
		session.Language = bodyStr
	default:
		http.Error(w, "Invalid language code", http.StatusBadRequest)
	}
}

func SaveSessions() {
	data, err := json.Marshal(&Sessions)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("sessions.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := os.ReadFile("sessions.json")

	if err != nil {
		fmt.Println(err)
	}

	Sessions = make(map[string]*Session)
	json.Unmarshal(data, &Sessions)

	fmt.Println(string(data))

	spew.Dump(data)

	http.HandleFunc("/", Handler)
	http.HandleFunc("/change", HandlerChange)
	http.ListenAndServe(":4000", nil)
	jsonStr, err := json.Marshal(Sessions)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
