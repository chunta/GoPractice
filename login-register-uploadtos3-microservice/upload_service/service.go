package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

const (
	// Cache-related constants
	cacheExpiration    = 3 * time.Minute // Cache duration for signed URLs
	cacheControlHeader = "public, max-age=180"
	signedUrlValidity  = 15 * time.Minute // Signed URL validity duration
	awsRegion          = "us-east-1"      // Change to your AWS region
	awsAccessKey       = "xxx"
	awsSecretKey       = "xxx"
	bucketName         = "paid-images"
)

var (
	// Create a global cache instance
	signedUrlCache = cache.New(cacheExpiration, 5*time.Minute)
)

func GetUserImages(c *gin.Context) {
	// Retrieve userId from the JWT context
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// Ensure the userId is an integer
	userIdInt, ok := userId.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is not an integer"})
		return
	}

	// Print the userId for debugging (optional)
	fmt.Println("User ID from token:", userIdInt)

	// Create an S3 session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		fmt.Println("Unable to create AWS session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create S3 client
	svc := s3.New(sess)

	// Prefix to filter images for the specific user
	prefix := fmt.Sprintf("uploads/%d/", userIdInt)

	// List objects in S3 with the userId prefix
	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		fmt.Println("Failed to list objects:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list objects"})
		return
	}

	// Generate signed URLs for the user
	var signedUrls []string
	for _, item := range result.Contents {
		// Cache key: userId + object key
		cacheKey := fmt.Sprintf("%d-%s", userIdInt, *item.Key)

		// Check if we already have the signed URL cached
		if cachedUrl, found := signedUrlCache.Get(cacheKey); found {
			// If cached, use the cached URL
			fmt.Println("Using cached URL for:", *item.Key)
			signedUrls = append(signedUrls, cachedUrl.(string))
			continue
		}

		// Generate a signed URL for the image
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    item.Key,
		})
		signedUrl, err := req.Presign(signedUrlValidity) // URL is valid for 15 minutes
		if err != nil {
			fmt.Println("Failed to sign URL:", err)
			continue
		}

		// Cache the signed URL for this specific image
		signedUrlCache.Set(cacheKey, signedUrl, cacheExpiration)

		// Print log indicating new URL is being generated
		fmt.Println("Generated new signed URL for:", *item.Key)

		// Append the signed URL to the list
		signedUrls = append(signedUrls, signedUrl)
	}

	// Set HTTP Cache-Control headers
	c.Header("Cache-Control", cacheControlHeader) // Cache for 3 minutes
	c.Header("Expires", time.Now().Add(cacheExpiration).Format(http.TimeFormat))

	// Respond with the signed URLs
	c.JSON(http.StatusOK, gin.H{"userId": userIdInt, "images": signedUrls})
}

func UploadImage(c *gin.Context) {
	// Retrieve userId from the context
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// Type assert userId to int
	userIdInt, ok := userId.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is not an integer"})
		return
	}

	// Retrieve file from the form-data
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	fileName := header.Filename

	// Generate a unique key for the file
	key := fmt.Sprintf("uploads/%d/%s", userIdInt, fileName)

	// Create an S3 session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		log.Println("Unable to create AWS session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Create an S3 client
	svc := s3.New(sess)

	// Create tags and encode them
	tags := url.QueryEscape(fmt.Sprintf("UserId=%d", userId)) // Properly encode tags

	// Upload the file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(header.Header.Get("Content-Type")),
		Tagging:     aws.String(tags), // Add tags in URL-encoded format
	})
	if err != nil {
		log.Println("Failed to upload file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "key": key})
}

func ListS3Buckets(c *gin.Context) {
	// Create a session using static credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		log.Fatal("Unable to create AWS session: ", err)
	}

	// Create an S3 client
	svc := s3.New(sess)

	// List S3 buckets
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Unable to list buckets: ", err)
	}

	// Output all S3 buckets
	buckets := []string{}
	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("- %s\n", *bucket.Name)
		buckets = append(buckets, *bucket.Name) // Fix here: Store the result of append back into the buckets slice
	}

	// Send the list of buckets in the response
	c.JSON(200, gin.H{"buckets": buckets})
}
