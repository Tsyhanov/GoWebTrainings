package concurrent

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
	//get posts
	var wgposts sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wgposts.Add(1)
		go GetPostFromEndpoint(&wgposts, strconv.Itoa(i))
	}
	wgposts.Wait()
	fmt.Println("concurrent:add posts wg done")
	//get comments for post
	var wgcomments sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wgcomments.Add(1)
		go GetCommentsFromEndpoint(&wgcomments, strconv.Itoa(i))
	}
	wgcomments.Wait()
}

//get posts from external endpoint
func GetPostFromEndpoint(wg *sync.WaitGroup, postid string) {
	defer wg.Done()

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

	db.AddPost(p)
}

//get comments from external endpoint
func GetCommentsFromEndpoint(wg *sync.WaitGroup, postid string) {
	defer wg.Done()

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

	//create subroutines to insert comments into db
	var wgcomments sync.WaitGroup
	for _, value := range comments {
		wgcomments.Add(1)
		go db.AddComment(value)
	}
	wgcomments.Wait()
}
