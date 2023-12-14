// main.go
package main

import (
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"webrtc-go/pkg"
)

func main() {
	pkg.AllRooms.Init()

	r := gin.Default()

	r.GET("/create", pkg.CreateRoomRequestHandler)
	r.GET("/join", pkg.JoinRoomRequestHandler)

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	// HTML5 history mode support
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	err := r.Run(":8000")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	// certFile := "cert.pem"
	// keyFile := "key.pem"

	// err := r.RunTLS(":443", certFile, keyFile)
	// if err != nil {
	// 	log.Fatal("Failed to start server: ", err)
	// }
}
