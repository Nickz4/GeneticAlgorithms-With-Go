package ga

type GenenticAlgorithms struct{
	MutationRate int
	CrossoverRate int
	PopulationSize int
	Generations int
	KeepBest bool
	UpperBoundary int
	LowerBooudary int
}
/*
golang的命名需要使用驼峰命名法，且不能出现下划线
golang中根据首字母的大小写来确定可以访问的权限。
无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；
如果首字母小写，则只能在本包中使用
*/
type GeneticAlgorithmsFunc interface{
	Mutation(chromosome interface{}, mutationRate int) interface{}
	Crossover(chromosome1,chromosome2 interface{},crossoverRate int) interface{}
	InitiateFirstPopulation(min,max,populationSize int) []interface{}
	Choice(population []interface{}) []interface{}
	Run(settings GenenticAlgorithms) (interface{}, error)
	Sort([]interface{})
	Evaluate(x int) int
}

