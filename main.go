package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

var jsonStr = []byte(`
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

type Things struct {
	Things []Combined `json:"things"`
}

type Combined struct {
	Name string `json:"name"`
	Age int `json:"age"`
	City string `json:"city"`
	Country string `json:"country"`
}

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type HumanDecoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})
	Print(interface{})
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func main() {
	logger := log.New(os.Stdout, "INFO:", 0)
	s := Service {
		log: logger,
	}
	persons, places := s.Decode(jsonStr)
	s.Sort(persons)
	s.Sort(places)

	s.Print(persons)
	s.Print(places)
}

func (s *Service) Decode(data []byte) ([]Person, []Place) {
	var things Things
	err := json.Unmarshal(data, &things)
	if err != nil {
		s.log.Fatalf("error %v", err)
	}

	persons := make([]Person, 0)
	places := make([]Place, 0)
	for _, v := range things.Things {
		if v.City == "" {
			person := Person{
				Name: v.Name,
				Age: v.Age,
			}
			if person.Age > 0 {
				persons = append(persons, person)
			}
		}

		if v.Name == "" {
			place := Place {
				City: v.City,
				Country: v.Country,
			}

			places = append(places, place)
		}
	}

	return persons, places
}

func (s *Service) Sort(dataToSort interface{}) {
	switch data := dataToSort.(type) {
	case []Person:
		sort.Slice(data, func(i, j int) bool {
			return len(data[i].Name) < len(data[j].Name)
		})
	case []Place:
		sort.Slice(data, func(i, j int) bool {
			return len(data[i].City) < len(data[j].City)
		})
	default:
		fmt.Println("cant sort")
	}
}

func (s *Service) Print(info interface{}) {
	s.log.Println(info)
}