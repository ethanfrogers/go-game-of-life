package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const size int = 20

func printWorld(wrld [size][size]string) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(" " + wrld[i][j] + " ")
		}
		fmt.Println("")
	}
}

func generateNewWorld() [size][size]string {
	// random number generator
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	var world [size][size]string

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if generator.Intn(size)%3 == 0 {
				world[i][j] = "x"
			} else {
				world[i][j] = "_"
			}
		}
	}

	return world
}

func mod(pos int, m int) int {
	return (pos%m + m) % m
}

func evolve(world [size][size]string) [size][size]string {

	var newWorld [size][size]string

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			numNeighbors := 0
			live := world[i][j] == "x"
			//above
			if world[i][mod(j-1, size)] == "x" {
				numNeighbors++
			}
			//above right
			if world[mod(i-1, size)][mod(j+1, size)] == "x" {
				numNeighbors++
			}
			//above left
			if world[mod(i-1, size)][mod(j-1, size)] == "x" {
				numNeighbors++
			}
			//below
			if world[mod(i+1, size)][j] == "x" {
				numNeighbors++
			}
			//below right
			if world[mod(i+1, size)][mod(j+1, size)] == "x" {
				numNeighbors++
			}
			//below left
			if world[mod(i+1, size)][mod(j-1, size)] == "x" {
				numNeighbors++
			}
			//left
			if world[mod(i-1, size)][j] == "x" {
				numNeighbors++
			}
			//right
			if world[mod(i+1, size)][j] == "x" {
				numNeighbors++
			}

			var newValue string

			if numNeighbors < 2 && live {
				newValue = "_"
			} else if (numNeighbors == 2 || numNeighbors == 3) && live {
				newValue = "x"
			} else if numNeighbors > 3 && live {
				newValue = "_"
			} else if numNeighbors == 3 && !live {
				newValue = "x"
			}

			newWorld[i][j] = newValue

		}
	}

	return newWorld
}

func main() {

	world := generateNewWorld()

	fmt.Println("----Base World----")
	printWorld(world)

	for i := 0; i < 5; i++ {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		world = evolve(world)
		fmt.Printf("----Iteration %d ----\n", i)
		printWorld(world)
		time.Sleep(1 * time.Second)

	}
}
