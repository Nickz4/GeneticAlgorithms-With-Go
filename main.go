package main

import (
	"GoWithGeneticAlgorithms/ga"
	"GoWithGeneticAlgorithms/problem"
	"fmt"
	"math/rand"
	"time"
)

func main(){
	settings := ga.GenenticAlgorithms{
		PopulationSize: 10,
		MutationRate: 10,
		CrossoverRate: 50,
		Generations: 20,
		KeepBest: true,
		UpperBoundary: 63,
		LowerBooudary:0,
	}
	//rand的seed在main里设置一次就行，不要在每个需要随机数的函数里都设置
	rand.Seed(time.Now().UnixNano())

	var a ga.GeneticAlgorithmsFunc = problem.QuadraticSolotion{}


	beforeQuadTime := time.Now()
	solution,error := a.Run(settings)
	afterQuadTime := time.Since(beforeQuadTime)

	if error != nil{
		println(error)
	}else {
		fmt.Printf("Best: x: %v  y: %v\n", solution, problem.QuadraticSolotion{}.Evaluate(solution.(int)))
	}
	fmt.Printf("%d\n", afterQuadTime.String())

}