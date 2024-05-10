package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/ramd"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
    title string `json:"title"`
	Director *Director `json:director"`
}
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string  `json;"Lastname"'
} 
var movies []Movie
fun getMovies(w http.ResposeWriter,r"http.Request){
	w.header().Set("Content:Type","application/json")}

	json.NewEncoder(w,http.ResponseWriter,r ).Encode(movies) 
func  deleteMovie(w http.ResponseWriter,r "http.Request)
w.Header()Set("Content Type","application/json")
params :-nux.Vars(r)
for index,item ;- range movies {
	if item ID== params["id"]{
		moveis= append(movies:[index]),movies[index+1:]...)
		break 
	}
}
for getMovies(w http.ResposeWriter,r "https "
}
func main(){
	movies= append(movies,Movie(ID:"1",Isbn:"45455",title:"Movie One",Director :&Director(Firstname:"John",Lastname:"Doe")))
	movies= append(movies,Movie(ID:"2",Isbn:"45435",title:"Movie Two",Director :&Director(Firstname:"Steve",Lastname:"Smith")))
	r :=mux.NewRounter()
	r.HandleFunc("/movies",getMovies).Methods("GET") 
	r.HandlerFunc("/movies/(id)",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.handleFunc("/movies/(id)",updateMovie).Methods("PUT")
	r.handleFunc("/movies/(id)",deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000/n")
	log.Fatal(http.ListenAndServer(:"8000",r )) 





} 