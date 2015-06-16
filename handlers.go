package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "io"
    "io/ioutil"
)

// Index is currently a placeholder for when the root
// server is hit. Simply displays "Welcome!" at the
// moment
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

// PostIndex returns the entire repo of messages.
func PostIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(posts); err != nil {
        panic(err)
    }
}

// PostShow returns the message with the given Id.
func PostShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postId := vars["postId"]

    i, err := strconv.Atoi(postId);

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    message := RepoFindPost(i)
    if err := json.NewEncoder(w).Encode(message); err != nil {
        panic(err)
    }
}

// Creates a new message and adds it to the repo.
func PostCreate(w http.ResponseWriter, r *http.Request) {
    var post Post
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &post); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreatePost(post)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}
