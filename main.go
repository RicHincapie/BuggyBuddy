// BuggyBuddy is a Golang server that fails a percentage of responses
// so you can test resiliency features.
package main

import (
	"fmt"
	"net/http"
	"BuggyBuddy/server"
	"BuggyBuddy/commons"
)



func main() {
	http.HandleFunc("/", server.HandleDefault)
	http.HandleFunc("/fail", server.HandleFail)
	http.HandleFunc("/setPercentage", server.HandleSetPercentage)
	http.HandleFunc("/getPercentage", server.HandleGetPercentage)
	fmt.Printf("Failing %d%% of reqs on /fail\n", commons.FailurePercentage)
	fmt.Println("Try /,/setPercentage?percentage=70 or /getPercentage")
	fmt.Println("\nServer listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
