package main

import (
	"fmt"
	"math"

	"github.com/kzhendarov/GoExample/internal/class"
	"github.com/kzhendarov/GoExample/internal/constants"
)

/* Область видимости всего пакета */
// Объявление переменных
var myVar int

// Объявление с инициализацией
var maVar2 float32 = 2.581

// Объявление переменной по типу присваемого значения:
var myVar3 = true

func main() {

	// Объявление с определением типа по присваемому занчению:
	p := math.Pi // объявит переменную p float со значением константы 3.1415

	// Форматированный вывод (описание шаблона формата в README.md)
	fmt.Printf("myVar = %d\nmyVar2 = %.3f\nmyVar3 = %t\nCONST Pi = %.9f\n\n", myVar, maVar2, myVar3, p)
	// Пример использования перечисления (поля данных и методы типа Direction описаны в ".\internal\constants\directions.go")
	var moving constants.Direction = constants.North

	// Вывод на печать элемента перечисления и выбора направления на его основе
	fmt.Println(moving, moving.DirectionSelect())

	// Создание экземпляра типа Person (поля данных и методы типа Person описаны в ".\internal\class\person.go")

	var profile class.Person = *class.NewPerson("Joe", "Black", 32, "j.black@e-mail.com", "+0(8956)999-00-01")

	// Вывод информации с данными экземпляра типа Person
	fmt.Println(profile)

	/*
		Все поля структуры Person приватные, получить значения запросом > profile.firstName возможности нет.
		 - Спецификация доступа определяется регистром первого литерала в названии поля или метода, в рамках структуры (имена с большой буквы публичные, с маленькой - приватные (в рамках пакета))
		Есть возможность получить значения в экземпляр дублирующего типа PersonData
	*/

	var profileData class.Person_data = profile.GetData()
	fmt.Printf("данные экземпляра: %v > Значение поля FirstName: %s", profileData, profileData.FirstName)

}
