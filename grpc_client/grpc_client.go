package grpc_client

import (
	"context"
	"grpc/graph/model"
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

func main() {
	// p := pb.Post{
	// 	Id:        "1",
	// 	Title:     "This is sample post",
	// 	Content:   "This sample post has this sample content! Please work!",
	// 	Author:    "Author",
	// 	Comments:  []string{"This is a samplpe comment"},
	// 	AuthorId:  "2",
	// 	CreatedAt: "12:23:23",
	// }
	print("Creating context")

	// conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Cannot connect to server, %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewCommunityServiceClient(conn)

	// ctx := context.Background()
	// ctx = metadata.NewOutgoingContext(
	// 	ctx,
	// 	metadata.Pairs("key1", "val1", "key2", "val2"),
	// )

	// p := pb.Post{}
	// created, err := client.CreatePost(ctx, &pb.CreatePostRequest{
	// 	Post: &p,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed creating a post, %v", err)
	// } else {
	// 	log.Printf("Post created successfully")
	// 	// log.Printf(created.GetSuccess())
	// 	if created.GetSuccess() {
	// 		log.Printf("True returned")
	// 	} else {
	// 		log.Printf("False returned")
	// 	}
	// }

	// updatedPost := &pb.CreatePostRequest{}
	// models.Prepare(updatedPost)

	// updated, err := client.UpdatePost(ctx, &pb.UpdatePostRequest{
	// 	Id:          "WsqZVokio63gxsmLPdSG55",
	// 	UpdatedPost: updatedPost.GetPost(),
	// })

	// if err != nil {
	// 	println("Updation failed")
	// } else {
	// 	println("Updation based on id successful")
	// 	println(updated)
	// }

	// GetPosts, err := client.GetPosts(ctx, &pb.GetPostsRequest{
	// 	Ids: []string{"hVYseaHQYqseh9ZheRLWRd", "MnwSHsuN9okaPEiSerDuA3", "oBvaQhfFqVWVdDBgAhaER6"},
	// })

	// if err != nil {
	// 	log.Fatalf("GetPost failed")
	// } else {
	// 	println("Get posts successful. Here are the posts")
	// 	for i := 0; i < len(GetPosts.Posts); i++ {
	// 		println("post ", i, " : \n", GetPosts.Posts[i])
	// 	}
	// }

	// DeletePosts, err := client.DeletePosts(ctx, &pb.DeletePostsRequest{
	// 	Ids: []string{"RzXdTZwGsf3HHLJ4QQN3N9"},
	// })

	// if err != nil {
	// 	log.Fatalf("DeletePost failed")
	// } else {
	// 	println("Delete posts successful. Here are the deleted posts")
	// 	for i := 0; i < len(DeletePosts.DeletedPosts); i++ {
	// 		println("user ", i, " : \n", "Deleted post id : ", DeletePosts.DeletedPosts[i].Id)
	// 	}
	// }

	// u := pb.User{}
	// created_user, err := client.CreateUser(ctx, &pb.CreateUserRequest{
	// 	User: &u,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed creating an user, %v", err)
	// } else {
	// 	log.Printf("User created successfully")
	// 	// log.Printf(created.GetSuccess())
	// 	if created_user.GetSuccess() {
	// 		log.Printf("True returned")
	// 	} else {
	// 		log.Printf("False returned")
	// 	}
	// }

	// updated_user, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
	// 	Id: "uLXqsAc4SZHF9zoyRGmnRY",
	// 	UpdatedUser: &pb.User{
	// 		Id:        "ajhgdvcfgacdgcaghwdfgc",
	// 		Nickname:  "Alex",
	// 		Email:     "something@something.com",
	// 		VehicleId: shortuuid.New(),
	// 		Picurl:    "gdhrsjtrdytf",
	// 		CreatedAt: time.Now().String(),
	// 	},
	// })
	// if err != nil {
	// 	println("UpdateUser failed")
	// } else {
	// 	println("Updated successfully")
	// 	print(updated_user)
	// }

	// GetUsers, err := client.GetUsers(ctx, &pb.GetUsersRequest{
	// 	Ids: []string{"ajhgdvcfgacdgcaghwdfgc", "ajhgdvcfgacdgcaghwdfgc"},
	// })

	// if err != nil {
	// 	log.Fatalf("GetUser failed")
	// } else {
	// 	println("Get users successful. Here are the users")
	// 	for i := 0; i < len(GetUsers.Users); i++ {
	// 		println("user ", i, " : \n", GetUsers.Users[i])
	// 	}
	// }

	// DeleteUsers, err := client.DeleteUsers(ctx, &pb.DeleteUsersRequest{
	// 	Ids: []string{"ajhgdvcfgacdgcaghwdfgc"},
	// })

	// if err != nil {
	// 	log.Fatalf("DeleteUser failed")
	// } else {
	// 	println("Delete users successfully. Here are the deleted users")
	// 	for i := 0; i < len(DeleteUsers.DeletedUsers); i++ {
	// 		println("user ", i, " : \n", "Deleted user id : ", DeleteUsers.DeletedUsers[i].Id)
	// 	}
	// }

	// c := pb.Coupon{}
	// created_coupon, err := client.CreateCoupon(ctx, &pb.CreateCouponRequest{
	// 	Coupon: &c,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed creating a coupon, %v", err)
	// } else {
	// 	log.Printf("Coupon created successfully")
	// 	// log.Printf(created.GetSuccess())
	// 	if created_coupon.GetSuccess() {
	// 		log.Printf("True returned")
	// 	} else {
	// 		log.Printf("False returned")
	// 	}
	// }

	// updated_coupon, err := client.UpdateCoupon(ctx, &pb.UpdateCouponRequest{
	// 	CouponCode: "ggrrpp",
	// 	UpdatedCoupon: &pb.Coupon{
	// 		CouponCode: "rtrtrtr",
	// 		Amount:     "1000",
	// 		CreatedAt:  time.Now().String(),
	// 	},
	// })
	// if err != nil {
	// 	println("UpdateCoupon failed")
	// } else {
	// 	println("Updated successfully")
	// 	print(updated_coupon)
	// }

	// GetCoupons, err := client.GetCoupons(ctx, &pb.GetCouponsRequest{
	// 	CouponCodes: []string{"rtrtrtr", "Kh9jbRkcfADThCNPLbUhLR"},
	// })

	// if err != nil {
	// 	log.Fatalf("GetCoupon failed")
	// } else {
	// 	println("Get coupons successful. Here are the coupons")
	// 	for i := 0; i < len(GetCoupons.Coupons); i++ {
	// 		println("coupon ", i, " : \n", GetCoupons.Coupons[i])
	// 	}
	// }

	// DeleteCoupons, err := client.DeleteCoupons(ctx, &pb.DeleteCouponsRequest{
	// 	CouponCodes: []string{"rtrtrtr", "Kh9jbRkcfADThCNPLbUhLR"},
	// })

	// if err != nil {
	// 	log.Fatalf("DeleteCoupon failed")
	// } else {
	// 	println("Delete coupons successfully. Here are the deleted coupons")
	// 	for i := 0; i < len(DeleteCoupons.DeletedCoupons); i++ {
	// 		println("coupon ", i, " : \n", "Deleted coupn code : ", DeleteCoupons.DeletedCoupons[i].CouponCode)
	// 	}
	// }

	// cm := pb.Comment{}
	// comment, err := client.CreateComment(ctx, &pb.CreateCommentRequest{
	// 	Comment: &cm,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed creating a comment, %v", err)
	// } else {
	// 	log.Printf("Comment created successfully")
	// 	// log.Printf(created.GetSuccess())
	// 	if comment.GetSuccess() {
	// 		log.Printf("True returned")
	// 	} else {
	// 		log.Printf("False returned")
	// 	}
	// }

	// updated_comment, err := client.UpdateComment(ctx, &pb.UpdateCommentRequest{
	// 	Id: "FMY5JoaFmc4B9G8TMLNe2i",
	// 	UpdatedComment: &pb.Comment{
	// 		Id:        "rtrhghfhtrtr",
	// 		Content:   "ashdgfahgjsvfhagsvfhjagsdvfahkscfasdvcahgd",
	// 		Author:    "Author",
	// 		CreatedAt: time.Now().String(),
	// 	},
	// })
	// if err != nil {
	// 	println("UpdateComment failed")
	// } else {
	// 	println("Updated successfully")
	// 	print(updated_comment)
	// }

	// GetComments, err := client.GetComments(ctx, &pb.GetCommentsRequest{
	// 	Ids: []string{"rtrhghfhtrtr", "TExBbivdM9YvDQHTTMs6rD"},
	// })

	// if err != nil {
	// 	log.Fatalf("GetComment failed")
	// } else {
	// 	println("Get comments successful. Here are the comments")
	// 	for i := 0; i < len(GetComments.Comments); i++ {
	// 		println("comment ", i, " : \n", GetComments.Comments[i])
	// 	}
	// }

	// DeleteComments, err := client.DeleteComments(ctx, &pb.DeleteCommentsRequest{
	// 	Ids: []string{"rtrhghfhtrtr"},
	// })

	// if err != nil {
	// 	log.Fatalf("DeleteComment failed")
	// } else {
	// 	println("Delete comments successfully. Here are the deleted comments")
	// 	for i := 0; i < len(DeleteComments.DeletedComments); i++ {
	// 		println("coupon ", i, " : \n", "Deleted comment id : ", DeleteComments.DeletedComments[i].Id)
	// 	}
	// }

}
