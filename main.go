package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

// site référence: https://go.dev/doc/tutorial/web-service-gin

// creation d'un type album
type album struct{
	ID	string `json:"id"`
	Title 	string `json:"title"`
	Artist	string `json:"artist"`
	Price	float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "John Coltrane", Price: 58.85},
	{ID: "2", Title: "Jeru", Artist: "Gery mullighan", Price: 45.25},
	{ID: "3", Title: "Yellow plane", Artist: "John Silver", Price: 15.96},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	// creation d'une reference newAlbum de type album
	var newAlbum album

	// appel BinfJSON pour lier le JSON reçu avec newAlbum
	if err:= c.BindJSON(&newAlbum); err != nil{
		return
	}

	// ajout de new album à la liste des albums déjà crée
	albums = append(albums, newAlbum)
	
	// on retourne une http reponse 201 avec le nouvel album crée
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context){
	// on recupere le parametre nommé id de l'url
	id := c.Param("id")

	// on boucle sur la liste d'albums pour recherche un album par son id
	for _,a := range albums{
		if a.ID == id {
		c.IndentedJSON(http.StatusOK, a)
		return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album not found"})
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:9090")
}
