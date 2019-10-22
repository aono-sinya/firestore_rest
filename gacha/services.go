package gacha

import (
	"encoding/json"
	"log"
	"sync"
)

func initGacha(userID string) *Gacha {
	g := &Gacha{
		wg: new(sync.WaitGroup),
		mtx: new(sync.Mutex),
		UserID: userID,
		Entities: make([]GachaEntity, 0, 10)}
	return g
}

func GachaExecute(userID string) []byte {
	g := initGacha(userID)
	for i := 0; i <= 10; i++ {
		g.Lock()
		go g.Lot()
	}
	g.Wait()
	res, err := json.Marshal(g.Entities); if err != nil {
		log.Fatalln(err)
	}
	return res
}
