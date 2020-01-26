package main
// add it to develop branch
//changed on 01/25/2020
import ( "log"
	"fmt"
	"time"
	"net/http"
)
var cnt int = 0

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint hit at %v", time.Now())
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	cnt += 1
	fmt.Println("Request counter at %v is %d ", time.Now(), cnt) 
	log.Fatal(http.ListenAndServe(":8082", nil))
}
func main() {
	handleRequests()
}
