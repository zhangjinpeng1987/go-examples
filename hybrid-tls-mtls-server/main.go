package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	pb "github.com/zhangjinpeng1987/go-examples/hybrid-tls-mtls-server/pkg/proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RPC Server
type rpcServer struct {
	pb.UnimplementedYourServiceServer
}

func (s *rpcServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Greeting: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

// Http Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

// func (s *rpcServer) mustEmbedUnimplementedYourServiceServer() {}

func main() {
	// Specify the paths to your server certificate and private key for both gRPC and HTTP
	certFile := flag.String("cert", "", "cert file")
	keyFile := flag.String("key", "", "key file")

	// Specify the path to your CA certificate for client verification
	httpCACertFile := flag.String("ca-cert", "", "ca cert file")

	// Specify the port to use for gRPC and HTTP/2
	addr := flag.String("addr", ":443", "addr")

	// mTLS or TLS
	mtls := flag.Bool("mtls", false, "use mtls or tls, 1 means mtls, 0 means tls, default is 0")

	flag.Parse()

	// Load the server certificate and private key
	serverCert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal("Error loading server certificate and key: ", err)
	}

	// Load the CA certificate for client verification
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(*httpCACertFile)
	if err != nil {
		log.Fatal("Error loading HTTP CA certificate: ", err)
	}
	if !certPool.AppendCertsFromPEM(ca) {
		log.Fatal("Failed to add HTTP CA certificate")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
	}
	if *mtls {
		log.Println("use mTLS")
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	// Create a TCP listener for both gRPC and HTTP
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("Error creating listener: ", err)
	}

	tlsListener := tls.NewListener(listener, tlsConfig)

	// Create a new multiplexer
	mux := cmux.New(tlsListener)

	// // Create a gRPC listener with more precise matching
	grpcListener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// // Create an HTTP listener for any remaining traffic
	httpListener := mux.Match(cmux.HTTP1Fast(), cmux.HTTP2())

	// Create a gRPC server with TLS configuration
	grpcServer := grpc.NewServer()

	// Register your gRPC service
	pb.RegisterYourServiceServer(grpcServer, &rpcServer{})

	// Enable reflection for gRPC
	reflection.Register(grpcServer)

	// Serve gRPC and HTTP using cmux
	go func() {
		// Start serving gRPC on a separate goroutine
		log.Println("gRPC server is now running on " + *addr)
		err := grpcServer.Serve(grpcListener)
		if err != nil {
			log.Fatal("gRPC server error: ", err)
		}
	}()

	// Create a Gin router
	router := gin.New()

	// Register your HTTP handlers
	router.GET("/hello", func(c *gin.Context) {
		log.Println("received request")
		c.String(http.StatusOK, "Hello, World!")
		log.Println("sent response")
	})

	// Create a new HTTP server
	httpServer := &http.Server{
		Handler: router,
	}

	// Serve HTTP using the same listener as gRPC
	go func() {
		log.Println("HTTP server is now running on" + *addr)
		err := httpServer.Serve(httpListener)
		if err != nil {
			log.Fatal("HTTP server error: ", err)
		}
	}()

	// Start serving muxed connections
	err = mux.Serve()
	if err != nil {
		log.Fatal("cmux server error: ", err)
	}
}
