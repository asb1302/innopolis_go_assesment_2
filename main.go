package main

import (
	"errors"
)

// GraphLists граф в виде списков смежности
type GraphLists struct {
	adjList map[int][]int
	weights map[int]map[int]int
}

func EvalSequence(matrix [][]int, userAnswer []int) (int, error) {
	if !validateMatrix(matrix) {
		return 0, errors.New("invalid matrix")
	}

	if !validateUserAnswer(matrix, userAnswer) {
		return 0, errors.New("invalid user answer")
	}

	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	if maxGrade == 0 {
		return 0, nil
	}

	percent := (userGrade * 100) / maxGrade

	return percent, nil
}

func validateMatrix(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		// матрица должна быть квадратной - количество строк должно совпадать с количеством столбцов
		if len(matrix[i]) != n {
			return false
		}

		// в графе не может быть петель - значения на главной диагонали должны быть равны нулю
		if matrix[i][i] != 0 {
			return false
		}

		// проверка на отрицательные значения
		for j := 0; j < n; j++ {
			if matrix[i][j] < 0 {
				return false
			}
		}
	}

	return true
}

func validateUserAnswer(matrix [][]int, userAnswer []int) bool {
	seen := make(map[int]bool)

	for _, v := range userAnswer {
		// ответы пользователя не должны выходить за диапазон матрицы
		if v < 0 || v >= len(matrix) {
			return false
		}

		// элементы в слайсе ответов пользователя должны быть уникальными
		if seen[v] {
			return false
		}

		seen[v] = true
	}

	return true
}

func calMaxGrade(matrix [][]int) int {
	graph := matrixToGraph(matrix)
	maxGrade := 0

	for vertex := range graph.adjList {
		visited := make(map[int]bool)

		grade := dfs(vertex, visited, graph, 0)

		if grade > maxGrade {
			maxGrade = grade
		}
	}

	return maxGrade
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	userGrade := 0

	for i := 0; i < len(userAnswer)-1; i++ {
		userGrade += matrix[userAnswer[i]][userAnswer[i+1]]
	}

	return userGrade
}

func matrixToGraph(matrix [][]int) GraphLists {
	// adjList: Список смежности для хранения соседей каждой вершины
	// weights: Мапа для хранения весов ребер. Представлена как map[int]map[int]int, где ключи - это вершины, а значения - это карты с весами ребер от вершины к ее соседям
	graph := GraphLists{
		adjList: make(map[int][]int),
		weights: make(map[int]map[int]int),
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				if graph.weights[i] == nil {
					graph.weights[i] = make(map[int]int)
				}

				graph.adjList[i] = append(graph.adjList[i], j)
				graph.weights[i][j] = matrix[i][j]
			}
		}
	}

	return graph
}

// Функция выполняет поиск в глубину для нахождения пути с максимальной суммой весов в графе
func dfs(node int, visited map[int]bool, graph GraphLists, currentGrade int) int {
	visited[node] = true

	maxGrade := currentGrade

	for _, neighbor := range graph.adjList[node] {
		if !visited[neighbor] {
			grade := dfs(neighbor, visited, graph, currentGrade+graph.weights[node][neighbor])

			if grade > maxGrade {
				maxGrade = grade
			}
		}
	}

	// Снимаем отметку о посещении для текущей вершины, чтобы она могла быть использована в других путях
	visited[node] = false

	return maxGrade
}
