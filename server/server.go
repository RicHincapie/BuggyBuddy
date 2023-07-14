package server

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"math/rand"
	"time"
	"BuggyBuddy/commons"
)

var (
	// FailurePercentage = 70
	mu                sync.Mutex
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Response Ok\n")
}

func HandleFail(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	// defer mu.Unlock()

	if shouldFail() {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 Internal Server Error")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Ok")
	}
}

func HandleSetPercentage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	percentageStr := r.URL.Query().Get("percentage")
	percentage, err := strconv.Atoi(percentageStr)
	USAGE := "USAGE: setPercentage?percentage=10"
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid percentage value.\n", USAGE)
		return
	}

	if percentage < 0 || percentage > 100 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Percentage value out of range")
		return
	}

	commons.FailurePercentage = percentage
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Now will fail %d%% of requests", commons.FailurePercentage)
}

func HandleGetPercentage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Currently failing ", commons.FailurePercentage, "%")
}

func shouldFail() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) < commons.FailurePercentage
}