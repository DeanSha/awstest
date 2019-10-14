package main
// add it to develop branch
import ( "log"
	"fmt"
	"time"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint hit at %v", time.Now())
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
