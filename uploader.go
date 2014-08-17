package uploader

import "fmt"
import "github.com/mitchellh/goamz/aws"
import "github.com/mitchellh/goamz/s3"
import "io/ioutil"

func Uploader() {
  fmt.Printf("go-uploader\n")

  auth, err := aws.EnvAuth()
  if err != nil {
    fmt.Printf("error connecting to AWS\n")
    panic(err)
  }

  fmt.Printf("connected to AWS\n")

  S3 := s3.New(auth, aws.EUWest)
  bucket := S3.Bucket("go-uploader")
  println("Bucket name: " + bucket.Name)

  println("Listing files in root folder:")
  objects, err := bucket.List("", "", "", 100)
  for _, value := range objects.Contents {
    fmt.Printf(value.Key + "\n")
  }

  data, err := bucket.Get("Erlend-Rosjo.jpg")

  err = ioutil.WriteFile("temp/Erlend-Rosjo.jpg", data, 0644)
  if err != nil {
    panic(err)
  }
}