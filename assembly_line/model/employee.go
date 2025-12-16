package model

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID             int
	ProcessedCount int
	mu             sync.Mutex
}

func NewEmployee(id int) *Employee {
	return &Employee{
		ID:             id,
		ProcessedCount: 0,
	}
}

func (e *Employee) ProcessItem(item Item) {
	e.mu.Lock()
	defer e.mu.Unlock()

	fmt.Printf("[%s] 員工 %d 開始處理 %s (ID: %d)\n",
		time.Now().Format("15:04:05.000"), e.ID, item.GetType(), item.GetID())

	defer func() {
		e.ProcessedCount++
		fmt.Printf("[%s] 員工 %d 完成處理 %s (ID: %d) - 已處理 %d 件\n",
			time.Now().Format("15:04:05.000"), e.ID, item.GetType(), item.GetID(), e.ProcessedCount)
	}()

	item.Process()
}

func (e *Employee) GetProcessedCount() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ProcessedCount
}
