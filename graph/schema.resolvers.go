package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"grpc/graph/model"
	"grpc/grpc_client"
	"grpc/models"
	"log"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, id *string) (bool, error) {
	// post := model.Post{}
	post := models.Prepare(model.Post{})

	res, err := grpc_client.CreatePost(":9090", post)
	if err != nil {
		log.Fatalf("Error while creating Post by graphQL ..... %v", err)
	} else {
		println("Creating Post by GraphQL successful .. message on GraphQl server, :))")
		println("Received response, ", res)
	}
	return res, nil
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.PostInput) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - UpdatePost"))
}

// DeletePost is the resolver for the DeletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postsID []string) ([]string, error) {
	panic(fmt.Errorf("not implemented: DeletePost - DeletePost"))
}

// GetPosts is the resolver for the GetPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, ids []*string) ([]*model.Post, error) {
	lis := []string{"oBvaQhfFqVWVdDBgAhaER6", "MnwSHsuN9okaPEiSerDuA3"}
	posts := grpc_client.GetPosts(":9090", lis)

	ret := []*model.Post{}
	for i := 0; i < len(posts.Posts); i++ {
		p := model.Post{
			ID:        posts.Posts[i].Id,
			Title:     posts.Posts[i].Title,
			Content:   posts.Posts[i].Content,
			Author:    posts.Posts[i].Author,
			Comments:  posts.Posts[i].Comments,
			AuthorID:  posts.Posts[i].AuthorId,
			CreatedAt: posts.Posts[i].CreatedAt,
		}

		ret = append(ret, &p)
	}
	return ret, nil
}

// GetCoupons is the resolver for the GetCoupons field.
func (r *queryResolver) GetCoupons(ctx context.Context, codes []*string) ([]string, error) {
	panic(fmt.Errorf("not implemented: GetCoupons - GetCoupons"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }