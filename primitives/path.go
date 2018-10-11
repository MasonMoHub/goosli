package primitives

type Path struct {
	Lines []Line
}

func (p Path) Reverse() Path {
	res := Path{}
	for _,l := range p.Lines {
		res.Lines = append(res.Lines, l.Reverse())
	}
	return res
}

func (p Path) Enclose() Path {
	s := len(p.Lines)
	if !p.Lines[s-1].P2.Equal(p.Lines[0].P1) {
		p.Lines = append(p.Lines, Line{p.Lines[s-1].P2, p.Lines[0].P1} )
	}
	return p
}

func JoinPaths(paths []Path) []Path {
	lookup := make(map[Point]Path, len(paths))
	for _, path := range paths {
		lookup[path.Lines[0].P1] = path
	}
	var result []Path
	for len(lookup) > 0 {
		var v Point
		for v = range lookup {
			break
		}
		var path Path
		for {
			if p, ok := lookup[v]; ok {
				path.Lines = append(path.Lines, p.Lines[0])
				delete(lookup, v)
				v = p.Lines[len(p.Lines)-1].P2
			} else {
				break
			}
		}
		result = append(result, path)
	}
	for i:=0; i<len(result);i++{
		for j:=i+1;j<len(result);j++ {
			jp := tryJoin(result[i], result[j])
			if jp != nil{
				result[i] = *jp
				result[j] = result[len(result)-1]
				result = result[:len(result)-1]
				i=-1
				break
			}
		}
	}
	// clean from small lines
	for i:=0;i< len(result);i++{
		for j:=0;j< len(result[i].Lines);j++ {
			if result[i].Lines[j].ToVector().Length() < 0.1 { //TODO: hardcode
				if j!=0 {
					result[i].Lines[j-1].P2 = result[i].Lines[j].P2
				}else if j<len(result[i].Lines)-1 {
					result[i].Lines[j+1].P1 = result[i].Lines[j].P1
				}
				result[i].Lines = append(result[i].Lines[:j] , result[i].Lines[j+1:]...)
				j--
			}
		}
		if len(result[i].Lines) ==0{
			result[i] = result[len(result)-1]
			result = result[:len(result)-1]
			i--
		}
	}
	return result
}

func tryJoin(p1, p2 Path) *Path {
	if p1.Lines[len(p1.Lines)-1].P2.Equal(p2.Lines[0].P1){
		p1.Lines = append(p1.Lines, p2.Lines...)
		//println("works right way")
		return &p1
	}
	if p2.Lines[len(p2.Lines)-1].P2.Equal(p1.Lines[0].P1){
		p2.Lines = append(p2.Lines, p1.Lines...)
		//println("works otherside")
		return &p2
	}
	if p1.Lines[len(p1.Lines)-1].P2.Equal(p2.Lines[len(p2.Lines)-1].P2){
		p1.Lines = append(p1.Lines, p2.Reverse().Lines...)
		println("something not good, reverse path found")
		return &p1
	}
	if p1.Lines[0].P1.Equal(p2.Lines[0].P1){
		p2.Lines = append(p2.Reverse().Lines, p1.Lines...)
		println("something not good, reverse path found")
		return &p2
	}
	return nil
}

func FindCentroid(path Path) Point { //TODO: refactorme
	a := 0.0
	a2 := 0.0
	a3 := 0.0
	cx := 0.0
	cy := 0.0
	cy2 := 0.0
	cz2 := 0.0
	cx2 := 0.0
	cz := 0.0
	for _, line := range path.Lines {
		p := line.P1
		p_1 := line.P2
		d := p.X*p_1.Y - p_1.X*p.Y
		a += d
		cx += (p.X + p_1.X) * d
		cy += (p.Y + p_1.Y) * d

		d2 := p.X*p_1.Z - p_1.X*p.Z
		a2 += d2
		cx2 += (p.X + p_1.X) * d2
		cz += (p.Z + p_1.Z) * d2

		d3 := p.Y*p_1.Z - p_1.Y*p.Z
		a3 += d3
		cy2 += (p.Y + p_1.Y) * d3
		cz2 += (p.Z + p_1.Z) * d3
	}
	var x, x1, y, y1, z, z1 float64
	//println(a, a2, a3, "<>", cx, cx2, cy, cy2, cz, cz2, "\n")
	a = a * 0.5
	if !AlmostZero(a){
		x = cx / 6 / a
		y = cy / 6 / a
	}
	a2 = a2 * 0.5
	if !AlmostZero(a2) {
		x1 = cx2 / 6 / a2
		z = cz / 6 / a2
	}
	a3 = a3 * 0.5
	if !AlmostZero(a3) {
		y1 = cy2 / 6 / a3
		z1 = cz2 / 6 / a3
	}

	//println("x:", x, " ", x1)
	//println("y:", y, " ", y1)
	//println("z:", z, " ", z1)

	if AlmostZero(a) && AlmostZero(a2){
		x = path.Lines[0].P1.X
		y = y1
		z = z1
	}
	if AlmostZero(a) && AlmostZero(a3) {
		y = path.Lines[0].P1.Y
		x = x1
	}
	if AlmostZero(a2) && AlmostZero(a3) {
		z = path.Lines[0].P1.Z
	}

	if AlmostZero(a) && !AlmostZero(a2){
		x = x1
	}
	if AlmostZero(a) && !AlmostZero(a3) {
		y = y1
	}
	if AlmostZero(a2) && !AlmostZero(a3) {
		z = z1
	}
	//println(x,x1,y,y1,z,z1,"\n")
	return Point{x, y, z}
}
