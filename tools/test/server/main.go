package main

import "net/http"

import "os/exec"

func main() {
	http.HandleFunc("/ping", restart)
	http.ListenAndServe(":5001", nil)
}

func restart(w http.ResponseWriter, r *http.Request) {
	refresh := r.URL.Query().Get("cmd")
	if refresh == "refresh" {
		if err := execCmd(); err != nil {
			w.Write([]byte("successfully"))
			return
		}
	}
	w.Write([]byte("failed"))
}

func execCmd() error {
	exec.Command("cd /project&&go run main.go")
	return nil
}
