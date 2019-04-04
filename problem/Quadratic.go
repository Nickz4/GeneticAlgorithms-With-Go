package problem

import (
	"GoWithGeneticAlgorithms/ga"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	//"time"

	//"./ga"
)
var(
	CrossoverDEBUG = false
	RunDEBUG = true
	MutationDEBUG = false
	ChoiceDEBUG = false
)
type DataProcess struct{
	DecimalNumber interface{}
}

func (d DataProcess) String() string{
	//现在是固定长6，之后应添加动态长度，根据二进制最长的基因，另一个也应达到同样长度
	return fmt.Sprintf("%06b",d.DecimalNumber)
}


type QuadraticSolotion struct{}


func (q QuadraticSolotion) Evaluate(x int) int {
	return  x*x-10*x+5 // minimum should be at x=5
}


func (q QuadraticSolotion) Mutation(chromosome interface{}, mutationRate int) interface{}{
	var result string
	c := DataProcess{chromosome}
	for _,q := range c.String(){
		if rand.Intn(100) < mutationRate {
			//String在获取的时候是以ascii形式获取的，48为0 49为1
			if q == 48 {
				q = 49
			} else {
				q = 48
			}
		}
		result += string(q)
		if MutationDEBUG{
			fmt.Printf("mutation result:%v\n",result)
		}
	}
	// 二进制转int8
	i, err := strconv.ParseInt(result, 2, 8)
	if err != nil {
		panic(err)
	}
	return int(i)

}
func (q QuadraticSolotion) InitiateFirstPopulation(min,max,populationSize int) []interface{} {
	initialPopulation := make([]interface{},0, populationSize)
	for i:=0; i < populationSize; i++{
		initialPopulation = append(initialPopulation, rand.Intn(max - min) + min)
	}

	return initialPopulation
}


func (q QuadraticSolotion) Crossover(chromesome1, chromosome2 interface{},crossoverRate int) interface{}{
	var result string
	c1 := DataProcess{chromesome1}
	c2 := DataProcess{chromosome2}
	if CrossoverDEBUG{
		fmt.Printf("crossover chromosome1:%v\n",c1)
		fmt.Printf("crossover chromosome2:%v\n",c2)
	}
	for i,q := range c1.String(){
		if rand.Intn(100) < crossoverRate{
			result += string(q)
		}else {
			result += string(c2.String()[i])
		}
		if CrossoverDEBUG{
			fmt.Printf("crossover result:%v\n",result)
		}
	}
	// 二进制转int8
	res, err := strconv.ParseInt(result, 2, 8)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func (q QuadraticSolotion) Sort(population []interface{}){
	//定义Less()方法，从大到小排序
	sort.Slice(population, func(i, j int) bool {
		return q.Evaluate(population[i].(int)) > q.Evaluate(population[j].(int))
	})
}

func (q QuadraticSolotion) Choice(population []interface{}) []interface{} {
	capacity:=0
	for i:=0; i < len(population); i++{
		capacity += i
	}
	if ChoiceDEBUG{
		fmt.Printf("capacity:%v\n",capacity)
	}
	newPopulation := make([]interface{},0, capacity)
	for i,chromosome := range population{
		//基因写入的个数取决于index，index越多的基因，在新population中占比就越多
		for j:=0; j< i;j++{
			newPopulation = append(newPopulation,chromosome)
		}
	}
	return newPopulation
}

func (q QuadraticSolotion) Run(settings ga.GenenticAlgorithms) (interface{}, error){
	//初始化种群
	population := q.InitiateFirstPopulation(settings.LowerBooudary,settings.UpperBoundary,settings.PopulationSize)
	//排序找到目前最优解
	q.Sort(population)
	if RunDEBUG{
		fmt.Printf("fist:%v\n",population)
	}

	bestSolution := population[len(population)-1]

	for i:=0; i< settings.Generations; i++{
		nextGeneration := make([]interface{}, 0, settings.PopulationSize)

		if settings.KeepBest {
			//保留上一代的最佳
			nextGeneration= append(nextGeneration, bestSolution)
		}
		//扩充上一代的基因池，根据解的优秀程度，越优秀基因占比越多
		choicePopulation := q.Choice(population)
		if RunDEBUG{
			fmt.Printf("choicePopulation:%v\n",choicePopulation)
		}
		//获取新种群的长度，如果保留了上一代最佳，现在应该为1
		populationLen := len(nextGeneration)
		//填充下一代种群
		for ; populationLen < settings.PopulationSize; populationLen ++{

			//从上一代基因池随机选择基因进行crossover
			chromosome1Index := rand.Int() % len(choicePopulation)
			chromosome2Index := rand.Int() % len(choicePopulation)
			if RunDEBUG{
				fmt.Printf("chromosome1:%v,chromosome2:%v\n",choicePopulation[chromosome1Index],choicePopulation[chromosome2Index])
			}

			//crossover
			child := q.Crossover(choicePopulation[chromosome1Index],choicePopulation[chromosome2Index],settings.CrossoverRate)
			if RunDEBUG{
				fmt.Printf("crossover:%v\n",child)
			}
			//mutation
			newChild := q.Mutation(child,settings.MutationRate)
			if RunDEBUG{
				fmt.Printf("mutation:%v\n",newChild)
			}
			nextGeneration = append(nextGeneration, newChild)
			if RunDEBUG{
				fmt.Printf("nextGeneration:%v\n",nextGeneration)
			}

		}
		population = nextGeneration

		q.Sort(population)
		if RunDEBUG {
			fmt.Printf("population:%v\n",population)
		}
		bestSolution = population[len(population)-1]
	}

	return bestSolution,nil
}