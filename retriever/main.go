package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/reals"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmooc",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}


func session(s RetrieverPoster)  {
	s.Get()
	s.Post()
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a faker"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	r = &reals.Retriever{
		UserAgent: "Mozilla/5.0",
		TimOut:    time.Minute,
	}
	inspect(r)
	fmt.Printf("%T %v\n", r, r)
	realRetriever := r.(*reals.Retriever)
	fmt.Println(realRetriever.TimOut)

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contens: ", v.Contents)
	case *reals.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)

	}
}
