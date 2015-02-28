package store

import (
"fmt"
"github.com/mitchellh/goamz/aws"
"github.com/mitchellh/goamz/s3"
"io"
"io/ioutil"
"mime"
"os"
)

func CreateBucket() *s3.Bucket {
  auth, err := aws.EnvAuth()
  if err != nil {
    println("store -- error connecting to AWS\n")
    panic(err)
  }

  fmt.Printf("store -- connected to AWS\n")

  awsRegion := os.Getenv("STORE_AWSREGION")
  if awsRegion == "" {
    println("store -- error: environment variable STORE_AWSREGION not set")
    panic("store -- STORE_AWSREGION not set")
  }

  region := aws.Regions[awsRegion]
  if region.Name == "" {
   panic("store -- region '" + awsRegion + "' is invalid") 
  }
  println("store -- AWS region: " + region.Name)

  S3 := s3.New(auth, aws.Regions["eu-west-1"])
  
  bucketName := os.Getenv("STORE_BUCKETNAME")
  if bucketName == "" {
    println("store -- error: environment variable STORE_BUCKETNAME not set")
    panic("store -- STORE_BUCKETNAME not set")
  }
  
  bucket := S3.Bucket(bucketName)
  println("store -- Bucket name: " + bucket.Name)
  return bucket
}

func Upload(filename string) {
  println("store -- Upload called with filename " + filename)
  bucket := CreateBucket()

  tempDir := "temp/"
  data, err := ioutil.ReadFile(tempDir + filename)
  if err != nil {
    println("store -- error reading file '" + tempDir + filename + "'")
    panic(err)
  }

  extension := filename[(len(filename) - 4):]

  err = bucket.Put(filename, data, mime.TypeByExtension(extension), s3.PublicRead)
  if err != nil {
    println("store -- error during bucket PUT")
    panic(err)
  }
}

func UploadReader(filename string, data io.Reader, dataLength int64) {
  println("store -- UploadReader called with filename " + filename)

  bucket := CreateBucket()

  extension := filename[(len(filename) - 4):]
  err := bucket.PutReader(filename, data, dataLength, mime.TypeByExtension(extension), s3.PublicRead)
  if err != nil {
    println("store -- error during bucket PUT using PutReader")
    panic(err)
  }
}

func Download(filename string) io.ReadCloser {
  println("store -- Download called with filename " + filename)
  bucket := CreateBucket()

  rc, err := bucket.GetReader(filename)

  if err != nil {
    println("store -- error during bucket GET")
    panic(err)
  }

  return rc
}
