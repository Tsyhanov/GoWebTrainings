package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"test-registration-form/models"
	"test-registration-form/pkg/db"
)

func GetPostsAndComments() {
	fmt.Println("concurrent:GetPostsAndComments")

	//TODO: rewrite it to use concurrency
	//get posts
	for i := 1; i <= 100; i++ {
		GetPostFromEndpoint(strconv.Itoa(i))
		//get comments for post
		GetCommentsFromEndpoint(strconv.Itoa(i))
	}

}

//get posts from external endpoint
func GetPostFromEndpoint(postid string) {
	req := "https://jsonplaceholder.typicode.com/posts/" + postid

	resp, err := http.Get(req)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll Failed: %s", err)
	}

	p := models.Post{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Printf("Post unmarshaling failed: %s", err)
		return
	}
	fmt.Println("body:", p)

	db.AddPost(p)
}

//get comments from external endpoint
func GetCommentsFromEndpoint(postid string) {
	req := "https://jsonplaceholder.typicode.com/comments?postId=" + postid

	resp, err := http.Get(req)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll Failed: %s", err)
	}

	comments := []models.Comment{}
	err = json.Unmarshal(body, &comments)
	if err != nil {
		log.Printf("Post unmarshaling failed: %s", err)
		return
	}
	fmt.Println("comments:", comments)

	//create subroutines to insert comments into db
	var wgcomments sync.WaitGroup
	for _, value := range comments {
		wgcomments.Add(1)
		go db.AddComment(value)
	}
	wgcomments.Wait()
}
