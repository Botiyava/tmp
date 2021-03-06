package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Sample struct {
	x     []int
	x_len int
	x_min int
	x_max int
	x_sum int
}

func (s *Sample) Init() (err error) {
	if len(s.x) >= 2 {
		sort.Ints(s.x)
		s.x_len = len(s.x)
		s.x_min = s.x[0]
		s.x_max = s.x[s.x_len-1]

		for _, elem := range s.x {
			s.x_sum += elem
		}

		return nil
	} else {
		return errors.New("выборка должна содержать минимум два значения")
	}

}

//Экспоненциальное распределение
func (s *Sample) expDistribution(x float64) float64 {
	var lambda float64
	lambda = float64(s.x_len) / float64(s.x_sum)
	return lambda * math.Exp(-lambda*x)

}

//Распределение Релея
func (s *Sample) ReleyDistribution(x float64) float64 {
	M := s.x_sum / s.x_len
	sigma := float64(M) / 1.253
	return (x / (sigma * sigma)) * math.Exp((-x*x)/(2*sigma*sigma))
}

//Правила разбиения интервалов
func intervalRule(rule int, N int) (n int) {
	switch rule {
	case 1:
		n = int(1 + 3.3*math.Log10(float64(N)))
	case 2:
		n = int(5 * math.Log10(float64(N)))
	case 3:
		n = int(math.Pow(float64(N), 0.5))
	case 4:
		n = int(math.Pow(float64(N), 0.33333333))
	}
	return
}

//Критерий Хи квадрат
func (s *Sample) Khi(f func(float64) float64, n int) {
	delta := (s.x_max - s.x_min) / n
	currentDelta := delta

	// Разбиение слайса на промежутки (количество промежутков определяется правилом n)
	intervals := make([][]int, n+1)
	j := 0
	for i := 0; i < s.x_len; i++ {
		if s.x[i] < currentDelta {
			intervals[j] = append(intervals[j], s.x[i])
		} else {
			if i == s.x_len-1 {
				intervals[j] = append(intervals[j], s.x[i])
				break
			}
			j++
			intervals[j] = append(intervals[j], s.x[i])
			currentDelta += delta
		}
	}

	//второй столбец таблицы, количество элементов в каждом интервале
	mi := make([]int, n)
	for i := 0; i < n; i++ {
		mi[i] = len(intervals[i])
	}
	delta = (s.x_max - s.x_min) / n

	//третий столбец, вероятность в зависимости от распределения. Интегрирование
	pi := make([]float64, n)
	pi_origin := make([]float64, n)
	iteration := 0
	for i := 1; i <= n; i++ {
		for j := 0; j < delta*i; j++ {
			pi[iteration] += f(float64(j))
			pi_origin[iteration] += f(float64(j))
		}
		if i != 1 && i != n {
			pi[iteration] -= pi_origin[iteration-1]
		}
		if i == n {
			tmp := 0.0
			for i := 0; i < n-1; i++ {
				tmp += pi[i]
			}
			pi[iteration] = 1 - tmp
		}
		iteration++
	}

	//4 и 5 столбец
	Np := make([]float64, n)
	mNp2 := make([]float64, n)
	khi := 0.0
	for i := 0; i < n; i++ {
		Np[i] = pi[i] * float64(s.x_len)
		mNp2[i] = math.Pow(float64(mi[i])-Np[i], 2) / Np[i]
		khi += mNp2[i]

	}
	fmt.Printf("mi:\t%v\n\np_i:\t%v\n\nNp_j:\t%v\n\n(mNp_j)^2/Np_j:\t%v\nkhi: %v", mi, pi, Np, mNp2, khi)
}

//Критерий Колмогорова
func (s *Sample) Kolmogorov(f func(float64) float64) float64 {
	res := 0.0
	for i := 0; i < s.x_len; i++ {
		res += f(float64(i)) - float64(i)/float64(s.x_max)
	}
	return res
}
func main() {
	//Исходные данные
	sample := Sample{x: []int{5, 45, 35, 88, 72, 5, 16, 46, 26, 65, 76, 29, 43, 10, 9, 43, 69, 21, 32, 15, 81, 55, 48, 47, 50, 13,
		77, 63, 76, 20, 9, 95, 62, 76, 31, 46, 69, 92, 62, 6}}

	if err := sample.Init(); err != nil {
		panic(err)
	}
	n := intervalRule(4, sample.x_len)                       // Выбор правила
	sample.Khi(sample.ReleyDistribution, n)                  //Критерий Хи. Правила по фурнкции f по правилу n
	fmt.Println(sample.Kolmogorov(sample.ReleyDistribution)) // Критерий  Колмогорова по функции f

}
