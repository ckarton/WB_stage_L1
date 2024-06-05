package main

// programm 1
// type Human struct {
//     Name string
// }

// func (h Human) Speak() {
//     fmt.Println("Hi, my name is", h.Name)
// }

// type Action struct {
//     Human
// }

// func main() {
//     action := Action{
//         Human: Human{
//             Name: "Bob",
//         },
//     }

//     action.Speak()
// }

// programm 2
// import (
//     "fmt"
//     "sync"
// )
// func main() {

//     numbers := []int{2, 4, 6, 8, 10}

//     results := make(chan int)

//     var wg sync.WaitGroup

//     for _, num := range numbers {
//         wg.Add(1)
//         go func(n int) {
//             defer wg.Done()
//             results <- n * n
//         }(num)
//     }

//     go func() {
//         wg.Wait()
//         close(results)
//     }()

//     for square := range results {
//         fmt.Println(square)
//     }
// }

// programm 3
// import (
//     "fmt"
//     "sync"
// )
// func main() {
// 	numbers := []int{2, 4, 6, 8, 10}
// 	results := make(chan int)
// 	var wg sync.WaitGroup

// 	for _, num := range numbers {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			results <- n * n
// 		}(num)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	sum := 0
// 	for square := range results {
// 		sum += square
// 	}
// 	fmt.Println("Сумма квадратов:", sum)
// }

// programm 4
// import (
//     "fmt"
//     "math/rand"
//     "os"
//     "os/signal"
//     "sync"
//     "syscall"
//     "time"
// )

// func worker(id int, jobs <-chan int, wg sync.WaitGroup) {
//     defer wg.Done()
//     for job := range jobs {
//         fmt.Printf("Worker %d received job: %d\n", id, job)
//     }
// }

// func main() {
//     numWorkers := 3 // Количество воркеров

//     if len(os.Args) > 1 {
//         fmt.Sscanf(os.Args[1], "%d", &numWorkers)
//     }

//     jobs := make(chan int)
//     var wg sync.WaitGroup

//     for i := 1; i <= numWorkers; i++ {
//         wg.Add(1)
//         go worker(i, jobs, wg)
//     }

//     stop := make(chan os.Signal, 1)
//     signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

//     go func() {
//         for {
//             select {
//             case <-stop:
//                 close(jobs)
//                 return
//             default:
//                 jobs <- rand.Intn(100)
//                 time.Sleep(500 *time.Millisecond)
//             }
//         }
//     }()

//     <-stop
//     close(jobs)  // Убедимся, что канал jobs закрыт
//     wg.Wait()
// }

// progamm 5
// import (
//     "fmt"
//     "time"
// )

// func main() {
//     N := 5
//     ch := make(chan int)
//     done := make(chan struct{})

//     go func() {
//         for {
//             select {
//             case val := <-ch:
//                 fmt.Println("Received:", val)
//             case <-done:
//                 return
//             }
//         }
//     }()

//     go func() {
//         for i := 0; ; i++ {
//             select {
//             case ch <- i:
//                 fmt.Println("Sent:", i)
//             case <-done:
//                 return
//             }
//             time.Sleep(time.Second)
//         }
//     }()

//     time.Sleep(time.Duration(N) * time.Second)

//     close(done)
// }

// programm 6
// import (
//     "fmt"
//     "time"
// )

// func worker(done chan struct{}) {
//     fmt.Println("Worker started")
//     defer fmt.Println("Worker stopped")
//     <-done
// }

// func main() {
//     done := make(chan struct{})
//     go worker(done)
//     time.Sleep(1 * time.Second) // представим что это рутина
//     close(done)
//     time.Sleep(1 * time.Second)
// }

// programm 7
// import (
//     "fmt"
//     "sync"
// )

// func main() {
//     data := make(map[string]int)
//     var mutex sync.Mutex
//     for i := 0; i < 5; i++ {
//         go func(index int) {
//             mutex.Lock()
//             data[fmt.Sprintf("key%d", index)] = index
//             mutex.Unlock()
//         }(i)
//     }
//     fmt.Println("Waiting for goroutines to finish...")
//     mutex.Lock()
//     defer mutex.Unlock()
//     fmt.Println("Data in map:")
//     for key, value := range data {
//         fmt.Printf("%s: %d\n", key, value)
//     }
// }

// programm 8
// import "fmt"

// // Устанавливает i-й бит в 1
// func setBit(num *int64, i uint) {
//     *num |= 1 << i
// }

// // Устанавливает i-й бит в 0
// func clearBit(num *int64, i uint) {
//     *num &^= 1 << i
// }

// func main() {
//     var num int64 = 42
//     fmt.Printf("Исходное число: %d\n", num)

//     // Установка i-го бита в 1
//     i := uint(2)
//     setBit(&num, i)
//     fmt.Printf("Число после установки %d-го бита в 1: %d\n", i, num)

//     // Установка i-го бита в 0
//     clearBit(&num, i)
//     fmt.Printf("Число после установки %d-го бита в 0: %d\n", i, num)
// }

// programm 9
// import "fmt"

// func main() {
//     input := make(chan int)
//     output := make(chan int)

//     go func() {
//         defer close(input)
//         numbers := []int{1, 2, 3, 4, 5}
//         for _, num := range numbers {
//             input <- num
//         }
//     }()

//     go func() {
//         defer close(output)
//         for num := range input {
//             output <- num * 2
//         }
//     }()

//     for result := range output {
//         fmt.Println(result)
//     }
// }

// programm 10
// import (
//     "fmt"
//     "math"
// )

// func groupTemperatures(temperatures []float64, step float64) map[float64][]float64 {
//     groups := make(map[float64][]float64)

//     for _, temp := range temperatures {
//         groupKey := math.Floor(temp/step) * step

//         groups[groupKey] = append(groups[groupKey], temp)
//     }

//     return groups
// }

// func main() {
//     temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
//     groupStep := 10.0

//     groups := groupTemperatures(temperatures, groupStep)

//     for key, group := range groups {
//         fmt.Printf("Диапазон: %.1f - %.1f\n", key, key+groupStep)
//         fmt.Println("Температуры:", group)
//         fmt.Println()
//     }
// }

// programm 11
// import "fmt"

// func intersection(set1, set2 map[int]bool) map[int]bool {
//     result := make(map[int]bool)

//     for element := range set1 {
//         if set2[element] {
//             result[element] = true
//         }
//     }

//     return result
// }

// func main() {

//     set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}

//     set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

//     intersectionSet := intersection(set1, set2)

//     fmt.Println("Пересечение множеств:", intersectionSet)
// }

// programm 12
// import "fmt"

// func main() {
//     strings := []string{"cat", "cat", "dog", "cat", "tree"}
//     set := make(map[string]bool)

//     for _, str := range strings {
//         set[str] = true
//     }

//     fmt.Println("Множество строк:", set)
// }

// programm 13
// import "fmt"

// func main() {
//     a := 5
//     b := 3

//     fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

//     a = a + b
//     b = a - b
//     a = a - b

//     fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
// }

// programm 14
// import (
//     "fmt"
// )

// func identifyType(v interface{}) {
//     switch v.(type) {
//     case int:
//         fmt.Println("Это int")
//     case string:
//         fmt.Println("Это string")
//     case bool:
//         fmt.Println("Это bool")
//     case chan int:
//         fmt.Println("Это channel int")
//     case chan string:
//         fmt.Println("Это channel string")
//     case chan bool:
//         fmt.Println("Это channel bool")
//     default:
//         fmt.Println("Неизвестный тип")
//     }
// }

// func main() {
//     var a int = 10
//     var b string = "hello"
//     var c bool = true
//     var d chan int = make(chan int)
//     var e chan string = make(chan string)
//     var f chan bool = make(chan bool)

//     identifyType(a)
//     identifyType(b)
//     identifyType(c)
//     identifyType(d)
//     identifyType(e)
//     identifyType(f)
// }

// programm 15
// import (
//     "fmt"
// )

// func createHugeString(size int) string {
//     hugeString := make([]byte, size)
//     for i := range hugeString {
//         hugeString[i] = 'a'
//     }
//     return string(hugeString)
// }

// func someFunc() string {
//     v := createHugeString(1 << 10)
//     if len(v) < 100 {
//         return v
//     }
//     return string(v[:100])
// }

// func main() {
//     justString := someFunc()
//     fmt.Println(justString)
// }

// programm 16
// import (
//     "fmt"
// )

// func quicksort(arr []int) []int {
//     if len(arr) <= 1 {
//         return arr
//     }

//     pivot := arr[0]
//     left := []int{}
//     right := []int{}

//     for _, v := range arr[1:] {
//         if v < pivot {
//             left = append(left, v)
//         } else {
//             right = append(right, v)
//         }
//     }

//     left = quicksort(left)
//     right = quicksort(right)

//     return append(append(left, pivot), right...)
// }

// func main() {
//     arr := []int{10, 7, 8, 9, 1, 5}
//     fmt.Println("Исходный массив:", arr)
//     sortedArr := quicksort(arr)
//     fmt.Println("Отсортированный массив:", sortedArr)
// }

// programm 17
// import (
//     "fmt"
//     "sort"
// )

// func main() {
//     // Инициализируем отсортированный массив
//     arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

//     // Искомое значение
//     target := 13

//     // Используем встроенную функцию sort.Search для бинарного поиска
//     index := sort.Search(len(arr), func(i int) bool {
//         return arr[i] >= target
//     })

//     // Проверяем, найдено ли значение
//     if index < len(arr) && arr[index] == target {
//         fmt.Printf("Значение %d найдено на позиции %d\n", target, index)
//     } else {
//         fmt.Printf("Значение %d не найдено\n", target)
//     }
// }

// programm 18
// import (
//     "fmt"
//     "sync"
// )
// type Counter struct {
//     mu    sync.Mutex
//     value int
// }
// func (c *Counter) Increment() {
//     c.mu.Lock()
//     c.value++
//     c.mu.Unlock()
// }
// func (c *Counter) Value() int {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     return c.value
// }

// func main() {
//     var wg sync.WaitGroup
//     counter := &Counter{}
//     numGoroutines := 520
//     for i := 0; i < numGoroutines; i++ {
//         wg.Add(1)
//         go func() {
//             defer wg.Done()
//             counter.Increment()
//         }()
//     }
//     wg.Wait()
//     fmt.Println("Итоговое значение счетчика:", counter.Value())
// }

// programm 19
// import (
//     "fmt"
// )

// func reverseString(input string) string {
//     runes := []rune(input)
//     n := len(runes)

//     for i := 0; i < n/2; i++ {
//         runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
//     }

//     return string(runes)
// }

// func main() {
//     input := "qwertyuio"
//     fmt.Println("Исходная строка:", input)
//     reversed := reverseString(input)
//     fmt.Println("Перевернутая строка:", reversed)
// }

// programm 20
// import (
//     "fmt"
//     "strings"
// )

// func reverseWords(input string) string {
//     words := strings.Fields(input)

//     n := len(words)
//     for i := 0; i < n/2; i++ {
//         words[i], words[n-1-i] = words[n-1-i], words[i]
//     }
//     return strings.Join(words, " ")
// }

// func main() {
//     input := "snow dog sun"
//     fmt.Println("Исходная строка:", input)
//     reversed := reverseWords(input)
//     fmt.Println("Строка с перевёрнутыми словами:", reversed)
// }

// programm 21
// import (
//     "fmt"
// )

// type Printer interface {
//     Show()
// }

// type LegacyPrinter struct{}

// func (lp *LegacyPrinter) Display() {
//     fmt.Println("Legacy Printer is displaying")
// }

// type PrinterImpl struct {
//     legacyPrinter *LegacyPrinter
// }

// func (pi *PrinterImpl) Show() {
//     pi.legacyPrinter.Display()
// }

// func main() {
//     legacyPrinter := &LegacyPrinter{}
//     adapter := &PrinterImpl{
//         legacyPrinter: legacyPrinter,
//     }
//     adapter.Show()
// }

// programm 22
// import (
//     "fmt"
// )

// func main() {
//     a := 2<<20 + 1 // Пример числа, большего, чем 2^20
//     b := 2<<20 + 2 // Еще один пример числа, большего, чем 2^20

//     // Умножение
//     multiplication := a * b
//     fmt.Printf("%d * %d = %d\n", a, b, multiplication)

//     // Деление
//     division := a / b
//     fmt.Printf("%d / %d = %d\n", a, b, division)

//     // Сложение
//     addition := a + b
//     fmt.Printf("%d + %d = %d\n", a, b, addition)

//     // Вычитание
//     subtraction := a - b
//     fmt.Printf("%d - %d = %d\n", a, b, subtraction)
// }

// programm 23
// import "fmt"

// func removeIndex(slice []int, index int) []int {
//     return append(slice[:index], slice[index+1:]...)
// }

// func main() {
//     slice := []int{1, 2, 3, 4, 5}
//     index := 0
//     fmt.Println("Исходный срез:", slice)
//     slice = removeIndex(slice, index)
//     fmt.Println("Срез после удаления элемента с индексом", index, ":", slice)
// }

// programm 24
// import (
//     "fmt"
//     "math"
// )

// type Point struct {
//     x, y float64
// }

// // Конструктор для структуры Point.
// func NewPoint(x, y float64) Point {
//     return Point{x, y}
// }

// // Метод для вычисления расстояния между текущей точкой и другой точкой.
// func (p Point) DistanceTo(other Point) float64 {
//     dx := p.x - other.x
//     dy := p.y - other.y
//     return math.Sqrt(dx*dx + dy*dy)
// }

// func main() {
//     point1 := NewPoint(1, 2)
//     point2 := NewPoint(4, 6)

//     distance := point1.DistanceTo(point2)

//     fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) составляет %.2f\n",
//         point1.x, point1.y, point2.x, point2.y, distance)
// }

// programm 25
// import (
//     "fmt"
//     "time"
// )

// func sleep(seconds int) {
//     time.Sleep(time.Duration(seconds) * time.Second)
// }

// func main() {
//     fmt.Println("Начало программы")
//     sleep(5) // Приостановка выполнения на 5 секунд
//     fmt.Println("Программа завершена")
// }

// programm 26
// import (
//     "fmt"
//     "strings"
// )

// func isUnique(str string) bool {
//     str = strings.ToLower(str)
//     charMap := make(map[rune]bool)

//     for _, char := range str {
//         if charMap[char] {
//             return false
//         }
//         charMap[char] = true
//     }
//     return true
// }

// func main() {
//     stringsToCheck := []string{"qwer", "asdf", "ghjk"}

//     for _, str := range stringsToCheck {
//         fmt.Printf("Строка \"%s\": %t\n", str, isUnique(str))
//     }
// }
