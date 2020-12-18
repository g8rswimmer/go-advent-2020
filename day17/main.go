package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	curr3d := entries()

	for cycle := 0; cycle <= 6; cycle++ {
		fmt.Printf("3D Cycle: %d Actives %d\n", cycle, curr3d.actives())

		next := newDimension()
		for z := curr3d.layouts[2].min - 1; z <= curr3d.layouts[2].max+1; z++ {
			for y := curr3d.layouts[1].min - 1; y <= curr3d.layouts[1].max+1; y++ {
				for x := curr3d.layouts[0].min - 1; x <= curr3d.layouts[0].max+1; x++ {
					an := curr3d.threedneighbors(x, y, z)
					active := curr3d.get(x, y, z, 0)
					switch {
					case active && (an != 2 && an != 3):
						next.set(x, y, z, 0, false)
					case active == false && an == 3:
						next.set(x, y, z, 0, true)
					default:
						next.set(x, y, z, 0, active)
					}
				}
			}
		}
		curr3d = next
	}
	fmt.Println()

	curr4d := entries()

	for cycle := 0; cycle <= 6; cycle++ {
		fmt.Printf("4D Cycle: %d Actives %d\n", cycle, curr4d.actives())

		next := newDimension()
		for w := curr4d.layouts[3].min - 1; w <= curr4d.layouts[3].max+1; w++ {
			for z := curr4d.layouts[2].min - 1; z <= curr4d.layouts[2].max+1; z++ {
				for y := curr4d.layouts[1].min - 1; y <= curr4d.layouts[1].max+1; y++ {
					for x := curr4d.layouts[0].min - 1; x <= curr4d.layouts[0].max+1; x++ {
						an := curr4d.fourdneighbors(x, y, z, w)
						active := curr4d.get(x, y, z, w)
						switch {
						case active && (an != 2 && an != 3):
							next.set(x, y, z, w, false)
						case active == false && an == 3:
							next.set(x, y, z, w, true)
						default:
							next.set(x, y, z, w, active)
						}
					}
				}
			}
		}
		curr4d = next
	}

}

type layout struct {
	min int
	max int
}

type dimension struct {
	coordinates map[int]map[int]map[int]map[int]bool
	layouts     []layout
}

func (d *dimension) set(x, y, z, w int, active bool) {
	yplane, xhas := d.coordinates[x]
	if xhas == false {
		yplane = map[int]map[int]map[int]bool{}
	}
	zplane, yhas := yplane[y]
	if yhas == false {
		zplane = map[int]map[int]bool{}
	}
	wplane, zhas := zplane[z]
	if zhas == false {
		wplane = map[int]bool{}
	}
	wplane[w] = active
	zplane[z] = wplane
	yplane[y] = zplane
	d.coordinates[x] = yplane

	d.layouts[0].min = min(d.layouts[0].min, x)
	d.layouts[0].max = max(d.layouts[0].max, x)

	d.layouts[1].min = min(d.layouts[1].min, y)
	d.layouts[1].max = max(d.layouts[1].max, y)

	d.layouts[2].min = min(d.layouts[2].min, z)
	d.layouts[2].max = max(d.layouts[2].max, z)

	d.layouts[3].min = min(d.layouts[3].min, w)
	d.layouts[3].max = max(d.layouts[3].max, w)
}

func (d dimension) get(x, y, z, w int) bool {
	yplane, xhas := d.coordinates[x]
	if xhas == false {
		return false
	}
	zplane, yhas := yplane[y]
	if yhas == false {
		return false
	}
	wphane, xhas := zplane[z]
	if xhas == false {
		return false
	}
	return wphane[w]
}

func (d dimension) actives() int {
	result := 0
	for _, yplane := range d.coordinates {
		for _, zplane := range yplane {
			for _, wplane := range zplane {
				for _, active := range wplane {
					if active {
						result++
					}
				}
			}
		}
	}
	return result
}
func (d dimension) xneighbors(x, y, z, w int) int {
	active := 0
	if d.get(x, y, z, w) {
		active++
	}
	if d.get(x-1, y, z, w) {
		active++
	}
	if d.get(x+1, y, z, w) {
		active++
	}
	return active
}

func (d dimension) yneighbors(x, y, z, w int) int {
	active := 0
	active += d.xneighbors(x, y, z, w)
	active += d.xneighbors(x, y-1, z, w)
	active += d.xneighbors(x, y+1, z, w)
	return active
}

func (d dimension) zneighbors(x, y, z, w int) int {
	active := 0
	active += d.yneighbors(x, y, z, w)
	active += d.yneighbors(x, y, z+1, w)
	active += d.yneighbors(x, y, z-1, w)
	return active
}

func (d dimension) threedneighbors(x, y, z int) int {
	active := d.zneighbors(x, y, z, 0)
	if d.get(x, y, z, 0) {
		active--
	}
	return active
}

func (d dimension) fourdneighbors(x, y, z, w int) int {
	active := 0
	active += d.zneighbors(x, y, z, w)
	active += d.zneighbors(x, y, z, w+1)
	active += d.zneighbors(x, y, z, w-1)
	if d.get(x, y, z, w) {
		active--
	}
	return active
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newDimension() *dimension {
	return &dimension{
		coordinates: map[int]map[int]map[int]map[int]bool{},
		layouts:     make([]layout, 4),
	}
}

func entries() *dimension {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	dimen := newDimension()
	y, z, w := 0, 0, 0
	for scan.Scan() {
		line := scan.Text()
		for x := 0; x < len(line); x++ {
			active := false
			if line[x] == '#' {
				active = true
			}
			dimen.set(x, y, z, w, active)
		}
		y++
	}
	return dimen
}
