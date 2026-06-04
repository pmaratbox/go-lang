package main

import "fmt"

func main() {
	const inf = 1 << 30
	n := 3
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = inf
			}
		}
	}
	add := func(u, v, w int) { dist[u][v] = w }
	add(0, 1, 3)
	add(1, 2, 1)
	add(0, 2, 5)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	fmt.Println(dist[0][2])
}
