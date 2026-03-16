package main

import (
	"fmt"
	"time"

	"github.com/BlazejUl/pwr-ite-pea-1/atsp"
	"github.com/BlazejUl/pwr-ite-pea-1/graph"
	"github.com/BlazejUl/pwr-ite-pea-1/utils"
)

func main() {
	var G graph.Graph
	OutName := "out\\"
	fileName := ""
	inputV := 0
	k := 1
	for {
		PrintMenu()
		var opt int
		if _, err := fmt.Scanln(&opt); err != nil {
			fmt.Println(err)
		}
		switch opt {
		case 1:
			fmt.Println("1 - załaduj macież z pliku")
			fmt.Println("2 - wygeneruj macież")
			if _, err := fmt.Scanln(&opt); err != nil {
				fmt.Println(err)
			}
			switch opt {
			case 1:
				fmt.Println("Podaj nazwę pliku .txt z macieżą (musi znajdować się w folderze in)")
				if _, err := fmt.Scanln(&fileName); err != nil {
					fmt.Println(err)
				} else {
					if G, err = utils.ReadGraphFromFile("in\\" + fileName); err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(G.ToString())
					}
				}

			case 2:
				fmt.Println("Podaj liczbę miast")
				if _, err := fmt.Scanln(&inputV); err != nil {
					fmt.Println(err)
				} else {
					if G, err = utils.GenerateAdMatrix(inputV); err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(G.ToString())
					}
				}
			default:
				fmt.Println("tylko numery od 1 - 2")
			}

		case 2:
			if G == nil {
				fmt.Println("graf nie został podany")
				break
			}

			bf := atsp.NewBruteforce(G)
			start := time.Now()
			cost, path := bf.Solve(0)
			lTimeInMikro := time.Since(start).Microseconds()
			name := fmt.Sprintf("%dBruteForce", k)
			k++
			info := fmt.Sprintf("czas: %d µs\nkoszt: %d\nścieżka: %d", lTimeInMikro, cost, path)
			if err := utils.WriteFile(OutName+name, info); err != nil {
				fmt.Printf("////////error %d", err)
			}
			lw := fmt.Sprintf("%d\n", G.GetVerticesNum())
			if err := utils.WriteFile(OutName+"LastMatrix.txt", lw+G.ToString()); err != nil {
				fmt.Printf("////////error %d", err)
			}

			fmt.Printf("czas: %d µs\nkoszt: %d\nścieżka: %d\n", lTimeInMikro, cost, path)

		case 3:

		case 4:

		case 5:

		case 6:
			return
		default:
			fmt.Println("tylko numery od 1 - 7")
		}
	}

}

func PrintMenu() {
	fmt.Println("program do testowania rozwiązań problemu atsp")
	fmt.Println("1 - załaduj macież")
	fmt.Println("2 - rozwiąż za pomocą bruteforce")
	fmt.Println("3 - rozwiąż za pomocą")
	fmt.Println("4 - rozwiąż za pomocą")
	fmt.Println("5 - rozwiąż za pomocą")
	fmt.Println("6 - Wyjdź")
}
