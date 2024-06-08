package client

import (
	"mime/multipart"
)

type IClient interface {
}

type Client struct {
}

type CreateDicomDocumentRequest struct {
	File multipart.FileHeader `json:"file"`
}

type CreateDicomDocumentResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func NewClient() IClient {
	return &Client{}
}
