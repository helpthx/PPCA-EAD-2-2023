package main

import (
	"fmt" )

type Item struct {
    x  int
    y  int
}

func main() {

   var n, m int
   var h int
   var tower_height = 0
   var acc = 0

   fmt.Scanf("%d %d", &n, &m)

   var tower = []Item{}
   //var field [][]rune
   field := make([][]rune, n, m)
   var aux rune
   for i := 0; i < n; i++{
	for j := 0; j < m; j++ {
		fmt.Scanf("%c", &aux)
		if aux == '\n'{
			fmt.Scanf("%c", &aux);
		}
		field[i] = append(field[i], aux)
		if aux == 't'{
		 	tower = append(tower, Item{x: i, y: j})
			tower_height++
		 }
		}
	}

	fmt.Println(tower)
	for i := 0; i < n; i++{
		for j := 0; j < m; j++ {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < tower_height; i++{
		fmt.Scanf("%d", &h)
		min_x := tower[i].x - h 
		max_x := tower[i].x + h
		for j := min_x; j <= max_x; j++{
			min_y := tower[i].y - h
			max_y := tower[i].y + h
			for k := min_y; k <= max_y; k++ {
				if j >= 0 && j < n && k >= 0 && k < m && field[j][k] == '#'{
					acc++
					field[j][k] = '.'
				} 
			}

		}
	}


	fmt.Println(acc)

	for i := 0; i < n; i++{
		for j := 0; j < m; j++ {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Println()
	}



}