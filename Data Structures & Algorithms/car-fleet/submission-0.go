// 同じタイミングでなくても早く着いていればOK、遅れてついたものだけを別のfleetとみなす
func carFleet(target int, position []int, speed []int) int {
	type Car struct {
		position int
		speed int
	}

	cars := make([]Car, len(position))
	for i := range position {
		cars[i] = Car{
			position: position[i],
			speed: speed[i],
		}
	}

	// positionを降順で並び替える
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})
	
	stack := []float64{}
	for _, car := range cars {
		time := float64(target - car.position) / float64(car.speed) 
		if len(stack) == 0 || time > stack[len(stack) - 1] {
			stack = append(stack, time)
		}
	}
	return len(stack)
}
