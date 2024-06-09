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
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type GetDicomDocumentResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type ListDicomDocumentsResponse struct {
	Documents []GetDicomDocumentResponse `json:"documents"`
}

type GetDocumentDataResponse struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type GetDicomDocumentDataByIDandTagRequest struct {
	Group   uint16 `json:"group"`
	Element uint16 `json:"element"`
}

type GetDicomDocumentDataByIDandTagResponse struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func NewClient() IClient {
	return &Client{}
}
