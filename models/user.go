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
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"

	pb "grpc/pb/community_grpc/proto"

	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userID uint64
var nickname string
var userEmail string

var picurl string
var vehicleID string

// Author model
// type Author struct {
// 	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
// 	Email     string    `gorm:"size:100;not null;unique" json:"email"`
// 	VehicleID string    `gorm:"size:100;not null;unique" json:"vehicleid"`
// 	Picurl    string    `gorm:"size:30000;not null;unique" json:"profile_pic_url"`
// 	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// }

// Hash for password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare password and hashcode
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//Prepare initilize user object

func PrepareUser(user *pb.CreateUserRequest) {
	user.User = &pb.User{
		Id:        shortuuid.New(),
		Nickname:  "John",
		Email:     "something@something.com",
		VehicleId: shortuuid.New(),
		Picurl:    "gdhrsjtrdytf",
		CreatedAt: time.Now().String(),
	}
}

// Validate Author
func ValidateUser(user *pb.CreateUserRequest, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.User.Nickname == "" {
			return errors.New("required nickname")
		}
		if user.User.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.User.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	case "login":
		if user.User.Nickname == "" {
			return errors.New("required nickname")
		}
		if user.User.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.User.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	default:
		if user.User.Nickname == "" {
			return errors.New("required nickname")
		}
		if user.User.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(user.User.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

// SaveUser persits data into database
func SaveUser(client *mongo.Client, user *pb.User) (*pb.CreateUserResponse, error) {

	collection := client.Database("crapi").Collection("user")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		println("Error while inserting user into collection")
		fmt.Println(err)
	}

	res := &pb.CreateUserResponse{
		Success: true,
	}
	return res, err
}

// Update user persisting into database
func UpdateUser(client *mongo.Client, user *pb.User, id string) (*pb.UpdateUserResponse, error) {
	collection := client.Database("crapi").Collection("user")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", user}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		println("Error while updating by id")
		fmt.Println(err)
	}

	res := &pb.UpdateUserResponse{
		Success: true,
	}
	return res, nil
}

// Get an array of all users having matching id
func GetUsers(client *mongo.Client, in []string) (*pb.GetUsersResponse, error) {
	collection := client.Database("crapi").Collection("user")
	var users [](*pb.User)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var result *pb.User
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatalf("Fetching documents from collection failed, %v", err)
		} else {
			users = append(users, result)
		}
	}
	res := &pb.GetUsersResponse{
		Users: users,
	}
	return res, nil
}

// Get an array of all deleted users having matching id
func DeleteUsers(client *mongo.Client, in []string) (*pb.DeleteUsersResponse, error) {
	collection := client.Database("crapi").Collection("user")
	var users [](*pb.User)
	for i := 0; i < len(in); i++ {
		filter := bson.D{{"id", in[i]}}
		var result *pb.User
		err_get := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err_get != nil {
			log.Fatalf(("Getting user by id failed"))
		}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatalf("Deleting documents from collection failed, %v", err)
		} else {
			users = append(users, result)
		}
	}
	res := &pb.DeleteUsersResponse{
		DeletedUsers: users,
	}
	return res, nil
}

// FindAuthorByEmail check user in database
// func FindAuthorByEmail(email string, db *gorm.DB) (*uint64, error) {
// 	var err error
// 	var id uint64
// 	var number *uint64
// 	var name string
// 	var picture []byte
// 	var uuid string
// 	userEmail = email

// 	//fetch id and number from for token user
// 	row := db.Table("user_login").Where("email LIKE ?", email).Select("id,number").Row()

// 	row.Scan(&id, &number)

// 	autherID = id
// 	//fetch name and picture from for token user
// 	row1 := db.Table("user_details").Where("user_id = ?", id).Select("name, lo_get(picture)").Row()
// 	row1.Scan(&name, &picture)
// 	if len(picture) > 0 {
// 		picurl = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(picture)
// 	}
// 	nickname = name
// 	row2 := db.Table("vehicle_details").Where("owner_id = ?", id).Select("uuid").Row()
// 	row2.Scan(&uuid)
// 	vehicleID = uuid
// 	return number, err
// }
