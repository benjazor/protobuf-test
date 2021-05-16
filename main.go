package main

import (
	"fmt"
	"log"

	"github.com/benjazor/protobuf-test/pkg/article"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Init article using generated constructor
	myArticle := &article.Article{
		Id:      1,
		Title:   "First Article",
		Content: "Very cool article, this is the first article 😋",
		Likes:   69,
	}

	// Marshaling
	data, err := proto.Marshal(myArticle)
	if err != nil {
		log.Fatal("Marshaling error: ", err)
	}

	// Unmarshaling
	newArticle := &article.Article{}
	err = proto.Unmarshal(data, newArticle)
	if err != nil {
		log.Fatal("Unmarshaling errror: ", err)
	}

	// Display result
	fmt.Println(data)
	fmt.Println(newArticle)
}
