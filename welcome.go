package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	setupLogging()
	mean()
	fmt.Println("--------")
	listChallenge()
	fmt.Println("--------")
	maxChallenge()
	fmt.Println("--------")
	mapTest()
	fmt.Println("--------")
	wordCountChallenge()
	fmt.Println("--------")
	div, mod := divmod(23, 6)
	fmt.Printf("div: %v, mod: %v\n", div, mod)
	fmt.Println("--------")
	num := 12
	doublePtr(&num)
	fmt.Printf("double is %v\n", num)
	fmt.Println("--------")
	s1, error := sqrt(2.0)
	if error != nil {
		//fmt.Printf("ERROR %s\n", error)
		log.Printf("ERROR %s", error)
	} else {
		fmt.Println(s1)
	}

	s2, error := sqrt(-2.0)
	if error != nil {
		//fmt.Printf("ERROR %s\n", error)
		log.Printf("ERROR %s", error)
	} else {
		fmt.Println(s2)
	}
	fmt.Println("--------")

	cType, body, error := contentTypeChallenge("http://faamili.de/faa")
	if error != nil {
		//fmt.Printf("ERROR: %s\n", error)
		log.Printf("ERROR %s", error)
	} else {
		fmt.Printf("Content type: %v\n", cType)
		fmt.Printf("Body length: %v\n", len(body))
	}
	fmt.Println("--------")
	tradeExample()
	fmt.Println("--------")

	square, error := NewSquare(1, 2, 10)
	if error != nil {
		//fmt.Printf("ERROR: %s\n", error)
		log.Printf("ERROR %s", error)
	}
	square.Move(2, 3)
	fmt.Printf("%+v\n", square)
	fmt.Println(square.Area())
	fmt.Println("--------")

	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello World from 12234325")
	fmt.Println("--------")

	var wg sync.WaitGroup
	urls := []string{"https://golang.org", "https://api.github.com", "https://httpbin.org/xml"}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Println("--------")

	channel()
	fmt.Println("--------")
	channelChallenge()
	fmt.Println("--------")
	selectChannel()
	fmt.Println("--------")

	md5Challenge()
	fmt.Println("--------")
}

// Log in the file
func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func mean() {
	x := 3.0
	y := 1.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result %v, type of %T\n", mean, mean)

	if frac := x / y; frac > 0.5 {
		fmt.Printf("x is more than half of y\n")
	}
}

func listChallenge() {

	for i := 1; i <= 20; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("fizz buzz\n")
		case i%5 == 0:
			fmt.Printf("buzz\n")
		case i%3 == 0:
			fmt.Printf("fizz\n")
		default:
			fmt.Printf("%v\n", i)
		}
	}

	loans := []string{"bugs", "daffy", "taz"}

	fmt.Printf("loans = %v (type %T)\n", loans, loans)

	fmt.Println(len(loans))
	fmt.Println(loans[2])

	for i, name := range loans {
		fmt.Printf("%s, at %d\n", name, i)
	}
}

func maxChallenge() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]

	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Printf("The maximum is %v\n", max)
}

func mapTest() {

	stocks := map[string]float64{"AMZN": 1699.0, "GOOG": 1129.19, "MSFT": 1024.12}

	fmt.Println(len(stocks))
	fmt.Printf("MSFT stock is %v$\n", stocks["MSFT"])

	stocks["TSLA"] = 345.24

	value, exist := stocks["TSLA"]
	if !exist {
		fmt.Println("TSLA not found")
	} else {
		fmt.Printf("TSLA stock is %v$\n", value)
	}

	delete(stocks, "GOOG")
	fmt.Println(stocks)
}

func wordCountChallenge() {
	text := `Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	wordsMap := map[string]int{}
	for _, word := range strings.Fields(text) {
		wordsMap[strings.ToLower(word)]++
	}

	fmt.Println(wordsMap)
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func doublePtr(n *int) {
	*n *= 2
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of nevative value (%f)", n)
	}
	return math.Sqrt(n), nil
}

func contentTypeChallenge(url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", "", err
	}

	return resp.Header.Get("Content-Type"), string(body), nil
}

type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be empty")
	}
	if volume < 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}
	if price < 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

func tradeExample() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Printf("Value of the trade: %v\n", t1.Value())

	t2, error := NewTrade("GOOG", 24, 34.34, true)
	if error != nil {
		fmt.Printf("Cannot create the trade instance: %v\n", error)
	} else {
		fmt.Printf("%+v\n", t2)
		fmt.Printf("Value of the trade: %v\n", t2.Value())
	}
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("length must be >= 0 (was %d)", length)
	}

	square := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return square, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (int, error) {
	// easy solution to upper case the string
	//s := string(p)
	//return c.wtr.Write([]byte(strings.ToUpper(s)))

	// may be a better solution
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.wtr.Write(out)
}

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR %s", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}
