package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed frontend/*
var embeddedFile embed.FS

func main(){

	fileSize := 2000 // 2GB

	CheckTestFile(fileSize);


	frontendFS, err := fs.Sub(embeddedFile, "frontend")
	if err != nil {
		log.Fatalf("Error creating sub filesystem: %v", err)
	}

	fileserver := http.FileServer(http.FS(frontendFS));
	

	mux := http.NewServeMux()
	mux.Handle("/", fileserver)
	mux.Handle("/download", http.HandlerFunc(DownloadTestFile))
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}


func DownloadTestFile(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Disposition", "attachment; filename=downloadFile.bin")
    w.Header().Set("Content-Type", "application/octet-stream")
    http.ServeFile(w, r, "./public/downloadFile.bin")

}


func CheckTestFile(fileSize int) {
	if _, err := os.Stat("./public/downloadFile.bin"); os.IsNotExist(err) {
		fmt.Println("File does not exist, creating it...")
		if err := os.MkdirAll("./public", os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		err := CreateLargeFile("./public/downloadFile.bin", fileSize)
        if err != nil {
                fmt.Println("Error executing command:", err)
                return
        }
        fmt.Println("File created successfully at ./public/downloadFile.bin")
	}
}


func CreateLargeFile(path string, sizeMB int) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := make([]byte, 1024*1024) // 1MB buffer
    for i := 0; i < sizeMB; i++ {
        if _, err := file.Write(buf); err != nil {
            return err
        }
    }
    return nil
}