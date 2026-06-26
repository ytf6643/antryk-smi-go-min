package main

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
   )

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
         fmt.Fprintln(w, "ok")
        })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
         out, err := exec.Command("nvidia-smi").CombinedOutput()
         fmt.Fprintln(w, "CUDA_VISIBLE_DEVICES="+os.Getenv("CUDA_VISIBLE_DEVICES"))
         if err != nil {
               fmt.Fprintln(w, err.Error())
              }
         w.Write(out)
        })

    port := os.Getenv("PORT")
    if port == "" {
         port = "8000"
        }
    http.ListenAndServe(":"+port, nil)
   }
