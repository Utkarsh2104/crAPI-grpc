/*
 * Licensed under the Apache License, Version 2.0 (the “License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import (
	"context"
	"errors"
	"fmt"
	"grpc/graph/model"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "grpc/pb/community_grpc/proto"

	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func PrepareCoupon(coupon *pb.CreateCouponRequest) {
// 	coupon.Coupon = &pb.Coupon{
// 		CouponCode: shortuuid.New(),
// 		Amount:     "500",
// 		CreatedAt:  time.Now().String(),
// 	}
// }
// func PrepareNewPost(post model.Post) model.Post {
// 	post = model.Post{
// 		ID:        shortuuid.New(),
// 		Title:     "Post Title - Filled by Prepare",
// 		Content:   "Post Content {.......} - Filled by Prepare",
// 		Author:    "Author - Made by Rajat",
// 		Comments:  []string{"comment1", "comment2", "comment3"},
// 		AuthorID:  "AuthorID - id " + shortuuid.New(),
// 		CreatedAt: time.Now().String(),
// 	}
// 	return post
// }

func PrepareNewCoupon(coupon model.Coupon) model.Coupon {
	coupon = model.Coupon{
		CouponCode: shortuuid.New(),
		Amount:     strconv.Itoa(rand.Intn(2000)),
		CreatedAt:  time.Now().String(),
	}
	return coupon
}

func PrepareUpdatedCoupon(coupon model.CouponInput) model.Coupon {
	// using pointers and reference for handling a warning
	uc := &model.Coupon{
		CouponCode: coupon.CouponCode,
		Amount:     coupon.Amount,
		CreatedAt:  coupon.CreatedAt,
	}
	return *uc
}

// Validate coupon
func ValidateCoupon(coupon *pb.CreateCouponRequest) error {

	if coupon.Coupon.CouponCode == "" {
		return errors.New("required coupon code")
	}
	if coupon.Coupon.Amount == "" {
		return errors.New("required coupon amount")
	}

	return nil
}

// SaveCoupon persits data into database
func SaveCoupon(client *mongo.Client, coupon *pb.Coupon) (*pb.CreateCouponResponse, error) {

	collection := client.Database("crapi").Collection("coupon")
	_, err := collection.InsertOne(context.TODO(), coupon)
	if err != nil {
		println("Error while inserting coupon into collection")
		fmt.Println(err)
	}

	res := &pb.CreateCouponResponse{
		Success: true,
	}
	return res, err
}

// Update coupon persisting into database
func UpdateCoupon(client *mongo.Client, coupon *pb.Coupon, couponcode string) (*pb.UpdateCouponResponse, error) {
	collection := client.Database("crapi").Collection("coupon")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"couponcode", couponcode}}
	update := bson.D{{"$set", coupon}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		println("Error while updating by couponcode")
		fmt.Println(err)
	}

	res := &pb.UpdateCouponResponse{
		Success: true,
	}
	return res, nil
}

// Get an array of all coupons having matching couponcode
func GetCoupons(client *mongo.Client, in []string) (*pb.GetCouponsResponse, error) {
	collection := client.Database("crapi").Collection("coupon")
	var coupons [](*pb.Coupon)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"couponcode", in[i]}}
		var result *pb.Coupon
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatalf("Fetching documents from collection failed, %v", err)
		} else {
			coupons = append(coupons, result)
		}
	}
	res := &pb.GetCouponsResponse{
		Coupons: coupons,
	}
	return res, nil
}

// Get an array of all deleted coupons having matching couponcode
func DeleteCoupons(client *mongo.Client, in []string) (*pb.DeleteCouponsResponse, error) {
	collection := client.Database("crapi").Collection("coupon")
	var coupons [](*pb.Coupon)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"couponcode", in[i]}}
		var result *pb.Coupon
		err_get := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err_get != nil {
			println("Cannot delete coupon with id " + in[i] + " .... Coupon does not exist in database")
			continue
		}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatalf("Deleting documents from collection failed, %v", err)
		} else {
			coupons = append(coupons, result)
		}
	}
	res := &pb.DeleteCouponsResponse{
		DeletedCoupons: coupons,
	}
	return res, nil
}

// // Coupon
// type Coupon struct {
// 	CouponCode string `bson:"coupon_code" json:"coupon_code"`
// 	Amount     string `json:"amount"`
// 	CreatedAt  time.Time
// }

// func (c *Coupon) Prepare() {
// 	c.CouponCode = html.EscapeString(strings.TrimSpace(c.CouponCode))
// 	c.Amount = html.EscapeString(strings.TrimSpace(c.Amount))
// 	c.CreatedAt = time.Now()

// }

// // Validate coupon
// func (c *Coupon) Validate() error {

// 	if c.CouponCode == "" {
// 		return errors.New("required coupon code")
// 	}
// 	if c.Amount == "" {
// 		return errors.New("required coupon amount")
// 	}

// 	return nil
// }

// // SaveCoupon save coupon database.
// func SaveCoupon(client *mongo.Client, coupon Coupon) (Coupon, error) {

// 	// Get a handle for your collection
// 	collection := client.Database("crapi").Collection("coupon")

// 	// Insert a single document
// 	insertResult, err := collection.InsertOne(context.TODO(), coupon)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

// 	return coupon, err
// }

// // ValidateCode write query in mongodb for check coupon code
// func ValidateCode(client *mongo.Client, db *gorm.DB, bsonMap bson.M) (Coupon, error) {
// 	var result Coupon

// 	// Get a handle for your collection
// 	collection := client.Database("crapi").Collection("coupons")

// 	err := collection.FindOne(context.TODO(), bsonMap).Decode(&result)
// 	if err != nil {
// 		return result, err
// 	}
// 	return result, err
// }
