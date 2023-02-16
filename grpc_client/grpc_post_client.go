package grpc_client

import (
	"context"
	"grpc/graph/model"
	"grpc/models"
	pb "grpc/pb/community_grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CreatePost(address string, p model.Post) (bool, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server, %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	created, err := client.CreatePost(ctx, &pb.CreatePostRequest{
		Post: &pb.Post{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			Author:    p.Author,
			Comments:  p.Comments,
			AuthorId:  p.AuthorID,
			CreatedAt: p.CreatedAt,
		}})

	if err != nil {
		log.Fatalf("Failed creating a post, %v", err)
	} else {
		log.Printf("Post created successfully")
		// log.Printf(created.GetSuccess())
		if created.GetSuccess() {
			log.Printf("True returned")
		} else {
			log.Printf("False returned")
		}
	}
	return created.GetSuccess(), nil
}

func UpdatePost(address string, id string, post model.PostInput) (bool, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server, %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	p := models.PrepareUpdatePost(post)

	updated, err := client.UpdatePost(ctx, &pb.UpdatePostRequest{
		Id: id,
		UpdatedPost: &pb.Post{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			Author:    p.Author,
			Comments:  p.Comments,
			AuthorId:  p.AuthorID,
			CreatedAt: p.CreatedAt,
		},
	})

	if err != nil {
		println("Updation failed")
	} else {
		println("Updation based on id successful")
		println(updated.GetSuccess())
	}
	return updated.GetSuccess(), nil
}

func GetPosts(address string, ids []string) *pb.GetPostsResponse {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server, %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	GetPosts, err := client.GetPosts(ctx, &pb.GetPostsRequest{
		Ids: ids,
	})

	if err != nil {
		log.Fatalf("GetPost failed")
		return nil
	} else {
		println("Get posts successful. Here are the posts")
		for i := 0; i < len(GetPosts.Posts); i++ {
			println("post ", i, " : \n", GetPosts.Posts[i])
		}
		return &pb.GetPostsResponse{
			Posts: GetPosts.Posts,
		}
	}
}

func DeletePost(address string, postsID []string) ([]string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server, %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunityServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("key1", "val1", "key2", "val2"),
	)

	DeletePosts, err := client.DeletePosts(ctx, &pb.DeletePostsRequest{
		Ids: postsID,
	})

	res := []string{}
	if err != nil {
		log.Fatalf("DeletePost failed")
	} else {
		println("Delete posts successful. Here are the deleted posts")
		for i := 0; i < len(DeletePosts.DeletedPosts); i++ {
			println("user ", i, " : \n", "Deleted post id : ", DeletePosts.DeletedPosts[i].Id)
			res = append(res, DeletePosts.DeletedPosts[i].Id)
		}
	}
	return res, nil
}
