package main

// import "fmt"
import (
	"jam-room/controllers"
	"jam-room/initializers"
	// "net/http"

	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")
}