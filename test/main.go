package main

import "fmt"

func main() {
	m := make(map[int]string)
	fmt.Println(m)

	m[1] = "hello"
	m[3] = ""
	fmt.Println(m)

	k := m[2]
	fmt.Println(k)

}

//type Expense struct {
//	Title    string
//	Category string
//	Amount   float64
//}
//
//// все расходы можно хранить в стуктуре
//type Tracker struct {
//	Expenses []Expense
//}
//
//func (t *Tracker) AddExpense(title, category string, amount float64) {
//	e := Expense{Title: title, Category: category, Amount: amount}
//	t.Expenses = append(t.Expenses, e)
//
//}
//
//func (t *Tracker) AllExpenses() {
//	if len(t.Expenses) == 0 {
//		fmt.Println("Расходов нет")
//		return
//	}
//	for index, value := range t.Expenses {
//		fmt.Printf("%d. %s | %s | %.2f rub. \n", index+1, value.Title, value.Category, value.Amount)
//	}
//}
//
//func (t *Tracker) TotalSum() {
//	//[Expense1, Expense2, Expense3]
//
//	var sum float64
//	for _, value := range t.Expenses {
//		sum = sum + value.Amount
//	}
//	fmt.Printf("Общая сумма расходов: %.2f\n", sum)
//}
//
//func (t *Tracker) ShowStatistic() {
//	if len(t.Expenses) == 0 {
//		fmt.Println("Расходов нет")
//		return
//	}
//	categorySum := make(map[string]float64)
//	for _, value := range t.Expenses {
//		categorySum[value.Category] = categorySum[value.Category] + value.Amount
//	}
//	fmt.Println(categorySum)
//	// Вывести ключ-значение в нормально читаемом виде
//	// Нужно пройтись циклом по мапе
//}
//
//func (t *Tracker) DeleteExpense(number int) {
//	if len(t.Expenses) <= number || number < 0 {
//		fmt.Println("Неверный номер расхода")
//		return
//	}
//	// Объединяем часть ДО индекса и часть ПОСЛЕ индекса
//	// slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
//	// [a,b,c,d,e,f] [:5]
//	t.Expenses = append(t.Expenses[:number], t.Expenses[number+1:]...)
//	fmt.Println("Расход удален")
//
//}
//
//func (t *Tracker) ShowMaxSumExpense() {
//	if len(t.Expenses) == 0 {
//		fmt.Println("Расходов нет")
//		return
//	}
//	maxAmount := t.Expenses[0].Amount
//	maxExpense := t.Expenses[0]
//	for i := 1; i < len(t.Expenses); i++ {
//		if maxAmount < t.Expenses[i].Amount {
//			maxAmount = t.Expenses[i].Amount
//			maxExpense = t.Expenses[i]
//		}
//	}
//	// Вывести наименование и категорию самого большого расхода
//	// Нужно хранить как maxAmount переменную maxExpense, где до цикла for она равна первому расходу
//
//	// в условии if она также должна подменяться
//	fmt.Println("Самый большой расход равен:", maxExpense.Title, maxExpense.Category, maxExpense.Amount)
//
//	// for
//	// for i := 0; i < 10; i++
//	// for i < 10
//	// for k, v := range t.Expenses
//}
//
//func main() {
//
//	tracker := Tracker{}
//
//	for {
//		printMenu()
//		scanner := bufio.NewScanner(os.Stdin)
//		scanner.Scan()
//		input := scanner.Text()
//		fmt.Println(input)
//		menu, err := strconv.Atoi(input)
//		if err != nil {
//			fmt.Println("Введите число!")
//			continue
//		}
//		switch menu {
//		case 1:
//			fmt.Println("Введите наименование расхода")
//			scanner.Scan()
//			title := scanner.Text()
//			fmt.Println("Введите категорию расхода")
//			scanner.Scan()
//			category := scanner.Text()
//			fmt.Println("Введите сумму")
//			scanner.Scan()
//			amountStr := scanner.Text()
//
//			//var amount float64
//			//запустить беск цикл, пока не введут нормальное число, если ввели не подходящее число, оставнавливаем цикл
//			val, err := strconv.ParseFloat(amountStr, 64)
//			if err != nil {
//				fmt.Println("Ошибка конвертации:", err)
//				return
//			}
//
//			tracker.AddExpense(title, category, val)
//
//		case 2:
//			tracker.AllExpenses()
//
//		case 3:
//			tracker.TotalSum()
//
//		case 4:
//			tracker.ShowStatistic()
//
//		case 5: // Удалить элемент в слайсе -1
//			//проверить наличие расходов, и если нет, то не запрашивать ввод номера расхода, завершать работу
//			tracker.AllExpenses()
//			fmt.Println("Введите номер расхода для удаления из базы")
//			scanner.Scan()
//			input := scanner.Text()
//			number, err := strconv.Atoi(input)
//			if err != nil {
//				fmt.Println("Введите число")
//				continue
//			}
//			tracker.DeleteExpense(number - 1)
//
//		case 6:
//			tracker.ShowMaxSumExpense()
//
//		case 7:
//			fmt.Println("Выходим из программы")
//			return
//
//		default:
//			fmt.Println("Такой функционал отсутствует")
//		}
//	}
//
//}
//
////мини приложение для учёта личных расходов
////лежать все должно в одном файле main.go
//// хранить запись о расходе в такой структуре -
//// так ты сможешь накинуть методы на структуру
//
//// вот такие возможности должен реализовывать наш проектик
//func printMenu() {
//	fmt.Println("\n====== Трекер расходов ======")
//	fmt.Println("1. Добавить расход")
//	fmt.Println("2. Показать все расходы")
//	fmt.Println("3. Показать общую сумму")
//	fmt.Println("4. Показать статистику по категориям")
//	fmt.Println("5. Удалить расход")
//	fmt.Println("6. Показать самый большой расход")
//	fmt.Println("7. Выход")
//	fmt.Print("Выберите пункт меню: ")
//}
