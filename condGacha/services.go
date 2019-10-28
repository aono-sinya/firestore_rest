package condGacha

import (
	"sync"
	"time"
)

func initGacha(userID string) *Gacha {
	g := &Gacha {
		userID: userID,
		entities: make([]GachaEntity, 0, 10)}
	return g
}

func GachaExecute(userID string) []byte {
	g := initGacha(userID)
	mtx := new(sync.Mutex)
	cond := sync.NewCond(mtx)
	for i := 0; i <= 10; i++ {
		go g.Lot(cond)
	}
	for i := 0; i <= 10; i++ {
		cond.Signal()
	}
	time.Sleep(3 * time.Second)
	return g.Json()
}