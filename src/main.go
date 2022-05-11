package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
}

// {
//     "Name": "Sanji",
//     "Age": 20,
//     "Department": "Management",
//     "Salary": 50000
// }
func newEmp(c *gin.Context) {
	var newemp Employee
	err := c.BindJSON(&newemp)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("input.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	agec := strconv.Itoa(newemp.Age)
	salc := strconv.Itoa(newemp.Salary)
	row := []string{newemp.Name, agec, newemp.Department, salc}
	er := w.Write(row)
	w.Flush()
	if er != nil {
		panic(er)
	}
	c.JSON(http.StatusCreated, newemp)
}

func getEmployees(c *gin.Context) {
	file, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	records, er := csv.NewReader(file).ReadAll()
	if er != nil {
		panic(er)
	}
	var employee []Employee
	for _, record := range records {
		agec, _ := strconv.Atoi(record[1])
		salc, _ := strconv.Atoi(record[3])
		d := Employee{
			Name:       record[0],
			Age:        agec,
			Department: record[2],
			Salary:     salc,
		}
		employee = append(employee, d)
	}
	c.JSON(http.StatusOK, employee)
}

func main() {
	router := gin.Default()
	router.GET("/empdata", getEmployees)
	router.POST("/empdata", newEmp)
	router.Run()
}
