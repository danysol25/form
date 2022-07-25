package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//base para parámetros. Datos que pediremos a API

type Base struct(
	ID int `json:"ID"`
	Nombre string `json:"Nombre"`
	Genero string `json:"Genero"`

)

type Peticiones []Base

var peticion = Peticiones{{
	ID: 1,
	Nombre: "Juan",
	Genero: "Masculino",
}}

//Funciones de ruta

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Bienvenido a mi API REST")
}

func main(){
	ruta := mux.NewRouter().StrictSlash(true)

	ruta.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8000", ruta))
}