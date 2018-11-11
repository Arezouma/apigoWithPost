package main

import ("encoding/json"

"log"

"net/http"
"github.com/gorilla/mux"
"github.com/gorilla/handlers")

type Article struct {
        ID string `json: "id"`
        Title string `json: "title"`
        Image string `json: "image"`
        Author string `json: "author"`
}

 

var articles = []Article{

        Article{ID: "1", Image: "https://imagesvc.timeincapp.com/v3/mm/image?url=https%3A%2F%2Fpeopledotcom.files.wordpress.com%2F2018%2F07%2Fchristening-b.jpg&w=700&q=85", Title: "The Duke and Duchess of Cambridge, Prince George, Princess Charlotte and Prince Louis.", Author: "Matt Holyoak"},
        Article{ID: "2", Image: "http://www2.pictures.livingly.com/mp/zvOAvZDJb0vl.jpg", Title: "You will love these rare andstunning photos of Princess Diana", Author: "Kimia Madani"},
        Article{ID: "3", Image: "https://cbsnews1.cbsistatic.com/hub/i/r/2018/07/10/1350f97a-d083-455a-b994-845440650a84/resize/620x/f946cfb0cee567a2d52ced58120292da/thailand.jpg", Title: "People on the street cheered and clapped when ambulances ferrying the boys arrived at the hospital in Chiang Rai city.", Author: "CBS NEWS"},
        Article{ID: "4", Image: "https://assets.bwbx.io/images/users/iqjWHBFdfxIU/iADgPrqC7yoI/v0/1000x-1.jpg", Title: "Meet the 40-Year-Old Erdogan Son-in-Law Running Turkey’s Economy", Author: "Onur Ant"},
        Article{ID: "5", Image: "https://fm.cnbc.com/applications/cnbc.com/resources/img/editorial/2018/07/02/105306577-1530537945343rts1ub47.530x298.jpg?v=1530542057", Title: "Tesla ditches reservations, opens up Model 3 car sales to all customers in North America", Author: "Lora Kolodny"},
        Article{ID: "6", Image: "https://peopledotcom.files.wordpress.com/2018/07/meghan-harry-7-2000.jpg?crop=0px%2C0px%2C2000px%2C1334.2318059299px&resize=742%2C495", Title: "Meghan Markle and Prince Harry Meet Nelson Mandela's Close Friend at Powerful New Exhibit", Author: "David Fisher"}}

 

func GetData(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        json.NewEncoder(w).Encode(articles)

}

func GetItem(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     params := mux.Vars(r)// Get params
     for _, item := range articles {
        if item.ID == params["id"] {
                json.NewEncoder(w).Encode(item)
                return
        }
     }
     json.NewEncoder(w).Encode(&Article{})   
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     params := mux.Vars(r)
     var arti Article
     _= json.NewDecoder(r.Body).Decode(&arti)
     arti.ID = params["id"]
     articles = append(articles, arti)
     json.NewEncoder(w).Encode(arti)   
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     params := mux.Vars(r)
     for index, item := range articles {
        if item.ID == params["id"] {
           articles = append(articles[:index], articles[index+1:]...)
           var arti Article
           _= json.NewDecoder(r.Body).Decode(&arti)
           arti.ID = params["id"]
           articles = append(articles, arti)
           json.NewEncoder(w).Encode(arti)
           return
        }
     }
     json.NewEncoder(w).Encode(articles)   
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     params := mux.Vars(r)
     for index, item := range articles {
        if item.ID == params["id"] {
           articles = append(articles[:index], articles[index+1:]...)
           break
        }
     }
     json.NewEncoder(w).Encode(articles)   
}



 

func main() {

        router := mux.NewRouter()

        router.HandleFunc("/data", GetData).Methods("GET")
        router.HandleFunc("/data/{id}", GetItem).Methods("GET")
        router.HandleFunc("/data/{id}", CreateArticle).Methods("POST")
        router.HandleFunc("/data/{id}", UpdateData).Methods("PUT")
        router.HandleFunc("/data/{id}", DeleteData).Methods("DELETE")

        log.Print("localhost:8000")
        corsObj := handlers.AllowedOrigins([]string{"*"})

        log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))

}