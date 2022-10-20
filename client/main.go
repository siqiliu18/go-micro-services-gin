package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	ginclient_proto "ginclient/proto"
)

// Client is a struct
type Client struct {
	ginClient ginclient_proto.ComputeServiceClient
}

func (c *Client) GuiFunc(ctx *gin.Context) {
	a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
		return
	}

	b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
		return
	}

	in := ginclient_proto.GcdRequest{
		A: a,
		B: b,
	}
	if res, err := c.ginClient.Compute(context.TODO(), &in); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func main() {
	fmt.Println("Go gRPC client with gin")
	conn, err := grpc.Dial("gcd-service-server:7000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dail server failed: %v", err)
	}
	defer conn.Close()

	clientService := Client{
		ginClient: ginclient_proto.NewComputeServiceClient(conn),
	}

	gateway := gin.Default()
	fmt.Println(gateway)

	gateway.GET("/gcd/:a/:b", clientService.GuiFunc)

	if err := gateway.Run(":7003"); err != nil {
		log.Fatalf("Failed to run client service: %v", err)
	}
}
