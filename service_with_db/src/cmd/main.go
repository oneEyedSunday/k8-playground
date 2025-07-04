package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/lib/pq"
)

var (
	dummySelectStatement *sql.Stmt
)

const (
	// Database
	dummySelect = "SELECT CURRENT_TIMESTAMP, RANDOM();"
)

func json(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

func plaintext(c *gin.Context) {
	c.String(200, "Hello, World!")
}

func db(c *gin.Context) {
	var res struct {
		CurrentTimestamp time.Time `json:"now"`
		Rand             float64   `json:"etag"`
	}
	err := dummySelectStatement.QueryRowContext(c.Request.Context()).Scan(&res.CurrentTimestamp, &res.Rand)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, &res)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Expose the default metrics for Go applications via http://localhost:{PORT}/metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// http.Handle("/metrics", promhttp.Handler())
	r.GET("/json", json)
	r.GET("/db", db)
	r.POST("/", insertStuff)
	r.GET("/stuff", getStuff)

	// TODO add POST route to test persistence

	r.GET("/plaintext", plaintext)

	// Start the HTTP server on port default port 80 or passed in port
	fmt.Println("Listening and serving HTTP")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(fmt.Sprintf(":%s", port))
}

func init() {
	// TODO use the uber package so it picks up limits from k8s
	// runtime.GOMAXPROCS(runtime.NumCPU())

	dbhost := os.Getenv("DB_HOST")
	port := 5432
	user := "app_user"
	password := os.Getenv("DB_PASSWORD")
	dbname := "database1"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbhost, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(3)
	dummySelectStatement, err = db.Prepare(dummySelect)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}
