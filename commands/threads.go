package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"caleox-spaceforum/models"
	"caleox-spaceforum/utils"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

const PageSize = 5

// Helper to safely get posts slice
func getPosts(thread map[string]interface{}) []interface{} {
	postsRaw, ok := thread["posts"]
	if !ok || postsRaw == nil {
		return []interface{}{}
	}
	return postsRaw.([]interface{})
}

// List threads
func ListThreads() {
	data, _ := utils.GetBin()
	threadsRaw, ok := data["threads"]
	if !ok || threadsRaw == nil {
		color.Cyan("No threads yet. Use 'post' to create one.")
		return
	}
	threads := threadsRaw.([]interface{})

	color.Cyan("Threads:")
	for i, t := range threads {
		thread := t.(map[string]interface{})
		color.HiCyan("%d. %s (by %s)", i+1, thread["title"], thread["author"])
	}
}

// Create thread
func CreateThread() {
	session, err := utils.LoadSession()
	if err != nil {
		color.Red("You need to login first!")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Thread title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	data, _ := utils.GetBin()
	threadsRaw, ok := data["threads"]
	var threads []interface{}
	if ok && threadsRaw != nil {
		threads = threadsRaw.([]interface{})
	}

	newThread := models.Thread{
		ID:      uuid.New().String(),
		Title:   title,
		Author:  session.Username,
		Created: time.Now().Format("2006-01-02"),
		Posts:   []models.Post{},
	}

	threads = append(threads, newThread)
	data["threads"] = threads
	utils.UpdateBin(data)
	color.Green("Thread created successfully!")
}

// View thread with pagination
func ViewThread(threadIndex string, page int) {
	data, _ := utils.GetBin()
	threadsRaw, ok := data["threads"]
	if !ok || threadsRaw == nil {
		color.Red("No threads found.")
		return
	}
	threads := threadsRaw.([]interface{})
	tIdx := atoi(threadIndex) - 1
	if tIdx < 0 || tIdx >= len(threads) {
		color.Red("Invalid thread number")
		return
	}

	thread := threads[tIdx].(map[string]interface{})
	color.Cyan("Title: %s | Author: %s | Created: %s\n", thread["title"], thread["author"], thread["created"])

	posts := getPosts(thread)
	total := len(posts)
	if total == 0 {
		color.Yellow("No posts yet. Be the first to reply!")
		return
	}

	start := (page - 1) * PageSize
	end := start + PageSize
	if start >= total {
		color.Red("No posts on this page.")
		return
	}
	if end > total {
		end = total
	}

	color.Cyan("Posts %d-%d of %d:", start+1, end, total)
	for i := start; i < end; i++ {
		post := posts[i].(map[string]interface{})
		color.HiWhite("%d. [%s] %s", i+1, post["author"], post["content"])
	}
}

// Reply to thread
func ReplyThread(threadIndex string) {
	session, err := utils.LoadSession()
	if err != nil {
		color.Red("You need to login first!")
		return
	}

	data, _ := utils.GetBin()
	threadsRaw, ok := data["threads"]
	if !ok || threadsRaw == nil {
		color.Red("No threads found.")
		return
	}
	threads := threadsRaw.([]interface{})
	tIdx := atoi(threadIndex) - 1
	if tIdx < 0 || tIdx >= len(threads) {
		color.Red("Invalid thread number")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your reply: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	newPost := models.Post{
		ID:      uuid.New().String(),
		Author:  session.Username,
		Content: content,
		Created: time.Now().Format("2006-01-02"),
	}

	thread := threads[tIdx].(map[string]interface{})
	posts := getPosts(thread)
	posts = append(posts, newPost)
	thread["posts"] = posts
	threads[tIdx] = thread
	data["threads"] = threads
	utils.UpdateBin(data)
	color.Green("Reply posted successfully!")
}

// Delete post
func DeletePost(threadIndex, postIndex string) {
	session, err := utils.LoadSession()
	if err != nil {
		color.Red("Login first!")
		return
	}

	data, _ := utils.GetBin()
	threadsRaw, ok := data["threads"]
	if !ok || threadsRaw == nil {
		color.Red("No threads found.")
		return
	}
	threads := threadsRaw.([]interface{})
	tIdx := atoi(threadIndex) - 1
	pIdx := atoi(postIndex) - 1
	if tIdx < 0 || tIdx >= len(threads) || pIdx < 0 {
		color.Red("Invalid numbers")
		return
	}

	thread := threads[tIdx].(map[string]interface{})
	posts := getPosts(thread)
	if pIdx >= len(posts) {
		color.Red("Invalid post number")
		return
	}

	post := posts[pIdx].(map[string]interface{})
	if post["author"] != session.Username {
		color.Red("You can only delete your own posts!")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure? (y/n): ")
	ans, _ := reader.ReadString('\n')
	ans = strings.TrimSpace(ans)
	if ans != "y" {
		color.Yellow("Cancelled.")
		return
	}

	posts = append(posts[:pIdx], posts[pIdx+1:]...)
	thread["posts"] = posts
	threads[tIdx] = thread
	data["threads"] = threads
	utils.UpdateBin(data)
	color.Green("Post deleted successfully!")
}

// Delete account
func DeleteAccount() {
	session, err := utils.LoadSession()
	if err != nil {
		color.Red("Login first!")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure you want to delete your account? (y/n): ")
	ans, _ := reader.ReadString('\n')
	ans = strings.TrimSpace(ans)
	if ans != "y" {
		color.Yellow("Cancelled.")
		return
	}

	data, _ := utils.GetBin()
	usersRaw, ok := data["users"]
	if !ok || usersRaw == nil {
		color.Red("No users found.")
		return
	}
	users := usersRaw.([]interface{})

	threadsRaw, ok := data["threads"]
	var threads []interface{}
	if ok && threadsRaw != nil {
		threads = threadsRaw.([]interface{})
	}

	newUsers := []interface{}{}
	for _, u := range users {
		user := u.(map[string]interface{})
		if user["username"] != session.Username {
			newUsers = append(newUsers, user)
		}
	}
	data["users"] = newUsers

	for i, t := range threads {
		thread := t.(map[string]interface{})
		posts := getPosts(thread)
		newPosts := []interface{}{}
		for _, p := range posts {
			post := p.(map[string]interface{})
			if post["author"] != session.Username {
				newPosts = append(newPosts, post)
			}
		}
		thread["posts"] = newPosts
		threads[i] = thread
	}
	data["threads"] = threads

	utils.UpdateBin(data)
	utils.ClearSession()
	color.Green("Account and all your posts deleted successfully!")
}

// Helper: convert string to int
func atoi(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
