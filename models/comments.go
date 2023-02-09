// /*
//  * Licensed under the Apache License, Version 2.0 (the “License”);
//  * you may not use this file except in compliance with the License.
//  * You may obtain a copy of the License at
//  *
//  *         http://www.apache.org/licenses/LICENSE-2.0
//  *
//  * Unless required by applicable law or agreed to in writing, software
//  * distributed under the License is distributed on an “AS IS” BASIS,
//  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  * See the License for the specific language governing permissions and
//  * limitations under the License.
//  */

package models

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc/pb/community_grpc/proto"

	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Comments struct {
// 	ID        string `gorm:"primary_key;auto_increment" json:"id"`
// 	Content   string `gorm:"size:255;not null;" json:"content"`
// 	CreatedAt time.Time
// 	Author    Author `json:"author"`
// }

func PrepareComment(comment *pb.CreateCommentRequest) {
	comment.Comment = &pb.Comment{
		Id:        shortuuid.New(),
		Content:   "Comment {.......}",
		Author:    "Author",
		CreatedAt: time.Now().String(),
	}
}

func SaveComment(client *mongo.Client, comment *pb.Comment) (*pb.CreateCommentResponse, error) {

	collection := client.Database("crapi").Collection("comment")
	_, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		println("Error while saving comment into collection")
		fmt.Println(err)
	}

	res := &pb.CreateCommentResponse{
		Success: true,
	}
	return res, nil
}

func UpdateComment(client *mongo.Client, comment *pb.Comment, id string) (*pb.UpdateCommentResponse, error) {
	collection := client.Database("crapi").Collection("comment")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", comment}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		println("Error while updating comment by id")
		fmt.Println(err)
	}

	res := &pb.UpdateCommentResponse{
		Success: true,
	}
	return res, nil
}

func GetComments(client *mongo.Client, in []string) (*pb.GetCommentsResponse, error) {
	collection := client.Database("crapi").Collection("comment")
	var comments [](*pb.Comment)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var result *pb.Comment
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatalf("Fetching documents from collection failed, %v", err)
		} else {
			comments = append(comments, result)
		}
	}
	res := &pb.GetCommentsResponse{
		Comments: comments,
	}
	return res, nil
}

func DeleteComments(client *mongo.Client, in []string) (*pb.DeleteCommentsResponse, error) {
	collection := client.Database("crapi").Collection("comment")
	var comments [](*pb.Comment)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var comment *pb.Comment
		err_get := collection.FindOne(context.TODO(), filter).Decode(&comment)
		if err_get != nil {
			log.Fatalf("Failed to get the comment from collection")
		}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatalf("Fetching documents from collection failed, %v", err)
		} else {
			comments = append(comments, comment)
		}
	}
	res := &pb.DeleteCommentsResponse{
		DeletedComments: comments,
	}
	return res, nil
}

// // CommentOnPost Add comment in post by id.
// func CommentOnPost(client *mongo.Client, postComment Comments) (Post, error) {
// 	var comments Comments
// 	//Comments.Author = Prepare()
// 	//Take data from Database by postId
// 	preData, err := GetPostByID(client, postComment.ID)
// 	updatePost := preData
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		comments.Content = postComment.Content
// 		comments.Author = Prepare()
// 		comments.CreatedAt = time.Now()
// 		//Add comment in post
// 		updatePost.Comments = append(updatePost.Comments, comments)

// 		update := bson.D{{"$set", bson.D{{"comments", updatePost.Comments}}}}

// 		collection := client.Database("crapi").Collection("post")

// 		_, err = collection.UpdateOne(context.TODO(), preData, update)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	return updatePost, err
// }
