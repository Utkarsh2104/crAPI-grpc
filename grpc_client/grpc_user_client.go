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

func CreateUser(address string, u model.User) (bool, error) {
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

	created, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Id:        u.ID,
			Nickname:  u.Nickname,
			Email:     u.Email,
			VehicleId: u.VehicleID,
			Picurl:    u.Picurl,
			CreatedAt: u.CreatedAt,
		}})

	if err != nil {
		log.Fatalf("Failed creating a user, %v", err)
	} else {
		log.Printf("User created successfully")
		// log.Printf(created.GetSuccess())
		if created.GetSuccess() {
			log.Printf("True returned")
		} else {
			log.Printf("False returned")
		}
	}
	return created.GetSuccess(), nil
}

func GetUsers(address string, ids []string) *pb.GetUsersResponse {
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

	GetUsers, err := client.GetUsers(ctx, &pb.GetUsersRequest{
		Ids: ids,
	})

	if err != nil {
		log.Fatalf("GetUser failed")
		return nil
	} else {
		println("Get users successful. Here are the users")
		for i := 0; i < len(GetUsers.Users); i++ {
			println("user ", i, " : \n", GetUsers.Users[i])
		}
		return &pb.GetUsersResponse{
			Users: GetUsers.Users,
		}
	}
}

func UpdateUser(address string, id string, user model.UserInput) (bool, error) {
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

	u := models.PrepareUpdateUser(user)

	updated, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id: id,
		UpdatedUser: &pb.User{
			Id:        u.ID,
			Nickname:  u.Nickname,
			Email:     u.Email,
			VehicleId: u.VehicleID,
			Picurl:    u.Picurl,
			CreatedAt: u.CreatedAt,
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

func DeleteUser(address string, usersID []string) ([]string, error) {
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

	DeleteUsers, err := client.DeleteUsers(ctx, &pb.DeleteUsersRequest{
		Ids: usersID,
	})

	res := []string{}
	if err != nil {
		log.Fatalf("DeleteUser failed")
	} else {
		println("Delete users successful. Here are the deleted users")
		for i := 0; i < len(DeleteUsers.DeletedUsers); i++ {
			println("user ", i, " : \n", "Deleted user id : ", DeleteUsers.DeletedUsers[i].Id)
			res = append(res, DeleteUsers.DeletedUsers[i].Id)
		}
	}
	return res, nil
}
