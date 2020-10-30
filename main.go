package main
import (
	"fmt"
	"net/http"
)
func main() {
	fmt.Println("hola mundo")
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":3000", nil)
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	fmt.Println(ip)
	http.Redirect(w, r, "https://www.twitch.tv/freddyfalso", 301)
	fmt.Fprintln(w, "hola mundo")
}
