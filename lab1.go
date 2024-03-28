package main

import (
	"fmt"
	"github.com/benoitmasson/plotters/piechart"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math/rand/v2"
)

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
	task11()
	task12()
}

func task1() {
	values := plotter.Values{1379, 1300, 1250, 1210, 1171, 1120, 1076, 1030, 998, 940}
	years := []string{"2024", "2025", "2026", "2027", "2028", "2029", "2030", "2031", "2032", "2033"}

	p := plot.New()
	p.Title.Text = "Прогноз щодо кількості навчальних закладів у наступні десять років"
	p.X.Label.Text = "Рік"
	p.Y.Label.Text = "Кількість навчальних закладів"

	w := vg.Points(40)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}
	barsA.Color = plotutil.Color(0)

	p.Add(barsA)
	p.NominalX(years...)
	p.Y.Min = 0
	p.Y.Max = 1600
	p.X.Min = 0
	p.X.Max = 10
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "task1.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task2() {
	values := plotter.Values{10, 100, 65, 30, 15}
	years := []string{"5-10тис.", "10-15тис.", "15-20тис.", "25-30тис.", "30тис. і вище"}

	p := plot.New()
	p.Title.Text = "Розподіл заробітної плати працівників компанії"
	p.X.Label.Text = "Розмір зарплати"
	p.Y.Label.Text = "Кількість співробітників"

	w := vg.Points(40)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}
	barsA.Color = plotutil.Color(4)

	p.Add(barsA)
	p.NominalX(years...)
	p.Y.Min = 0
	p.Y.Max = 120
	p.X.Min = 0
	p.X.Max = 5
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "task2.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task3() {
	fuelTypes := []string{"А-92", "А-95", "Дизель"}
	prices := [][]float64{
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5},
	}
	qualities := [][]float64{
		{0, 3, 2, 4, 3, 2},
		{0, 4, 3, 4, 4, 3},
		{0, 2, 2, 3, 3, 2},
	}

	p := plot.New()

	p.Title.Text = "Якість палива за одиницю ціни"
	p.X.Label.Text = "Одиниця ціни"
	p.Y.Label.Text = "Якість палива (від 1 до 5)"

	for i, fuelType := range fuelTypes {
		points := make(plotter.XYs, len(prices[i]))
		for j := range prices[i] {
			points[j].X = prices[i][j]
			points[j].Y = qualities[i][j]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			fmt.Println("Error creating line plot:", err)
			return
		}

		line.Color = plotutil.Color(i)
		line.Width = vg.Points(3)
		p.Add(line)

		p.Legend.Add(fuelType, line)
	}
	p.Y.Min = 0
	p.Y.Max = 5
	p.X.Min = 0
	p.X.Max = 5.5
	p.Legend.Top = true
	p.Add(plotter.NewGrid())

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "task3.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task4() {
	units := []string{"Підрозділ 1", "Підрозділ 2", "Підрозділ 3", "Підрозділ 4", "Підрозділ 5", "Підрозділ 6"}
	values := plotter.Values{3.3, 3.5, 3.2, 3.11, 3.27, 3.4}

	p := plot.New()

	p.Title.Text = "Рівень плинності кадрів у підрозділах у вересні"
	p.X.Label.Text = "Підрозділ"
	p.Y.Label.Text = "Рівень плинності кадрів (у відсотках)"

	w := vg.Points(40)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}
	barsA.Color = plotutil.Color(2)

	p.Add(barsA)
	p.NominalX(units...)
	p.Y.Min = 0
	p.Y.Max = 5
	p.X.Min = 0
	p.X.Max = 5.4

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "task4.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task5() {
	p := plot.New()
	p.HideAxes()
	responsibilities := []string{"Продаж товару", "Інвентаризація", "Робота з клієнтами", "Підведення звітів", "Обід"}
	percentages := []float64{30, 20, 15, 25, 10}
	colors := []color.Color{color.RGBA{11, 255, 130, 255}, color.RGBA{255, 200, 100, 255}, color.RGBA{R: 140, G: 50, B: 100, A: 255}, color.RGBA{R: 30, G: 10, B: 210, A: 255}, color.RGBA{R: 33, G: 161, B: 13, A: 255}}
	offset := 0.0
	for i := range percentages {
		pie, err := piechart.NewPieChart(plotter.Values{percentages[i]})
		if err != nil {
			panic(err)
		}
		pie.Color = colors[i]
		pie.Total = 100
		pie.Labels.Nominal = []string{responsibilities[i]}
		pie.Labels.Values.Show = true
		pie.Offset.Value = offset
		pie.Labels.Values.Percentage = true
		pie.Labels.Font.Size = vg.Points(10)
		p.Add(pie)
		offset += percentages[i]
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "task5.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
		return
	}
}

func task6() {
	values := plotter.Values{5, 5, 5, 5, 5}
	years := []string{"менше 3-х років", "3-5 років", "5-10 років", "10-15 років", "15 років і довше"}

	p := plot.New()
	p.Title.Text = "Надбавка до зарплатні з часом роботи"
	p.X.Label.Text = "Робочий стаж"
	p.Y.Label.Text = "Надбавка до зарплатні (у відсотках)"

	for i := range years {
		points := make(plotter.XYs, len(years))
		for i := range values {
			points[i].X = float64(i)
			points[i].Y = values[i]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			fmt.Println("Error creating line plot:", err)
			return
		}

		line.Color = plotutil.Color(i)
		line.Width = vg.Points(3)
		p.NominalX(years...)
		p.Add(line)
	}
	p.Y.Min = 0
	p.Y.Max = 20
	p.X.Min = 0
	p.X.Max = 3
	p.Add(plotter.NewGrid())

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "task6.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task7() {
	ageGroup := []string{"18-25 років", "25-30 років", "30-35 років", "35+ років"}
	values := plotter.Values{3.5, 2.7, 12, 5}

	p := plot.New()

	p.Title.Text = "Рівень плинності кадрів у вікових групах"
	p.X.Label.Text = "Вік"
	p.Y.Label.Text = "Плинність кадрів (у відсотках)"

	w := vg.Points(40)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}
	barsA.Color = plotutil.Color(1)

	p.Add(barsA)
	p.NominalX(ageGroup...)
	p.Y.Min = 0
	p.Y.Max = 30
	p.X.Min = 0
	p.X.Max = 3.5

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "task7.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task8() {
	values := plotter.Values{34, 27, 25, 21, 17}
	regions := []string{"Північ", "Захід", "Південь", "Схід", "Центр"}

	p := plot.New()

	p.Title.Text = "Народжуваність у регіонах країни"
	p.X.Label.Text = "Народжуваність на 1000 людей"
	p.Y.Label.Text = "Регіон"
	barChart, err := plotter.NewBarChart(values, vg.Points(40))
	if err != nil {
		log.Fatalf("Error creating horizontal bar chart: %v", err)
	}

	barChart.Horizontal = true
	barChart.Color = color.RGBA{R: 0, G: 128, B: 255, A: 255}
	p.NominalY(regions...)
	p.Add(barChart)

	p.X.Max = 60
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "task8.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}
}

func task9() {
	values := plotter.Values{5, 7, 9, 3}
	years := []string{"Січень", "Лютий", "Березень", "Квітень"}

	p := plot.New()
	p.Title.Text = "Прибутковість акцій компанії"
	p.X.Label.Text = "Місяць"
	p.Y.Label.Text = "Прибутковість (у відсотках)"

	w := vg.Points(50)

	barsA, err := plotter.NewBarChart(values, w)
	if err != nil {
		panic(err)
	}
	barsA.Color = plotutil.Color(0)

	p.Add(barsA)
	p.NominalX(years...)
	p.Y.Min = 0
	p.Y.Max = 20
	p.X.Min = 0
	p.X.Max = 3.5

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "task9.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}

func task10() {
	p := plot.New()
	p.Title.Text = "Розподіл діяльності фондів"
	p.HideAxes()
	responsibilities := []string{"Виробництво", "Фінансовий сектор", "Будівництво", "С/Г",
		"Логістика", "Енергетика"}
	percentages := []float64{40, 20, 15, 10, 8, 7}
	colors := []color.RGBA{
		{R: 210, G: 210, B: 210, A: 255},
		{R: 255, G: 255, B: 204, A: 255},
		{R: 204, G: 255, B: 204, A: 255},
		{R: 204, G: 229, B: 255, A: 255},
		{R: 255, G: 204, B: 255, A: 255},
		{R: 255, G: 204, B: 204, A: 255},
	}
	offset := 0.0
	for i := range percentages {
		pie, err := piechart.NewPieChart(plotter.Values{percentages[i]})
		if err != nil {
			panic(err)
		}
		pie.Color = colors[i]
		pie.Total = 100
		pie.Labels.Nominal = []string{responsibilities[i]}
		pie.Labels.Values.Show = true
		pie.Offset.Value = offset
		pie.Labels.Values.Percentage = true
		pie.Labels.Font.Size = vg.Points(10)
		p.Add(pie)
		offset += percentages[i]
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "task10.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
		return
	}
}

func task11() {
	incomes := generateRandomData(100, 50000)

	p := plot.New()

	s, err := plotter.NewScatter(incomes)
	if err != nil {
		log.Fatalf("Error creating scatter plot: %v", err)
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	p.Add(s)

	p.Title.Text = "Взаємозв'язок доходу і зарплати"
	p.X.Label.Text = "Зарплата"
	p.Y.Label.Text = "Дохід"

	p.X.Max = 52000
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "task11.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}
}

func generateRandomData(n int, max float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		salary := rand.Float64() * max

		income := salary*0.5 + rand.Float64()*10000

		pts[i].X = salary
		pts[i].Y = income
	}
	return pts
}

func task12() {
	units := []string{"Студент 1", "Студент 2", "Студент 3", "Студент 4", "Студент 5", "Студент 6", "Студент 7", "Студент 8"}
	month := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	monthName := []string{"Січень", "Лютий", "Березень", "Квітень", "Травень", "Червень", "Липень", "Серпень",
		"Вересень", "Жовтень", "Листопад", "Грудень"}

	averageGrade := [][]float64{
		{9.5, 9.3, 9.3, 9.4, 9.0, 9.3, 9.2, 9.10, 8.8, 8.9, 9.0, 9.1},
		{8.8, 8.7, 8.7, 8.65, 8.53, 8.65, 8.8, 8.10, 8.2, 8.3, 8.3, 8.25},
		{7.95, 8.25, 8.3, 8.4, 8.8, 8.95, 8.9, 8.8, 8.8, 8.8, 8.85, 8.9},
		{9.2, 9.2, 9.4, 9.7, 9.8, 9.55, 9.15, 9.00, 9.3, 9.5, 9.4, 9.4},
		{6.2, 6.6, 6.9, 7.55, 8.20, 8.55, 8.9, 8.7, 8.8, 8.9, 8.65, 8.5},
		{8.8, 8.8, 8.2, 8.05, 8.15, 8.12, 8.20, 8.3, 8.4, 8.9, 9.0, 9.1},
		{7.3, 7.2, 7.4, 7.6, 8.0, 8.4, 8.8, 9.20, 9.4, 9.4, 9.5, 9.6},
		{7.8, 8.3, 8.55, 8.4, 8.6, 8.8, 8.75, 9.45, 9.15, 9.2, 9.3, 9.4},
	}

	p := plot.New()

	p.Title.Text = "Успішність студентів протягом року"
	p.X.Label.Text = "Місяць"
	p.Y.Label.Text = "Середній бал (у десятибальній шкалі)"

	for i, student := range units {
		points := make(plotter.XYs, len(month))
		for j := range month {
			points[j].X = month[j]
			points[j].Y = averageGrade[i][j]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			fmt.Println("Error creating line plot:", err)
			return
		}

		line.Color = plotutil.Color(i)
		line.Width = vg.Points(3)
		p.Add(line)

		p.Legend.Add(student, line)
	}
	p.Y.Min = 6
	p.Y.Max = 10
	p.X.Min = 0
	p.X.Max = 12.5
	p.NominalX(monthName...)
	p.Legend.Top = true
	p.Add(plotter.NewGrid())

	if err := p.Save(12*vg.Inch, 6*vg.Inch, "task12.png"); err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
}
