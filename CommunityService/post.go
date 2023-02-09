package CommunityService

import (
	"context"
	"errors"
	"grpc/config"
	"grpc/models"
	pb "grpc/pb/community_grpc/proto"
	"log"

	"google.golang.org/grpc/metadata"
)

type Server config.Server

func (s *Server) CreatePost(ctx context.Context, post *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	body, err := metadata.FromIncomingContext(ctx)
	if !err {
		log.Printf("Failed to read context")
	} else {
		log.Printf("Body of context : ", body)
	}

	models.Prepare(post)
	postSaved, er := models.SavePost(s.Client, post)
	if er != nil {
		log.Fatalf("Post not saved, %v", err)
		return nil, errors.New("Can not save post")
	}
	return postSaved, nil
}
