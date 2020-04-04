package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const port = 8090

func main() {
	router := httprouter.New()
	router.POST("/backend", backend)
	router.GET("/function", functionGet)
	// router.POST("/function", functionPost)
	// router.POST("/compile", compile)
	router.POST("/format", format)

	handler := cors.AllowAll().Handler(router)
	log.Printf("Authentication Server listenning on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func backend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	obj := struct {
		ID string `json:"id"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func functionGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content, err := ioutil.ReadFile("server/server.go")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		Data string `json:"data"`
	}{
		Data: string(content),
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func format(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	obj := struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	f, err := ioutil.TempFile("formatter", "tmp")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	defer os.Remove(f.Name())
	if _, err := io.Copy(f, strings.NewReader(obj.Data)); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := os.Rename(f.Name(), "formatter/server.go"); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("make", "fmt")
	cmd.Dir = "./formatter"
	out, err := cmd.CombinedOutput()
	if err != nil {
		s := strings.TrimPrefix(string(out), "go fmt server.go\n")
		res := struct {
			Err string `json:"err"`
		}{
			Err: s,
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resJSON)
		return
	}

	// copied from functionGet
	content, err := ioutil.ReadFile("formatter/server.go")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		Data string `json:"data"`
	}{
		Data: string(content),
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
