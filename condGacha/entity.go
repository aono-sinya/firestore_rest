package condGacha

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"
	"time"
)

type GachaEntity struct {
	UserID string
	Value int
}

type Gacha struct {
	cond *sync.Cond
	userID string
	entities []GachaEntity
}

func (g *Gacha) Lot(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
	rand.Seed(time.Now().UnixNano())
	en := GachaEntity{
		UserID: g.userID,
		Value: rand.Intn(100)}
	g.entities = append(g.entities, en)
}

func (g *Gacha) Json() []byte {
	res, err := json.Marshal(g.entities); if err != nil {
		log.Fatalln(err)
	}
	return res
}