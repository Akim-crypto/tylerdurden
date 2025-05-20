package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)



func main() {
	village := Village{}

	// Создаем жителей деревни
	resident1 := &Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	village.AddElement(animal1)
	village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}



type VillageElement interface {
	Update()
	FlushInfo() string
}


type Village struct {
	Elements []VillageElement
}


func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}


func (v *Village) UpdateAll() {
	for _ , e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _ , e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}

type Resident struct {
	Name string
	Age int
	Married bool
	Alive bool
	Events []string
}


func (r *Resident) addYear(){
	r.Age++
}

func (r *Resident) changeMarriedStatus(){
	if r.Married {
		r.Married = false
		r.Events = append(r.Events, "Развод , больше я не в браке")
	} else {
		r.Married = true
		r.Events = append(r.Events, "Наконец то , спутник жизни найден")
	}
}

func (r *Resident) die(){
	r.Alive = false
	r.Events = append(r.Events, fmt.Sprintf("Ушел на покой на %d , году жизни",r.Age))
}


func (r *Resident) Update(){
	if !r.Alive{
		return
	}
	r.addYear()
	if rand.IntN(100) < 15 {
		r.changeMarriedStatus()
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Нашел новую работу")
	}
	if r.Married && rand.IntN(100) < 25 {
		r.Events = append(r.Events, "Поруглался с супргуой/ом")
	}

	if rand.IntN(100) < 5 {
		r.die()
	}

}

func (r *Resident)FlushInfo() string {
	info := fmt.Sprintf("Житель %s умер в возрасте %d.",r.Name,r.Age)
	if r.Alive {
		marriedStatus := "холост"
		if r.Married {
			marriedStatus = "в браке"
		}
		events := "нет"
		if len(events) > 0 {
			events = strings.Join(r.Events, "\n")
		}
		info = fmt.Sprintf("Житель %s (возраст %d), статус: %s\nСобытия:%s:\n",r.Name,r.Age,marriedStatus,events)

	}
	r.Events = []string{}
	return info
}



type Animal struct {
	Name string
	Age int
	Type string
	Alive bool
	Events []string
}


func (r *Animal) addYear(){
	r.Age++
}


func (a *Animal) die(){
	a.Alive = false
	a.Events = append(a.Events, fmt.Sprintf("Умер на %d  году жизни",a.Age))
}



func (a *Animal) Update(){
	if !a.Alive{
		return
	}
	a.addYear()
	if rand.IntN(100) < 7 {
		a.Events = append(a.Events, "Сломал лапу")
	}
	if  a.Type == "собака" && rand.IntN(100) < 25 {
		a.Events = append(a.Events, "Покусал прохожего")
	}
	if  a.Type == "кошка" && rand.IntN(100) < 25 {
		a.Events = append(a.Events, "Убежала из дома")

	if rand.IntN(100) < 7 {
		a.die()
	}
	}
}

func (r *Animal)FlushInfo() string {
	info := fmt.Sprintf("Житель %s умер в возрасте %d.",r.Name,r.Age)
	if r.Alive {
		
		events := "нет"
		if len(events) > 0 {
			events = strings.Join(r.Events, "\n")
		}
		info = fmt.Sprintf("Животное %s (возраст %d),\nСобытия:%s:\n",r.Name,r.Age,events)

	}
	r.Events = []string{}
	return info
}
