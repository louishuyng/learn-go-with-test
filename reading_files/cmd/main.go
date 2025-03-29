package main

import (
	"log"
	"os"
	blogposts "reading_files/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		log.Println(post)
		log.Println("----")
	}
}
