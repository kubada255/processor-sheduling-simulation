package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type polecenie struct {
	czas_p    int
	czas_d    int
	priorytet int
}

var suma int

// fcfs scheduling
func planowanie_fcfs(polecenia []polecenie) {
	fmt.Println("FCFS")
	ilosc := len(polecenia)
	kolejka := []polecenie{}
	kolejka = append(kolejka, polecenia[0])
	polecenia = polecenia[1:]
	takt := 0
	czas := 0

	for len(kolejka) > 0 {

		for {
			if len(polecenia) > 0 && polecenia[0].czas_p <= takt {
				kolejka = append(kolejka, polecenia[0])
				polecenia = polecenia[1:]
			} else {
				break
			}
		}
		kolejka[0].czas_d--
		takt++

		if kolejka[0].czas_d == 0 {
			czas = czas + (takt - kolejka[0].czas_p)
			kolejka = kolejka[1:]

		}

	}
	fmt.Println("Średni czas oczekiwania: ", float64(czas-suma)/float64(ilosc))
}

// sjf scheduling
func planowanie_sjf(polecenia []polecenie) {
	fmt.Println("SJF")
	ilosc := len(polecenia)
	kolejka := []polecenie{}
	takt := 0
	czas := 0
	for {
		for {
			if len(polecenia) > 0 && polecenia[0].czas_p <= takt {
				kolejka = append(kolejka, polecenia[0])
				polecenia = polecenia[1:]
			} else {
				break
			}
		}
		if len(kolejka) == 0 {
			break
		}
		if kolejka[len(kolejka)-1].czas_p == 0 {
			//sorting by czas_d
			for i := 0; i < len(kolejka); i++ {
				for j := 0; j < len(kolejka)-1; j++ {
					if kolejka[j].czas_d > kolejka[j+1].czas_d {
						kolejka[j], kolejka[j+1] = kolejka[j+1], kolejka[j]
					}
				}
			}

		}
		kolejka[0].czas_d--
		takt++
		if kolejka[0].czas_d == 0 {
			czas = czas + (takt - kolejka[0].czas_p)
			kolejka = kolejka[1:]
			//sorting by czas_d
			for i := 0; i < len(kolejka); i++ {
				for j := 0; j < len(kolejka)-1; j++ {
					if kolejka[j].czas_d > kolejka[j+1].czas_d {
						kolejka[j], kolejka[j+1] = kolejka[j+1], kolejka[j]
					}
				}
			}
		}
	}
	fmt.Println("Średni czas oczekiwania: ", float64(czas-suma)/float64(ilosc))
}
func priorytetetowe(polecenia []polecenie, postarzanie int) {
	fmt.Println("Priorytetowe")
	ilosc := len(polecenia)
	kolejka := []polecenie{}
	takt := 0
	czas := 0
	for {
		for {
			if len(polecenia) > 0 && polecenia[0].czas_p <= takt {
				kolejka = append(kolejka, polecenia[0])
				polecenia = polecenia[1:]
			} else {
				break
			}
		}
		if len(kolejka) == 0 {
			break
		}
		if kolejka[len(kolejka)-1].czas_p == 0 {
			//sorting by priorytet
			for i := 0; i < len(kolejka); i++ {
				for j := 0; j < len(kolejka)-1; j++ {
					if kolejka[j].priorytet > kolejka[j+1].priorytet {
						kolejka[j], kolejka[j+1] = kolejka[j+1], kolejka[j]
					}
				}
			}

		}
		kolejka[0].czas_d--
		takt++
		if takt%postarzanie == 0 {
			for i := 1; i < len(kolejka); i++ {
				if kolejka[i].priorytet > 0 {
					kolejka[i].priorytet--
				}
			}
		}
		if kolejka[0].czas_d == 0 {
			czas = czas + (takt - kolejka[0].czas_p)
			kolejka = kolejka[1:]
			//sorting by priorytet
			for i := 0; i < len(kolejka); i++ {
				for j := 0; j < len(kolejka)-1; j++ {
					if kolejka[j].priorytet > kolejka[j+1].priorytet {
						kolejka[j], kolejka[j+1] = kolejka[j+1], kolejka[j]
					}
				}
			}
			//sorting by czas_d
			for i := 0; i < len(kolejka); i++ {
				if kolejka[i].priorytet == kolejka[0].priorytet {
					for j := 0; j < len(kolejka)-1; j++ {
						if kolejka[j].czas_d > kolejka[j+1].czas_d {
							kolejka[j], kolejka[j+1] = kolejka[j+1], kolejka[j]
						}
					}
				} else {
					break
				}
			}
		}
	}
	fmt.Println("Średni czas oczekiwania: ", float64(czas-suma)/float64(ilosc))
}
func main() {
	var fileName string
	fmt.Print("Podaj nazwę pliku: ")
	fmt.Scanln(&fileName)
	polecenia := []polecenie{}
	czas_p := []int{}
	czas_d := []int{}
	priorytet := []int{}
	postarzanie := 0
	//Reading file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {

		fmt.Println("Blad czytania pliku", err)
		fmt.Scanf(" ")
		return
	}
	linie := strings.Split(string(data), "\r\n")

	//Spliting lines
	for i, linia := range linie {
		slowa := strings.Split(linia, " ")
		for _, slowo := range slowa {

			liczba, err := strconv.Atoi(slowo)
			if err != nil {
				fmt.Println("Blad czytania pliku", err)
				fmt.Scanf(" ")
				return
			}
			if i == 0 {
				czas_p = append(czas_p, liczba)
			}
			if i == 1 {
				czas_d = append(czas_d, liczba)
			}
			if i == 2 {
				priorytet = append(priorytet, liczba)
			}
			if i == 3 {
				postarzanie = liczba
			}
		}

	}
	//Adding to slice
	for i := 0; i < len(czas_p); i++ {
		polecenia = append(polecenia, polecenie{czas_p[i], czas_d[i], priorytet[i]})
	}
	suma = 0
	for _, i := range czas_d {
		suma = suma + i
	}
	planowanie_fcfs(polecenia)
	planowanie_sjf(polecenia)
	priorytetetowe(polecenia, postarzanie)

	fmt.Scanln()

}
