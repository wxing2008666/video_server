/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:40
 *  @Description:
 */
package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data modal
type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id      string `json:"id"`
	AideoId string `json:"aideo_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
