package server

import (
	pb "github.com/rafaelmbcosta/golang-blog-experiment/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorId string             `json:"author_id" bson:"author_id"`
	Title    string             `json:"title" bson:"title"`
	Content  string             `json:"content" bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
