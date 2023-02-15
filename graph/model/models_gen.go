// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Author    string `json:"author"`
}

type CommentInput struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Author    string `json:"author"`
}

type Coupon struct {
	CouponCode string `json:"coupon_code"`
	Amount     string `json:"amount"`
	CreatedAt  string `json:"created_at"`
}

type CouponInput struct {
	CouponCode string `json:"coupon_code"`
	Amount     string `json:"amount"`
	CreatedAt  string `json:"created_at"`
}

type Post struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	Comments  []string `json:"comments"`
	AuthorID  string   `json:"author_id"`
	CreatedAt string   `json:"created_at"`
}

type PostInput struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	Comments  []string `json:"comments"`
	AuthorID  string   `json:"author_id"`
	CreatedAt string   `json:"created_at"`
}

type User struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	VehicleID string `json:"vehicle_id"`
	Picurl    string `json:"picurl"`
	CreatedAt string `json:"created_at"`
}

type UserInput struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	VehicleID string `json:"vehicle_id"`
	Picurl    string `json:"picurl"`
	CreatedAt string `json:"created_at"`
}