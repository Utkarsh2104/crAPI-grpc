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

func CreateComment(address string, c model.Comment) (bool, error) {
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

	created, err := client.CreateComment(ctx, &pb.CreateCommentRequest{
		Comment: &pb.Comment{
			Id:        c.ID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			Author:    c.Author,
		}})

	if err != nil {
		log.Fatalf("Failed creating a comment, %v", err)
	} else {
		log.Printf("Comment created successfully")
		// log.Printf(created.GetSuccess())
		if created.GetSuccess() {
			log.Printf("True returned")
		} else {
			log.Printf("False returned")
		}
	}
	return created.GetSuccess(), nil
}

func GetComments(address string, ids []string) *pb.GetCommentsResponse {
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

	GetComments, err := client.GetComments(ctx, &pb.GetCommentsRequest{
		Ids: ids,
	})

	if err != nil {
		log.Fatalf("GetComment failed")
		return nil
	} else {
		println("Get comments successful. Here are the comments")
		for i := 0; i < len(GetComments.Comments); i++ {
			println("comment ", i, " : \n", GetComments.Comments[i])
		}
		return &pb.GetCommentsResponse{
			Comments: GetComments.Comments,
		}
	}
}

func UpdateComment(address string, id string, comment model.CommentInput) (bool, error) {
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

	c := models.PrepareUpdateComment(comment)

	updated, err := client.UpdateComment(ctx, &pb.UpdateCommentRequest{
		Id: id,
		UpdatedComment: &pb.Comment{
			Id:        c.ID,
			Content:   c.Content,
			Author:    c.Author,
			CreatedAt: c.CreatedAt,
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

func DeleteComment(address string, commentsID []string) ([]string, error) {
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

	DeleteComments, err := client.DeleteComments(ctx, &pb.DeleteCommentsRequest{
		Ids: commentsID,
	})

	res := []string{}
	if err != nil {
		log.Fatalf("DeleteUser failed")
	} else {
		println("Delete users successful. Here are the deleted users")
		for i := 0; i < len(DeleteComments.DeletedComments); i++ {
			println("comment ", i, " : \n", "Deleted comment id : ", DeleteComments.DeletedComments[i].Id)
			res = append(res, DeleteComments.DeletedComments[i].Id)
		}
	}
	return res, nil
}
