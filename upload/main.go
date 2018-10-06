package main

import (
    "crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
    "path/filepath"
)
const maxUploadSize = 2 * 1024
const uploadPath = "./tmp"
func uploadFileHandler() http.HandlerFunc {
	return http.HandleFunc(func(w http.ResponseWriter,r *http.Request) {
    r.Body = http.MaxBytesReader(w,r.Body,maxUploadSize)
   if err := r.ParseMultipartForm(maxUploadSize); err != nil {
	   renderError(w,"FILE_TO_BIG",http.StatusBadRequest)
	   return
   }
   fileType := r.PostFormValue("type")
   file,_,err := r.FormFile("uploadFile")
   if err != nil {
	   renderError(w,"INVALID_FILE",http.StatusBadRequest)
	   return
   }
   defer file.Close()
   fileBytes, err := ioutil.ReadAll(file)
   if err != nil {
	   renderError(w,"INVALID_FILE",http.StatusBadRequest)
	   return
   }
   filetype := http.DetectContentType(fileBytes)
  switch filetype {
    case "image/jpeg", "image/jpg":
    case "image/gif","image/png":
    case "application/pdf":
	   break
   default:
	renderError(w,"INVALID_FILE",http.StatusBadRequest)
	return	   
  }
  fileName := randToken(12)
  fileEndings, err := mime.ExtensionsByType(fileType)
 if err != nil {
	  renderError(w,"CANT_READ_TYPE",http.StatusInternalServerError)
	  return
   }
   newPath := filepath.Join(uploadPath,fileName+fileEndings[0])
   fmt.Printf("Filetype: %s,File:%s\n",fileType,newPath)
   newFile,err := os.Create(newPath)
   if err != nil {
	   renderError(w,"CANT_WRITE_FILE",http.StatusInternalServerError)
	   return
   }
   defer newFile.Close()
   if _,err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
   renderError(w,"CANT_WRITE file",http.StatusInternalServerError)   
    return
   }
   w.Write([]byte("SUCCESS"))
  } )
}  
func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
func randToken(len int) string {
	b := make([]byte,len)
	rand.Read(b)
	return fmt.Sprintf("%x",b)
}
func main() {
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}