package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type SqrRes struct {
	Massage string
	Result  float64
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/res", calc_handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("html/resources"))))
	http.ListenAndServe(":8000", nil)
}
func calc_handler(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	if r.Method != "POST" || r.FormValue("formNumber") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inputValue := r.FormValue("formNumber")
	number, _ := strconv.ParseInt(inputValue, 10, 0)
	temp.Execute(w, SqrRes{"√" + inputValue + " = ", Round(math.Sqrt(float64(number)), .5, 2)})
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	temp.Execute(w, nil)
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
