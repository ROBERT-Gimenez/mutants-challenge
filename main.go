package main

import (
	"Challenge/api/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	/* "context"
	"net/http/httptest"
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda" */
)

var router = mux.NewRouter()
 
func init() {
	routes.MutantsRoutes(router)
}
/*  Codigo en Lampda
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Recibiendo solicitud: %s %s", request.HTTPMethod, request.Path)
	
	reqBody := []byte(request.Body)
	req, err := http.NewRequest(request.HTTPMethod, request.Path, bytes.NewReader(reqBody))
	if err != nil {
		log.Println("Error creando la solicitud HTTP:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	return events.APIGatewayProxyResponse{
		StatusCode: recorder.Code,
		Body:       recorder.Body.String(),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}


func main() {
	lambda.Start(handler)
}
  */

func main() {
	routes.MutantsRoutes(router)

	 log.Println("Servidor iniciado en http://localhost:8080")
	 log.Fatal(http.ListenAndServe(":8080", router))
 }