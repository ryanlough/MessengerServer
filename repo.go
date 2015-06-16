package main

import (
    "fmt"
    "time"
)

var currentId int

var posts Posts

// Used to initialize the repo with data. Empty at the moment.
func init() {

}

func RepoFindPost(id int) Post {
    for _, p := range posts {
        if p.Id == id {
            return p
        }
    }
    // return empty Post if not found
    return Post{}
}

func RepoCreatePost(p Post) Post {
    currentId += 1
    p.Id = currentId
    p.Time = time.Now().Local()
    posts = append(posts, p)
    return p 
}

func RepoDestroyPost(id int) error {
    for i, p := range posts {
        if p.Id == id {
            posts = append(posts[:i], posts[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Post with id of %d to delete", id)
}
