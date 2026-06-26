package main

import (
   "net/http"
   "os"
   "os/exec"
  )

func main() {
   http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
       w.Write([]byte("ok
                      "))
                       })
                http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                    b, e := exec.Command("nvidia-smi").CombinedOutput()
                    w.Write([]byte("CUDA_VISIBLE_DEVICES=" + os.Getenv("CUDA_VISIBLE_DEVICES") + "
                                   "))
                                     if e != nil {
                                          w.Write([]byte(e.Error() + "
                                                         "))
                                                           }
                                                           w.Write(b)
                                                          })
                                                   p := os.Getenv("PORT")
                                                   if p == "" {
                                                       p = "8000"
                                                      }
                                                   http.ListenAndServe(":"+p, nil)
                                                  }
                                                  
