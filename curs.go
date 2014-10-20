package main

import (
	"math"
    //"math/big"
    "strconv"
    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
    "runtime"
    "fmt"
    "os"
)

const (
	lns int = 20
)

type Rand struct {
	rnd  int64
	iter int64
	base int64
	l    int64
}

func NewGen(Base, L int64) *Rand {
	out := new(Rand)
	out.base = Base
	out.iter = 1
	out.l = L
	return out
}

func (rnd *Rand) slice() []int64 {
	base := rnd.base
	out := make([]int64, lns)
	temp_base := int64(1)
	for i := lns - 1; i >= 0; i-- {
		out[i] = temp_base
        //temp_base.Mul(temp_base,base)
		temp_base *= base
	}
	return out
}

func (rnd *Rand) bslice() []int64 {
	out := make([]int64, lns)
	temp := rnd.rnd
	base_t := rnd.slice()
	for i := 0; i < lns; i++ {
        //temp_s.Div(temp,base_t[i])
		temp_s := temp / base_t[i]
		out[i] = temp_s
		temp %= base_t[i]
        //temp.Mod(temp,base_t[i])
	}
	return out
}

func (rnd *Rand) Next() float64 {
	rnd.rnd = int64(math.Floor(float64(rnd.iter)) *
        math.Sqrt(float64(rnd.l)*float64(rnd.iter)))
	//rnd.rnd = rnd.iter * int(math.Floor(math.Sqrt(float64(rnd.base))))
	slice := rnd.bslice()
	out := float64(0)
	for i := 0; i < lns; i++ {
		out += float64(slice[lns-1-i]) *
        math.Pow(float64(rnd.base), float64(-i-1))
	}
	rnd.iter++
	return out
}

func Plot(Base,L,lens int64) error {
    pts := PointsArray(Base,L,lens)
    p, err := plot.New()
    if err != nil {
       return err
    }
    p.Title.Text = "Base: " + strconv.FormatInt(Base,10) +
        " L: " + strconv.FormatInt(L,10)
    p.X.Label.Text = "Count"
    p.Y.Label.Text = "Value"
    err = plotutil.AddLinePoints(p,"", pts,)
    if err != nil {
        return err
    }
    name := "B_" + strconv.FormatInt(Base,10) +
        "_L_" + strconv.FormatInt(L,10) + ".png"
    if err := p.Save(14, 5, name); err != nil {
            return err
    }
    fmt.Println("Finish " + strconv.FormatInt(Base,10) + " " +
        strconv.FormatInt(L,10))
    runtime.GC()
    return nil
}

func PointsArray(Base, L, lens int64) plotter.XYs {
        pts := make(plotter.XYs, lens)
        gen := NewGen(Base,L)
        for i := range pts {
                pts[i].X = float64(i)
                pts[i].Y = gen.Next()
        }
        return pts
}

func PlotFreqTest(Base, L, lens int64) error {
    pts := PointsArray(Base, L ,lens)
    value := make([]int,10)
    for i := range pts {
        temp := pts[i].Y
        switch {
        case temp < 0.1 && temp > 0:
            value[0]++
        case temp < 0.2 && temp >= 0.1:
            value[1]++
        case temp < 0.3 && temp >= 0.2:
            value[2]++
        case temp < 0.4 && temp >= 0.3:
            value[3]++
        case temp < 0.5 && temp >= 0.4:
            value[4]++
        case temp < 0.6 && temp >= 0.5:
            value[5]++
        case temp < 0.7 && temp >= 0.6:
            value[6]++
        case temp < 0.8 && temp >= 0.7:
            value[7]++
        case temp < 0.9 && temp >= 0.8:
            value[8]++
        case temp <= 1  && temp >= 0.9:
            value[9]++
        }
    }
    p, err := plot.New()
    if err != nil {
       return err
    }
    p.Title.Text = "Test Freq Base: " + strconv.FormatInt(Base,10) +
        " L: " + strconv.FormatInt(L,10)
    p.X.Label.Text = "Count"
    p.Y.Label.Text = "Value"
    err = plotutil.AddLinePoints(p,
            "", pts,
            strconv.FormatInt(int64(value[0]),10),line(lens,0.1),
            strconv.FormatInt(int64(value[1]),10),line(lens,0.2),
            strconv.FormatInt(int64(value[2]),10),line(lens,0.3),
            strconv.FormatInt(int64(value[3]),10),line(lens,0.4),
            strconv.FormatInt(int64(value[4]),10),line(lens,0.5),
            strconv.FormatInt(int64(value[5]),10),line(lens,0.6),
            strconv.FormatInt(int64(value[6]),10),line(lens,0.7),
            strconv.FormatInt(int64(value[7]),10),line(lens,0.8),
            strconv.FormatInt(int64(value[8]),10),line(lens,0.9),
            strconv.FormatInt(int64(value[9]),10),line(lens,1),
        )
    if err != nil {
        return err
    }
    name := "FT_B_" + strconv.FormatInt(Base,10) +
        "_L_" + strconv.FormatInt(L,10) + ".png"
    if err := p.Save(14, 5, name); err != nil {
            return err
    }
    fmt.Println("Finish Freq Test" + strconv.FormatInt(Base,10) + " " +
        strconv.FormatInt(L,10))
    runtime.GC()
    return nil

}

func line(lens int64, level float64) plotter.XYs {
    pts := make(plotter.XYs,lens)
    for i := range pts {
        pts[i].X = float64(i)
        pts[i].Y = level

    }
    return pts
}

func Test(Base,L,lens int64) {
    Plot(Base,L,lens)
    PlotFreqTest(Base,L,lens)
}

func main() {
    arg := os.Args
    base,_ := strconv.ParseInt(arg[1],10,0)
    l,_ := strconv.ParseInt(arg[2],10,0)
    Test(base,l,20000)
}
