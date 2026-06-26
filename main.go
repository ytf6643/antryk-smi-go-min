package main

import (
     "fmt"
     "net/http"
     "os"
     "os/exec"
     "path/filepath"
    )

func run(w http.ResponseWriter, name string) {
     fmt.Fprintln(w, "### "+name)
     out, err := exec.Command(name).CombinedOutput()
     if err != nil {
           fmt.Fprintln(w, err.Error())
          }
     if len(out) == 0 {
           fmt.Fprintln(w, "no output")
          } else {
           w.Write(out)
           fmt.Fprintln(w)
          }
    }

func showFile(w http.ResponseWriter, path string) {
     fmt.Fprintln(w, "### file "+path)
     b, err := os.ReadFile(path)
     if err != nil {
           fmt.Fprintln(w, err.Error())
           return
          }
     w.Write(b)
     fmt.Fprintln(w)
    }

func showGlob(w http.ResponseWriter, pattern string) {
     fmt.Fprintln(w, "### glob "+pattern)
     matches, err := filepath.Glob(pattern)
     if err != nil {
           fmt.Fprintln(w, err.Error())
           return
          }
     if len(matches) == 0 {
           fmt.Fprintln(w, "none")
           return
          }
     for _, m := range matches {
           info, err := os.Stat(m)
           if err != nil {
                  fmt.Fprintln(w, m, err.Error())
                  continue
                 }
           fmt.Fprintln(w, m, info.Mode().String())
           if info.IsDir() {
                  entries, err := os.ReadDir(m)
                  if err != nil {
                          fmt.Fprintln(w, "read", err.Error())
                          continue
                         }
                  for _, e := range entries {
                          fmt.Fprintln(w, " "+e.Name())
                         }
                 }
          }
    }

func main() {
     http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
           fmt.Fprintln(w, "ok")
          })

     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
           fmt.Fprintln(w, "CUDA_VISIBLE_DEVICES="+os.Getenv("CUDA_VISIBLE_DEVICES"))
           fmt.Fprintln(w, "PATH="+os.Getenv("PATH"))
           for _, p := range []string{"nvidia-smi", "/usr/bin/nvidia-smi", "/usr/local/nvidia/bin/nvidia-smi", "/usr/local/cuda/bin/nvidia-smi"} {
                  run(w, p)
                 }
           showGlob(w, "/dev/nvidia*")
           showGlob(w, "/proc/driver/nvidia*")
           showGlob(w, "/proc/driver/nvidia/gpus/*")
           showFile(w, "/proc/driver/nvidia/version")
           infos, _ := filepath.Glob("/proc/driver/nvidia/gpus/*/information")
           for _, p := range infos {
                  showFile(w, p)
                 }
          })

     port := os.Getenv("PORT")
     if port == "" {
           port = "8000"
          }
     http.ListenAndServe(":"+port, nil)
    }
