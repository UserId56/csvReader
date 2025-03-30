package main

import (
	"csvReader/internal/csvReader"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	userId string
	name   string
	age    int
	city   string
}

func main() {
	var (
		path     string
		typeWork uint8
	)
	fmt.Print("Укажите путь до csv файла, включая имя и расширение: ")
	fmt.Scan(&path)
	fmt.Print("Укажите 0 если читать файл полностью и 1 - построчно: ")
	_, err := fmt.Scan(&typeWork)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch typeWork {
	case 0:
		result, err := csvReader.ReadCSVAll(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
		personsData := getPersonfromData(result)
		fmt.Println(personsData)
		average, count := 0, 0
		for _, v := range personsData {
			count++
			average += v.age
		}
		fmt.Printf("Средний возраст: %d", average/count)
	case 1:
		result, err := csvReader.ReadLargeCSV(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
		personsData := getPersonfromData(result)
		fmt.Println(personsData)
		average, count := 0, 0
		for _, v := range personsData {
			count++
			average += v.age
		}
		fmt.Printf("Средний возраст: %d", average/count)
	}
}

func getPersonfromData(dataPErson [][]string) []Person {
	persons := []Person{}
	for i := 1; i < len(dataPErson)-1; i++ {
		age, _ := strconv.Atoi(dataPErson[i][2])
		persons = append(persons, Person{
			userId: dataPErson[i][0],
			name:   dataPErson[i][1],
			age:    age,
			city:   dataPErson[i][3],
		})
	}
	return persons
}
