package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "grpc/pb/community_grpc/proto"

	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post Field
// type Post struct {
// 	ID        string     `gorm:"primary_key;auto_increment" json:"id"`
// 	Title     string     `gorm:"size:255;not null;unique" json:"title"`
// 	Content   string     `gorm:"size:255;not null;" json:"content"`
// 	Author    Author     `json:"author"`
// 	Comments  []Comments `json:"comments"`
// 	AuthorID  uint64     `sql:"type:int REFERENCES users(id)" json:"authorid"`
// 	CreatedAt time.Time
// }

// func (post *Post) Prepare() {
// 	post.ID = shortuuid.New()
// 	post.Title = html.EscapeString(strings.TrimSpace(post.Title))
// 	post.Content = html.EscapeString(strings.TrimSpace(post.Content))
// 	post.Author = Prepare()
// 	post.AuthorID = autherID
// 	post.Comments = []Comments{}
// 	post.CreatedAt = time.Now()
// }

func Prepare(post *pb.CreatePostRequest) {
	post.Post = &pb.Post{
		Id:        shortuuid.New(),
		Title:     "Post Title",
		Content:   "Post Content {.......}",
		Author:    "Author",
		Comments:  []string{"Coment1", "Comment2", "Comment3"},
		AuthorId:  "AuthorID",
		CreatedAt: time.Now().String(),
	}
}

func Validate(post *pb.CreatePostRequest) error {

	if post.Post.Title == "" {
		return errors.New("required title")
	}
	if post.Post.Content == "" {
		return errors.New("required content")
	}
	if len(post.Post.AuthorId) < 1 {
		return errors.New("required author")
	}
	return nil
}

// // Prepare initialize Field
// func PrepareAuth() Author {
// 	var u Author
// 	u.Nickname = nickname
// 	u.Email = userEmail
// 	u.VehicleID = vehicleID
// 	u.CreatedAt = time.Now()
// 	u.Picurl = picurl
// 	return u
// }

// SavePost persits data into database
func SavePost(client *mongo.Client, post *pb.Post) (*pb.CreatePostResponse, error) {

	collection := client.Database("crapi").Collection("post")
	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		println("Error while inserting post into collection")
		fmt.Println(err)
	}

	res := &pb.CreatePostResponse{
		Success: true,
	}
	return res, nil
}

// Update posts persisting into database
func UpdatePost(client *mongo.Client, post *pb.Post, id string) (*pb.UpdatePostResponse, error) {
	collection := client.Database("crapi").Collection("post")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", post}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		println("Error while updating by id")
		fmt.Println(err)
	}

	res := &pb.UpdatePostResponse{
		Success: true,
	}
	return res, nil
}

// Get an array of all posts having matching id
func GetPosts(client *mongo.Client, in []string) (*pb.GetPostsResponse, error) {
	collection := client.Database("crapi").Collection("post")
	var posts [](*pb.Post)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var result *pb.Post
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatalf("Fetching documents from collection failed, %v", err)
		} else {
			posts = append(posts, result)
		}
	}
	res := &pb.GetPostsResponse{
		Posts: posts,
	}
	return res, nil
}

// Get an array of all deleted posts having matching id
func DeletePosts(client *mongo.Client, in []string) (*pb.DeletePostsResponse, error) {
	collection := client.Database("crapi").Collection("post")
	var posts [](*pb.Post)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var result *pb.Post
		err_get := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err_get != nil {
			log.Fatalf(("Getting post by id failed"))
		}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatalf("Deleting documents from collection failed, %v", err)
		} else {
			posts = append(posts, result)
		}
	}
	res := &pb.DeletePostsResponse{
		DeletedPosts: posts,
	}
	return res, nil
}

// // GetPostByID fetch post by postId
// func GetPostByID(client *mongo.Client, ID string) (Post, error) {
// 	var post Post

// 	//filter := bson.D{{"name", "Ash"}}
// 	collection := client.Database("crapi").Collection("post")
// 	filter := bson.D{{"id", ID}}
// 	err := collection.FindOne(context.TODO(), filter).Decode(&post)

// 	return post, err

// }

// // FindAllPost return all recent post
// func FindAllPost(client *mongo.Client) ([]interface{}, error) {
// 	post := []Post{}

// 	options := options.Find()
// 	options.SetSort(bson.D{{"_id", -1}})
// 	options.SetLimit(10)
// 	collection := client.Database("crapi").Collection("post")
// 	cur, err := collection.Find(context.Background(), bson.D{}, options)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(cur)
// 	objectType := reflect.TypeOf(post).Elem()
// 	var list = make([]interface{}, 0)
// 	defer cur.Close(context.Background())
// 	for cur.Next(context.Background()) {
// 		result := reflect.New(objectType).Interface()
// 		err := cur.Decode(result)

// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}

// 		list = append(list, result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil, err
// 	}

// 	return list, err
// }
