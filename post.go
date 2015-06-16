package main

import "time"

type Todo struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Message   string    `json:"message"`
    Time      time.Time `json:"time"`
}

type Todos []Todo
