package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/MariaRogulenko/lines"
	"github.com/MariaRogulenko/lines/api"
)

var (
	gRPCPort  = flag.Int("grpc-port", 8090, "GRPC Server port")
	httpPort  = flag.Int("http-port", 8080, "HTTP Server port")
	staticDir = flag.String("static-dir", "./static", "Folder to serve html/js/css")
)

// runService runs the gRPC service.
func runService(service *lines.Service) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	api.RegisterGameServer(srv, service)
	go func() {
		if err := srv.Serve(l); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

// registerRPCProxy registers the HTTP handler to proxy REST to gRPC.
func registerRPCProxy(ctx context.Context) {
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterGameHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", *gRPCPort), dialOpts); err != nil {
		log.Fatal(err)
	}
	http.Handle("/api/", mux)
}

// registerHTMLHandler registers the HTTP handler to server HTML/JS/CSS.
func registerHTMLHandler() {
	indexPage, err := ioutil.ReadFile(fmt.Sprintf("%s/index.html", *staticDir))
	if err != nil {
		log.Fatalf("failed to read file %q: %v", fmt.Sprintf("%s/index.html", *staticDir), err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write(indexPage) })
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(*staticDir))))
}

func main() {
	flag.Parse()
	s := lines.NewService()
	lines.RegisterAndOpenDB()
	//lines.WriteDB()
	runService(s)
	registerRPCProxy(context.Background())
	registerHTMLHandler()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil))
}
