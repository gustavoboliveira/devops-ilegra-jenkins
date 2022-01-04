package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

const bite = 64
var hist []string

func calcule (w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    x,errX := strconv.ParseFloat(params["x"], bite)
    y,errY := strconv.ParseFloat(params["y"], bite)
    op,_ := params["op"]

    if errX != nil{
        fmt.Fprintf(w, "Error: can't convert first number")
    } else if errY != nil {
        fmt.Fprintf(w, "Error: can't convert second number")
    } else {
        switch op {
            case "sum":
                result := x + y
                build_expression(x, y, result, "+")
                fmt.Fprintf(w, "%.2f", result)
            case "sub":
                result := x - y
                build_expression(x, y, result, "-")
                fmt.Fprintf(w, "%.2f", result)
            case "mul":
                result := x * y
                build_expression(x, y, result, "*")
                fmt.Fprintf(w, "%.2f", result)
            case "div":
                if y == 0 {
                    fmt.Fprintf(w, "Error: can't possible divide by zero")
                } else {
                    result := x / y
                    build_expression(x, y, result, "/")
                    fmt.Fprintf(w, "%.2f", result)
                }
            default:
                fmt.Fprintf(w, "Error: operation not available")
        }
    }
}

func history (w http.ResponseWriter, r *http.Request) {
    if len(hist) == 0 {
        fmt.Fprintf(w, "Error: no operations registred yet")
    } else {
        for _, h := range hist {
            fmt.Fprintf(w, h + "\n")
        }
    }
}

func save_history(h string){
    hist = append(hist, h)
}

func build_expression(x, y, result float64, op string) string{
    h := fmt.Sprintf("%.2f %s %.2f = %.2f", x, op, y, result)
    save_history(h)
    return h
}

func main() {
    fmt.Println("Server running on localhost:8080/")
	router := mux.NewRouter()
	router.HandleFunc("/calc/{op}/{x}/{y}", calcule)
	router.HandleFunc("/calc/history", history)
	http.ListenAndServe(":8080", router)
}