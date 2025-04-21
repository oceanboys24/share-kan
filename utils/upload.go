package utils

import (
	"context"
	"mime/multipart"
	"net/url"
	"share-kan/config"
	"time"

	"github.com/minio/minio-go/v7"
)



func UploadFile(bucketName, objectName string, file multipart.File, fileSize int64, contentType string)  (*url.URL,error) {
	_ , err := config.S3Client.PutObject(context.Background(),bucketName,objectName,file,fileSize,minio.PutObjectOptions{ContentType: contentType})
	
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", `attachment; filename="`+objectName+`"`)


	presignedUrl, err := config.S3Client.PresignedGetObject(
		context.Background(),
		bucketName,
		objectName,
		time.Hour,
		reqParams,
	)
	if err != nil {
		return nil , err
	}
	
	return presignedUrl,nil
}