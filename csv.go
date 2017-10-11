package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Data for convenience
type Data = [][]string

// Matrix for convenience
type Matrix = [][]float64

// Line for convenience
type Line = []string

// Write writes csv file
func Write(List Data, name string) {
	Title := name
	file, err := os.Create(Title)
	checkError("Cannot create a file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range List {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	fmt.Println("Complete to Write")
}

// Read read csv file
func Read(directory string) Data {
	Title := directory
	file, err := os.Open(Title)
	checkError("Cannot open file", err)

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	return rows
}

// Str2Float converts Data to Matrix
func Str2Float(str Data) Matrix {
	Temp := make(Matrix, len(str), len(str))
	for i := range str {
		for j := range str[i] {
			temp, _ := strconv.ParseFloat(str[i][j], 64)
			Temp[i] = append(Temp[i], temp)
		}
	}
	return Temp
}

// Float2Str converts Matrix to Data
func Float2Str(flt Matrix) Data {
	Temp := make(Data, len(flt), len(flt))
	for i := range flt {
		for j := range flt[i] {
			temp := fmt.Sprint(flt[i][j])
			Temp[i] = append(Temp[i], temp)
		}
	}
	return Temp
}

// Transpose transpose matrix
func Transpose(A Data) Data {
	Temp := make(Data, len(A[0]), len(A[0]))
	// Make Space
	for j := range Temp {
		Temp[j] = make([]string, len(A), len(A))
	}

	// Transpose
	for i := range A {
		for j := range A[i] {
			Temp[j][i] = A[i][j]
		}
	}
	return Temp
}

// AddHeader adds headline to Data
func AddHeader(A Data, L Line) Data {
	if len(A[0]) != len(L) {
		log.Fatal("Different Length detected!")
	}

	B := make(Data, len(A)+1, len(A)+1)

	B[0] = L
	for i := range A {
		B[i+1] = A[i]
	}

	return B
}

//checkError checks Error
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
