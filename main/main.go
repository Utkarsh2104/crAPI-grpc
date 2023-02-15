package main

import (
	"context"
	"fmt"
	"grpc/models"
	pb "grpc/pb/community_grpc/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:9090"
)

type server struct {
	pb.UnimplementedCommunityServiceServer
}

var mongoClient *mongo.Client

func InitializeMongo(DbUser string, DbPassword string, DbHost string, DbPort string) *mongo.Client {
	DBURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", DbUser, DbPassword, DbHost, DbPort)
	clientOptions := options.Client().ApplyURI(DBURL)

	Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		println("Mongo Initialization failed")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return Client
}

func (s *server) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	println("Client made a call to create Post ..... calling prepare")
	// models.Prepare(in)
	post := in.GetPost()

	log.Printf("post is : %v", post)

	savedPost, err := models.SavePost(mongoClient, post)
	if err != nil {
		log.Fatalf("Can not save file to mysql, %v", err)
	} else {
		log.Printf("Post saved successfully")
		print(savedPost)
	}
	println("Now its working")
	response := &pb.CreatePostResponse{
		Success: true,
	}
	return response, nil
}

func (s *server) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	println("Client made a call to UpdatePost ...... updating")
	id := in.GetId()
	println("Id to update is ", id)

	updatedPost, err := models.UpdatePost(mongoClient, in.GetUpdatedPost(), id)

	if err != nil {
		log.Fatalf("Could not update the data in DB")
	} else {
		println("Post updated succcessfully")
		print(updatedPost)
	}
	res := &pb.UpdatePostResponse{
		Success: true,
	}
	return res, nil
}

func (s *server) GetPosts(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	println("Client made a call to GetPosts() ...... getting all the posts requested")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	getPosts, err := models.GetPosts(mongoClient, ids)
	if err != nil {
		log.Fatalf("GetPosts failed!! %v", err)
	} else {
		println("GetPost successful :) ")
	}
	// res := &pb.GetPostsResponse{
	// 	Posts: getPosts.Posts[],
	// }
	return getPosts, err
}

func (s *server) DeletePosts(ctx context.Context, in *pb.DeletePostsRequest) (*pb.DeletePostsResponse, error) {
	println("Client made a call to DeletePosts() ...... getting all the deleted posts")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	DeletePosts, err := models.DeletePosts(mongoClient, ids)
	if err != nil {
		log.Fatalf("DeletePosts failed!! %v", err)
	} else {
		println("DeletePost successful :) ")
		for i := 0; i < len(DeletePosts.DeletedPosts); i++ {
			println("Deleted post id : ", DeletePosts.DeletedPosts[i].Id)
		}
	}
	return DeletePosts, err
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	println("Client made a call to create User ..... calling prepareUser")
	models.PrepareUser(in)
	user := in.GetUser()

	log.Printf("user is : %v", user)

	savedUser, err := models.SaveUser(mongoClient, user)
	if err != nil {
		// println("Error while saving user to mongodb")
		log.Fatalf("Can not save user to mysql, %v", err)
	} else {
		log.Printf("User saved successfully")
		print(savedUser)
	}
	println("Now its working")
	response := &pb.CreateUserResponse{
		Success: true,
	}
	return response, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	println("Client made a call to UpdateUser ...... updating")
	id := in.GetId()
	println("Id to update is ", id)

	updatedUser, err := models.UpdateUser(mongoClient, in.GetUpdatedUser(), id)

	if err != nil {
		log.Fatalf("Could not update the user in DB")
	} else {
		println("User updated succcessfully")
		print(updatedUser)
	}
	res := &pb.UpdateUserResponse{
		Success: true,
	}
	return res, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	println("Client made a call to GetUsers() ...... getting all the users requested")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	getUsers, err := models.GetUsers(mongoClient, ids)
	if err != nil {
		log.Fatalf("GetUsers failed!! %v", err)
	} else {
		println("GetUser successful :) ")
	}
	// res := &pb.GetPostsResponse{
	// 	Posts: getPosts.Posts[],
	// }
	return getUsers, err
}

func (s *server) DeleteUsers(ctx context.Context, in *pb.DeleteUsersRequest) (*pb.DeleteUsersResponse, error) {
	println("Client made a call to DeleteUsers() ...... getting all the deleted users")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	DeleteUsers, err := models.DeleteUsers(mongoClient, ids)
	if err != nil {
		log.Fatalf("DeleteUsers failed!! %v", err)
	} else {
		println("DeleteUser successful :) ")
		for i := 0; i < len(DeleteUsers.DeletedUsers); i++ {
			println("Deleted user id : ", DeleteUsers.DeletedUsers[i].Id)
		}
	}

	return DeleteUsers, err
}

func (s *server) CreateCoupon(ctx context.Context, in *pb.CreateCouponRequest) (*pb.CreateCouponResponse, error) {
	println("Client made a call to create Coupon ..... calling prepareCoupon")
	models.PrepareCoupon(in)
	coupon := in.GetCoupon()

	log.Printf("coupon is : %v", coupon)

	savedCoupon, err := models.SaveCoupon(mongoClient, coupon)
	if err != nil {
		// println("Error while saving user to mongodb")
		log.Fatalf("Can not save coupon to mysql, %v", err)
	} else {
		log.Printf("Coupon saved successfully")
		print(savedCoupon)
	}
	println("Now its working")
	response := &pb.CreateCouponResponse{
		Success: true,
	}
	return response, nil
}

func (s *server) UpdateCoupon(ctx context.Context, in *pb.UpdateCouponRequest) (*pb.UpdateCouponResponse, error) {
	println("Client made a call to UpdateCoupon ...... updating")
	couponcode := in.GetCouponCode()
	println("Coupon code to update is ", couponcode)

	updatedCoupon, err := models.UpdateCoupon(mongoClient, in.GetUpdatedCoupon(), couponcode)

	if err != nil {
		log.Fatalf("Could not update the coupon in DB")
	} else {
		println("Coupon updated succcessfully")
		print(updatedCoupon)
	}
	res := &pb.UpdateCouponResponse{
		Success: true,
	}
	return res, nil
}

func (s *server) GetCoupons(ctx context.Context, in *pb.GetCouponsRequest) (*pb.GetCouponsResponse, error) {
	println("Client made a call to GetCoupons() ...... getting all the coupons requested")
	couponcodes := in.CouponCodes
	println("couponcodes requested are : ")
	for i := 0; i < len(couponcodes); i++ {
		print(couponcodes[i], " ")
	}

	getCoupons, err := models.GetCoupons(mongoClient, couponcodes)
	if err != nil {
		log.Fatalf("GetCoupons failed!! %v", err)
	} else {
		println("GetCoupon successful :) ")
	}

	return getCoupons, err
}

func (s *server) DeleteCoupons(ctx context.Context, in *pb.DeleteCouponsRequest) (*pb.DeleteCouponsResponse, error) {
	println("Client made a call to DeleteCoupons() ...... getting all the deleted coupons")
	couponcodes := in.CouponCodes
	println("couponcodes requested are : ")
	for i := 0; i < len(couponcodes); i++ {
		print(couponcodes[i], " ")
	}

	DeleteCoupons, err := models.DeleteCoupons(mongoClient, couponcodes)
	if err != nil {
		log.Fatalf("DeleteCoupons failed!! %v", err)
	} else {
		println("DeleteCoupon successful :) ")
		for i := 0; i < len(DeleteCoupons.DeletedCoupons); i++ {
			println("Deleted coupon code : ", DeleteCoupons.DeletedCoupons[i].CouponCode)
		}
	}

	return DeleteCoupons, err
}

func (s *server) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	println("Client made a call to create Comment ..... calling prepare")
	models.PrepareComment(in)
	comment := in.GetComment()

	print("comment is : %v", comment)

	savedComment, err := models.SaveComment(mongoClient, comment)
	if err != nil {
		log.Fatalf("Can not save comment to mongo, %v", err)
	} else {
		log.Printf("Comment saved successfully")
		print(savedComment)
	}
	println("Now its working")
	response := &pb.CreateCommentResponse{
		Success: true,
	}
	return response, nil
}

func (s *server) UpdateComment(ctx context.Context, in *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	println("Client made a call to UpdateComment ...... updating")
	id := in.GetId()
	println("Id to update is ", id)

	updatedComment, err := models.UpdateComment(mongoClient, in.GetUpdatedComment(), id)

	if err != nil {
		log.Fatalf("Could not update the comment in DB")
	} else {
		println("Comment updated succcessfully")
		print(updatedComment)
	}
	res := &pb.UpdateCommentResponse{
		Success: true,
	}
	return res, nil
}

func (s *server) GetComments(ctx context.Context, in *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	println("Client made a call to GetComment() ...... getting all the Comment requested")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	getComment, err := models.GetComments(mongoClient, ids)
	if err != nil {
		log.Fatalf("GetComment failed!! %v", err)
	} else {
		println("GetComment successful :) ")
	}
	// res := &pb.GetPostsResponse{
	// 	Posts: getPosts.Posts[],
	// }
	return getComment, err
}

func (s *server) DeleteComments(ctx context.Context, in *pb.DeleteCommentsRequest) (*pb.DeleteCommentsResponse, error) {
	println("Client made a call to DeleteComment() ...... Deleteting all the Comment requested")
	ids := in.Ids
	println("ids requested are : ")
	for i := 0; i < len(ids); i++ {
		print(ids[i], " ")
	}

	deleteComment, err := models.DeleteComments(mongoClient, ids)
	if err != nil {
		log.Fatalf("DeleteComment failed!! %v", err)
	} else {
		println("DeleteComment successful :) ")
		for i := 0; i < len(deleteComment.DeletedComments); i++ {
			println("Deleted coupon code : ", deleteComment.DeletedComments[i].Id)
		}
	}
	// res := &pb.DeletePostsResponse{
	// 	Posts: DeletePosts.Posts[],
	// }
	return deleteComment, err
}

func main() {
	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Printf("Fatal error! Did not connect %v", err)
	}
	log.Printf("Worked! Server listening at %v", lis.Addr())

	print("Starting mongo server ... ")

	var DbUser = "root"
	var DbPassword = "root"
	var DbHost = "localhost"
	var DbPort = "27017"
	mongoClient = InitializeMongo(DbUser, DbPassword, DbHost, DbPort)
	print("Mongo Server started successfully")

	grpcServer := grpc.NewServer()

	pb.RegisterCommunityServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	print("Server registered")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Cannot start server")
	}

	log.Printf("Wohoo! sERVER started too! serving at %v", grpcServer.GetServiceInfo())

	// defer cancel()
	// r, err := c.CreatePost(ctx, &pb.CreatePostRequest{Post: &p})
	// if err != nil {
	// 	log.Printf("Error! Did not get any response")
	// }
	// log.Printf("r.id %s", r.String())
}
