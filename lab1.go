package main

import (
	"fmt"
	"sort"
	"strconv"


)

func main(){
	rawData := []float64{68.63,823.26,353.06,217.66,216.60,821.00,160.98,863.55,509.23,
		378.66,360.73,158.75,972.65,880.97,678.66,665.12,502.98,697.36,369.11,86.56,896.68,
		880.63,198.57,392.96,999.92,693.17,592.69,208.61,159.68,198.16,803.65,166.05,900.65,
		236.58,703.73,193.31,171.29,725.58,51.16,736.68,762.61,129.03,681.15,666.78,332.53,561.02,
		506.38,370.17,953.08,727.69,791.82,669.15,1007.71,876.75,657.56,238.68,108.67,900.38,291.09,
		1010.00,909.72,885.67,159.00,331.66,759.68,676.63,858.18,1007.07,673.35,828.99,990.06,262.09,
		563.13,516.11,68.65,856.18,855.25,875.26,560.16,999.99,500.60,169.28,261.31,210.52,239.98,680.73,
		238.07,108.20,660.56,1018.06,766.97,822.96,332.20,268.73,59.71,361.09,131.51, 138.15 ,697.11, 650.05}
	N := len(rawData)
	var sum float64
	sort.Float64s(rawData)
	for i:= 0; i < N; i++{

}
numberOfIntervals := 11
interval := 0  //Текущий интервал времени
gap := rawData[len(rawData)]/float64(numberOfIntervals)
fmt.Println(gap)
//gap := 100.0   // разница по времени  между двумя интервалами (в часах)
overallR := 0 // Общее количество вышедших из строя объектов

list := make(map[int][]float64) //списки объектов по интервалам

	for i:= 0; i < N; i++{
		sum += rawData[i]
		this:
		if rawData[i] < gap{
list[interval] = append(list[interval],rawData[i])
		}else{
			overallR += len(list[interval])
			// Δr - количество отказов в данных интервал времени, r - общее количество отказов за интервал времени
			// от начала работы до текущего времени включительно.
			fmt.Print("Δr(" + strconv.Itoa(interval) + ") = ")
			fmt.Print(len(list[interval]))
			fmt.Print("\tr(" + strconv.Itoa(interval) + ") = ")
			fmt.Println(overallR)

			interval += 100
			gap += 100.0
			goto this
		}

	}
	// Δr - количество отказов в данных интервал времени, r - общее количество отказов за интервал времени
	// от начала работы до текущего времени включительно. ( Для последнего интервала)
	overallR += len(list[interval])
	fmt.Print("Δr(" + strconv.Itoa(interval) + ") = ")
	fmt.Print(len(list[interval]))
	fmt.Print("\tr(" + strconv.Itoa(interval) + ") = ")
	fmt.Println(overallR)
	interval = 0
	overallR = 0
	gap = 100.0


//--------------         Рассчёт наработки на отказ

T := sum/float64(len(rawData))
	formatT:= fmt.Sprintf("%.4f", T)
	fmt.Println("\n---------------------------------------------------------------------------------------------")
fmt.Print("|  Интервалы\t|Кол-во отказов\t|Инт-ть отказов  | Част отказов |Ф-ия надежности| λ = p/α   |")
for i:= 0; i< len(list)+1; i++ {

	fmt.Println("\n---------------------------------------------------------------------------------------------")


	// --------------   Вероятность безотказной работы (Пятый столбец вывода)

	P := float64(1.0 - (float64(overallR)/float64(N)))
	formatP:= fmt.Sprintf("%.2f", P) //Форматирование вывода (Сам вывод в конце)


	//---------------   Текущий интервал времени (Первый столбец в таблице)
	fmt.Print("| (",interval, ",", interval + 100, ")\t" )

	// --------------   Количество отказов  (Второй столбец в таблице)
	fmt.Print("|\t ", len(list[interval]),"\t" )


	// --------------   Интенсивность отказов (Третий столбец в таблице)
	CurrentN := N - overallR
	overallR += len(list[interval])
	lambda := float64(len(list[interval])) / (((float64(CurrentN) + float64(N-overallR)) / 2) * gap)
	formatLambda := fmt.Sprintf("%.5f", lambda) //Форматирование вывода
	fmt.Print("|\t ",formatLambda, " |\t")



	// --------------   Частота отказа (Пятый столбец в таблице)
	frequency := float64(len(list[interval])) / (float64(N) * gap)
	fmt.Print(frequency, "\t| ")


	//// --------------  Контроль (Шестой столбец в таблице)
	fmt.Print(formatP, "\t\t| ") //Печать вероятности безотказной работы
	control := frequency/P
	formatControl := fmt.Sprintf("%.5f", control)
	fmt.Print(formatControl, "   | ")

interval += 100 //Переходим к рассчету количественных характеристик надежности для следующего интервала


}
	fmt.Println("\n---------------------------------------------------------------------------------------------")
fmt.Println("\nНаработка на отказ = ", formatT , " часов")


}
