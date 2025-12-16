package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"assembly_line/model"
)

type AssemblyLine struct {
	employees []*model.Employee
	items     []model.Item
	itemChan  chan model.Item
	wg        sync.WaitGroup
	startTime time.Time
	endTime   time.Time
}

func NewAssemblyLine(employeeCount int) *AssemblyLine {
	al := &AssemblyLine{
		employees: make([]*model.Employee, employeeCount),
		items:     make([]model.Item, 0),
		itemChan:  make(chan model.Item, 100),
	}

	for i := range al.employees {
		al.employees[i] = model.NewEmployee(i)
	}

	return al
}

func (al *AssemblyLine) AddItems(item1Count, item2Count, item3Count int) {
	for i := 1; i <= item1Count; i++ {
		al.items = append(al.items, model.NewItem1(i))
	}

	for i := 1; i <= item2Count; i++ {
		al.items = append(al.items, model.NewItem2(i))
	}

	for i := 1; i <= item3Count; i++ {
		al.items = append(al.items, model.NewItem3(i))
	}

	al.shuffleItems()
}

func (al *AssemblyLine) shuffleItems() {
	rand.Shuffle(len(al.items), func(i, j int) {
		al.items[i], al.items[j] = al.items[j], al.items[i]
	})
}

func (al *AssemblyLine) Start() {
	fmt.Println("========================================")
	fmt.Println("流水線開始運作")
	fmt.Printf("員工數量: %d\n", len(al.employees))
	fmt.Printf("物品總數: %d\n", len(al.items))
	fmt.Println("========================================")

	al.startTime = time.Now()

	for _, employee := range al.employees {
		al.wg.Add(1)
		go al.worker(employee)
	}

	go func() {
		for _, item := range al.items {
			al.itemChan <- item
		}
		close(al.itemChan)
	}()

	al.wg.Wait()
	al.endTime = time.Now()

	al.printStatistics()
}

func (al *AssemblyLine) worker(employee *model.Employee) {
	defer al.wg.Done()

	for item := range al.itemChan {
		employee.ProcessItem(item)
	}
}

func (al *AssemblyLine) printStatistics() {
	fmt.Println("\n========================================")
	fmt.Println("流水線處理完成")
	fmt.Println("========================================")

	workTime := al.endTime.Sub(al.startTime)
	fmt.Printf("總處理時間: %v\n\n", workTime)

	fmt.Println("員工處理統計:")
	totalProcessed := 0
	for _, employee := range al.employees {
		count := employee.GetProcessedCount()
		totalProcessed += count
		fmt.Printf("  員工 %d: 處理了 %d 件物品\n", employee.ID, count)
	}

	fmt.Printf("\n總共處理物品: %d 件\n", totalProcessed)
	fmt.Println("========================================")
}
