package main

import "time"

type Post struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Message   string    `json:"message"`
    Time      time.Time `json:"time"`
}

type Posts []Post
