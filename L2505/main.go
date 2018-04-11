package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	r *httprouter.Router
}