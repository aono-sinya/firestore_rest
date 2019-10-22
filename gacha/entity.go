package gacha

import (
	"math/rand"
	"sync"
	"time"
)

type GachaEntity struct {
	UserID string
	Value int
}

type Gacha struct {
	wg *sync.WaitGroup
	mtx *sync.Mutex
	UserID string
	Entities []GachaEntity
}

func (g *Gacha) Lot() {
	defer g.wg.Done()
	defer g.mtx.Unlock()

	rand.Seed(time.Now().UnixNano())
	en := GachaEntity{
		UserID: g.UserID,
		Value: rand.Intn(100)}
	g.Entities = append(g.Entities, en)
}

func (g *Gacha) Lock() {
	g.wg.Add(1)
	g.mtx.Lock()
}

func (g *Gacha) Wait() {
	g.wg.Wait()
}