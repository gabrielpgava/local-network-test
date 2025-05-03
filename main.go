package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)


func main(){

	fileSize := 2000 // 2GB

	CheckTestFile(fileSize);

	fileserver := http.FileServer(http.Dir("./frontend"));
	

	mux := http.NewServeMux()
	mux.Handle("/", fileserver)
	mux.Handle("/download", http.HandlerFunc(DownloadTestFile))
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
		cmd := exec.Command("dd", "if=/dev/zero", "of=./public/downloadFile.bin", "bs=1M", fmt.Sprintf("count=%d", fileSize))
		output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing command:", err)
                return
        }
        fmt.Println(string(output))
	}
}
