package main

import (
	"time"
)

// ImageInfo holds metadata for an uploaded image
type ImageInfo struct {
	ID          int       `json:"id"`
	Filename    string    `json:"filename"`
	Size        int64     `json:"size"`
	UploadAt    time.Time `json:"upload_at"`
	StoragePath string    `json:"storage_path"`
}

// NewImageInfo creates a new ImageInfo object
func NewImageInfo(filename string, size int64, path string) *ImageInfo {
	return &ImageInfo{
		Filename:    filename,
		Size:        size,
		UploadAt:    time.Now(),
		StoragePath: path,
	}
}
