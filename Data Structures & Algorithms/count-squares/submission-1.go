type CountSquares struct {
    pointMap map[Point]int
	x_max int
	y_max int
	x_min int
	y_min int
}

type Point struct {
	x int
	y int
}

func Constructor() CountSquares {
    return CountSquares {
		pointMap: make(map[Point]int),
		x_max: -1,
		y_max: -1,
		x_min: 1001,
		y_min: 1001,
	}
}

func (this *CountSquares) Add(point []int)  {
	this.pointMap[Point{x: point[0], y: point[1]}]++

	this.x_max = max(this.x_max, point[0])
	this.y_max = max(this.y_max, point[1])
	this.x_min = min(this.x_min, point[0])
	this.y_min = min(this.y_min, point[1])
}

func (this *CountSquares) Count(point []int) int {
    if point[0] < this.x_min || point[0] > this.x_max || point[1] < this.y_min || point[1] > this.y_max {
		return 0
	}

	total := 0

	// left side
	max_square_len_left := point[0]-this.x_min
		// top
		max_square_len_left = min(max_square_len_left, this.y_max-point[1])
		for i:=1; i<=max_square_len_left; i++ {
			totalSquares := this.pointMap[Point{x: point[0]-i, y: point[1]}] * this.pointMap[Point{x:point[0]-i, y:point[1]+i}] * this.pointMap[Point{x:point[0], y: point[1]+i}]
			total += totalSquares
		}
		// bottom
		max_square_len_left = min(point[0]-this.x_min, point[1]-this.y_min)
		for i:=1; i<=max_square_len_left; i++ {
			totalSquares := this.pointMap[Point{x: point[0]-i, y: point[1]}] * this.pointMap[Point{x:point[0]-i, y:point[1]-i}] * this.pointMap[Point{x:point[0], y: point[1]-i}]
			total += totalSquares
		}

	// right side
	max_square_len_right := this.x_max-point[0]
		// top
		max_square_len_right = min(max_square_len_right, this.y_max-point[1])
		for i:=1; i<=max_square_len_right; i++ {
			totalSquares := this.pointMap[Point{x: point[0]+i, y: point[1]}] * this.pointMap[Point{x:point[0]+i, y:point[1]+i}] * this.pointMap[Point{x:point[0], y: point[1]+i}]
			total += totalSquares
		}
		// bottom
		max_square_len_right = min(this.x_max-point[0], point[1]-this.y_min)
		for i:=1; i<=max_square_len_right; i++ {
			totalSquares := this.pointMap[Point{x: point[0]+i, y: point[1]}] * this.pointMap[Point{x:point[0]+i, y:point[1]-i}] * this.pointMap[Point{x:point[0], y: point[1]-i}]
			total += totalSquares
		}

	return total
}
