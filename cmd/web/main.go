package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/archit-p/MicroserviceTemplate/pkg/models"
	"github.com/archit-p/MicroserviceTemplate/pkg/models/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// application represents an app with custom loggers and access to
// a database
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	samples  models.SampleModel
}

// newApplication creates an application with given parameters
func newApplication(errorLog, infoLog *log.Logger, samples models.SampleModel) *application {
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		samples:  samples,
	}
}

// createDBConnection starts a connection with MongoDB at the
// given URL
func createDBConnection(dbURL string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	addr := flag.String("addr", ":4000", "Address to start server on")
	dbURL := flag.String("db", "mongodb://localhost:27017", "Address for MongoDB server")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	client, err := createDBConnection(*dbURL)
	if err != nil {
		log.Fatal(err)
	}

	sampleModel := &mongodb.SampleMongo{
		Collection: client.Database("example").Collection("sample"),
	}

	app := newApplication(errorLog, infoLog, sampleModel)

	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	go func() {
		infoLog.Println("Starting server on ", *addr)

		err := srv.ListenAndServe()
		if err != nil {
			errorLog.Fatalf("Error starting server: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
