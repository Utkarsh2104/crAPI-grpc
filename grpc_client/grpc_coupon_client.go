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

func CreateCoupon(address string, cp model.Coupon) (bool, error) {
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

	created, err := client.CreateCoupon(ctx, &pb.CreateCouponRequest{
		Coupon: &pb.Coupon{
			CouponCode: cp.CouponCode,
			Amount:     cp.Amount,
			CreatedAt:  cp.CreatedAt,
		}})

	if err != nil {
		log.Fatalf("Failed creating a Coupon, %v", err)
	} else {
		log.Printf("Coupon created successfully")
		// log.Printf(created.GetSuccess())
		if created.GetSuccess() {
			log.Printf("True returned")
		} else {
			log.Printf("False returned")
		}
	}
	return created.GetSuccess(), nil
}

func UpdateCoupon(address string, id string, coupon model.CouponInput) (bool, error) {
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

	uc := models.PrepareUpdatedCoupon(coupon)

	updated, err := client.UpdateCoupon(ctx, &pb.UpdateCouponRequest{
		CouponCode: id,
		UpdatedCoupon: &pb.Coupon{
			CouponCode: uc.CouponCode,
			Amount:     uc.Amount,
			CreatedAt:  uc.CreatedAt,
		},
	})

	if err != nil {
		println("Coupon Updation Failed failed")
	} else {
		println("Coupon Updation based on CouponId successful")
		println(updated.GetSuccess())
	}
	return updated.GetSuccess(), nil
}

func GetCoupon(address string, ids []string) *pb.GetCouponsResponse {
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

	GetCoupon, err := client.GetCoupons(ctx, &pb.GetCouponsRequest{
		CouponCodes: ids,
	})

	if err != nil {
		log.Fatalf("Get Coupons failed")
		return nil
	} else {
		println("Get Coupons successful. Here are the Coupons")
		for i := 0; i < len(GetCoupon.Coupons); i++ {
			println("post ", i, " : \n", GetCoupon.Coupons[i])
		}
		return &pb.GetCouponsResponse{
			Coupons: GetCoupon.Coupons,
		}
	}
}

func DeleteCoupon(address string, CouponCode []string) ([]string, error) {
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

	DeleteCoupon, err := client.DeleteCoupons(ctx, &pb.DeleteCouponsRequest{
		CouponCodes: CouponCode,
	})

	res := []string{}
	if err != nil {
		log.Fatalf("DeleteCoupon failed")
	} else {
		println("Delete Coupons successful. Here are the deleted Coupons")
		for i := 0; i < len(DeleteCoupon.DeletedCoupons); i++ {
			println("user ", i, " : \n", "Deleted post id : ", DeleteCoupon.DeletedCoupons[i].CouponCode)
			res = append(res, DeleteCoupon.DeletedCoupons[i].CouponCode)
		}
	}
	return res, nil
}
