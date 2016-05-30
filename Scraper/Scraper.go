/*
	Tarea 3 - CIT-2001 - Diseño y Análisis de Algoritmos - Universidad Diego Portales
   	Autor: Guillermo Iglesias Birkner
   	Profesor: Francisco Claude - Ayudante: Marcello Tavano
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Declaración de variables a utilizar.
	conf_5 := 0
	conf_3 := 0
	journal_5 := 0
	journal_3 := 0

	URL_prefix := "http://dblp.uni-trier.de/pers/hd/"

	// Input: Nombre Apellido.
	person := os.Args

	fmt.Println("Nombre:", person[1])
	fmt.Println("Apellido:", person[2])

	// Concatenacion de la URL.
	URL_final := []string{URL_prefix, strings.ToLower(string(person[2][0])), "/", person[2], ":", person[1]}
	fmt.Println("URL:", strings.Join(URL_final, ""))

	// Extracción de la URL.
	web, err := http.Get(strings.Join(URL_final, ""))

	if err != nil {
		log.Fatal(err)
	}

	// Lectura de la URL.
	source, err := ioutil.ReadAll(web.Body)

	web.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(source), "Error 404") == true {
		log.Fatalln("Error 404: Pagina no encontrada.")
	}

	// Busqueda de cantidad de Revistas y Conferencias.
	c_prefix := regexp.MustCompile("inproceedings..id=.conf/............................")
	c_matches := c_prefix.FindAllString(string(source), -1)

	j_prefix := regexp.MustCompile("article..id=.journals/............................")
	j_matches := j_prefix.FindAllString(string(source), -1)

	// Busquedad de cantidad de Conferencias en los últimos años.
	for i := range c_matches {
		numbers := regexp.MustCompile("[0-9][0-9]")
		n_matches := numbers.FindString(c_matches[i])

		year, err := strconv.Atoi(n_matches)

		if err != nil {
			log.Fatal(err)
		}

		if year > 10 && year < 50 {
			conf_5++
			if year > 12 {
				conf_3++
			}
		}
	}

	// Busquedad de cantidad de Revistas en los últimos años.
	for i := range j_matches {
		numbers := regexp.MustCompile("[0-9][0-9]")
		j_matches := numbers.FindString(j_matches[i])

		year, err := strconv.Atoi(j_matches)

		if err != nil {
			log.Fatal(err)
		}

		if year > 10 && year < 50 {
			journal_5++
			if year > 12 {
				journal_3++
			}
		}
	}

	// Impresión cantidad de Revistas y Conferencias.
	fmt.Println("Revistas:", len(j_matches))
	fmt.Println("Conferencias:", len(c_matches))
	fmt.Println("Revistas (5 años):", journal_5)
	fmt.Println("Conferencias (5 años):", conf_5)
	fmt.Println("Revistas (5 años):", journal_3)
	fmt.Println("Conferencias (5 años):", conf_3)

}
