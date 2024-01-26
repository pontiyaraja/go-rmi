package server

import (
	"context"
	"grpc_blog/model/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBlog(t *testing.T) {
	blgServer := newServer()
	assert.Empty(t, blgServer.blogs)
	resp, err := blgServer.Create(context.Background(), &proto.Blog{PostID: 12, Content: "test Content", Author: "I'm",
		PublicationDate: "23-01-24", Tags: "my tag"})
	assert.Nil(t, err)
	assert.Equal(t, int32(12), resp.PostID)
	resp, err = blgServer.Create(context.Background(), nil)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func TestReadBlog(t *testing.T) {
	blgServer := newServer()
	assert.Empty(t, blgServer.blogs)
	resp, err := blgServer.Create(context.Background(), &proto.Blog{PostID: 12, Content: "test Content", Author: "I'm",
		PublicationDate: "23-01-24", Tags: "my tag"})
	assert.Nil(t, err)
	resp, err = blgServer.Read(context.Background(), &proto.PostID{PostID: 12})
	assert.Nil(t, err)
	assert.Equal(t, int32(12), resp.PostID)
	assert.Equal(t, "23-01-24", resp.PublicationDate)
	resp, err = blgServer.Read(context.Background(), nil)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func TestUpdateBlog(t *testing.T) {
	blgServer := newServer()
	assert.Empty(t, blgServer.blogs)
	resp, err := blgServer.Create(context.Background(), &proto.Blog{PostID: 12, Content: "test Content", Author: "I'm",
		PublicationDate: "23-01-24", Tags: "my tag"})
	assert.Nil(t, err)
	resp, err = blgServer.Read(context.Background(), &proto.PostID{PostID: 12})
	assert.Nil(t, err)
	assert.Equal(t, int32(12), resp.PostID)
	assert.Equal(t, "23-01-24", resp.PublicationDate)
	resp, err = blgServer.UPdate(context.Background(), &proto.Blog{PostID: 12, Content: "test Content", Author: "I'm",
		PublicationDate: "25-01-24", Tags: "my tag"})
	assert.Nil(t, err)
	assert.Equal(t, int32(12), resp.PostID)
	assert.Equal(t, "25-01-24", resp.PublicationDate)
	resp, err = blgServer.UPdate(context.Background(), &proto.Blog{PostID: 15, Content: "test Content", Author: "I'm",
		PublicationDate: "25-01-24", Tags: "my tag"})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestDeleteBlog(t *testing.T) {
	blgServer := newServer()
	assert.Empty(t, blgServer.blogs)
	resp, err := blgServer.Create(context.Background(), &proto.Blog{PostID: 12, Content: "test Content", Author: "I'm",
		PublicationDate: "23-01-24", Tags: "my tag"})
	assert.Nil(t, err)
	resp, err = blgServer.Read(context.Background(), &proto.PostID{PostID: 12})
	assert.Nil(t, err)
	assert.Equal(t, int32(12), resp.PostID)
	assert.Equal(t, "23-01-24", resp.PublicationDate)
	resp, err = blgServer.Delete(context.Background(), &proto.PostID{PostID: 12})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(12), resp.PostID)
	assert.Equal(t, "23-01-24", resp.PublicationDate)
	resp, err = blgServer.Delete(context.Background(), &proto.PostID{PostID: 12})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}
