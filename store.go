package store

import (
  "fmt"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
  "io/ioutil"
  "mime"
)


func CreateBucket() *s3.Bucket {
  auth, err := aws.EnvAuth()
  if err != nil {
    println("store -- error connecting to AWS\n")
    panic(err)
  }

  fmt.Printf("store -- connected to AWS\n")

  S3 := s3.New(auth, aws.EUWest)
  bucket := S3.Bucket("go-uploader")
  println("store -- Bucket name: " + bucket.Name)
  return bucket
}

func Upload(filename string) {
  println("store -- Upload called with filename " + filename)
  bucket := CreateBucket()

  tempDir := "temp/"
  data, err := ioutil.ReadFile(tempDir + filename)
  if err != nil {
    println("store -- error reading file" + tempDir + filename)
  }

  extension := filename[(len(filename)-4):]

  err = bucket.Put(filename, data, mime.TypeByExtension(extension), s3.PublicRead)
  if err != nil {
    panic(err)
  }
}
