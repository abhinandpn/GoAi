package Router

import (
	database "github.com/abhinandpn/GoCassandra/pkg/Database"
	"github.com/abhinandpn/GoCassandra/pkg/Users"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {

	// r := database.SetupDBConnection()
	r, err := database.SetupStorage()

	if err != nil {
		fmt.Println(err)
	}

	us := Users.NewService(r)

	router := chi.NewRouter()
	router.Mount("/api/users", Users.UsersRoutes(us))
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	return router
}
