package main

import "fmt"

func main() {

	Processor := &DataProcessor{}
	Persistence := &PersistentData{}
	tranform := &TransformData{}

	stockData := &StockData{
		ticker:    "abc",
		price:     100.0,
		observers: make([]StockDataObserver, 0),
	}
	stockData.RegisterObserver(Processor)
	stockData.RegisterObserver(Persistence)
	stockData.RegisterObserver(tranform)

	stockData.NotifyObservers()
	stockData.UpdateStockPrice(100)

}

type Subject interface {
	RegisterObserver(observer StockDataObserver)
	RemoveObserver(observer StockDataObserver)
	NotifyObservers()
}

// concrete subject
type StockData struct {
	ticker    string
	price     float64
	observers []StockDataObserver
}

func (s *StockData) RegisterObserver(observer StockDataObserver) {
	s.observers = append(s.observers, observer)
}

func (s *StockData) RemoveObserver(observer StockDataObserver) {

	for i, o := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *StockData) UpdateStockPrice(amount float64) {
	s.price = s.price + amount
	s.NotifyObservers()
}

func (s StockData) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}

}

type StockDataObserver interface {
	Update(context StockData)
}

// concrete observers

type DataProcessor struct{}

func (d *DataProcessor) Update(data StockData) {
	fmt.Println("prcessing data ", data.price, data.ticker)
}

type PersistentData struct {
}

func (d *PersistentData) Update(data StockData) {
	fmt.Println("saving  data ", data.price, data.ticker)
}

type TransformData struct{}

func (d *TransformData) Update(data StockData) {
	fmt.Println("transforming data", data.price, data.ticker)
}
