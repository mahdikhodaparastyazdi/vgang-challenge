package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	// "github.com/mahdikhodaparast/vgang-challenge/pkg/utils"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// dataGetterService is a service that gets data from vgang api.
type dataGetterService struct {
	remainItems     int
	config          *pkg.Config
	productURlChan  chan [2]int
	vgangRepository model.VgangRepository
	productIDChan   chan int
	Lock            sync.Mutex
	token           *model.GetTokenResponse
	j               int
}

// NewDataGetterService creates a new dataGetterService
func NewDataGetterService(config *pkg.Config, vgangRepository model.VgangRepository, productURlChan chan [2]int, productIDChan chan int, token *model.GetTokenResponse) (model.Runner, error) {
	return &dataGetterService{
		config:          config,
		vgangRepository: vgangRepository,
		productURlChan:  productURlChan,
		productIDChan:   productIDChan,
		token:           token,
		remainItems:     24,
		j:               0,
	}, nil
}

// Run runs the service
func (d *dataGetterService) Run(ctx context.Context) error {
	ctx, done := context.WithCancel(ctx)
	defer done()
	defer close(d.productIDChan)
	waitCh := make(chan struct{})

	var wg sync.WaitGroup
	go func() {
		for i := 0; i < d.config.DataGetterWorkerSize; i++ {
			wg.Add(1)
			go d.worker(&wg, done)
		}
		wg.Wait()
		fmt.Println("shod")
		close(waitCh)
	}()

	log.Println("Data getter Service Started with ", d.config.DataGetterWorkerSize, " workers")
	for {
		select {
		// case <-ctx.Done():
		// 	if d.RemainItems > 0 {
		// 		return fmt.Errorf("there is %d items left", d.RemainItems)
		// 	}
		// 	return nil
		// case <-waitCh:
		// 	if d.RemainItems > 0 {
		// 		return fmt.Errorf("there is %d items left", d.RemainItems)
		// 	}
		// 	return nil
		case <-waitCh:
			fmt.Println("yoh~ahha")
			return nil

		case <-ctx.Done():
			fmt.Println("asdadada")
			return nil
		}
	}

}

// worker is a worker function of the service. It will run until the  channel is closed.
func (d *dataGetterService) worker(wg *sync.WaitGroup, done context.CancelFunc) {
	defer wg.Done()
	// fmt.Println("data_getter:before range productURlChan")

	for categoryWithOffset := range d.productURlChan {
		// if categoryWithOffset[1] == -1 {
		// 	fmt.Println("total:", categoryWithOffset[0])
		// 	fmt.Println("d.remainItems:", d.remainItems)
		// 	d.remainItems = categoryWithOffset[0] - d.remainItems + 1
		// 	fmt.Println("d.remainItems:", d.remainItems)
		// 	// done()
		// 	// return
		// }
		// d.Lock.Lock()
		// d.remainItems--
		// d.Lock.Unlock()

		if d.remainItems <= 0 {
			fmt.Println("d.remainItems")
			// done()
			// close(d.categoryWithOffset)
			// return
		}

		// categoryWithOffset, err := d.vgangRepository.GetData(categoryWithOffset[0])
		// if err != nil || (categoryWithOffset == nil && err == nil) {
		// 	continue
		// }
		fmt.Println("offest:", categoryWithOffset[1])
		productIDsWithTotal, err := d.vgangRepository.GetProductIDs(d.token, categoryWithOffset[0], 48, categoryWithOffset[1])
		// fmt.Println("offest:asdad")
		if err != nil || (productIDsWithTotal.ProductId == nil && err == nil) {
			fmt.Println(err)
			continue
		}

		for _, productID := range productIDsWithTotal.ProductId {
			// d.Lock.Lock()
			// d.j++
			// d.Lock.Unlock()
			d.Lock.Lock()
			d.j++
			d.Lock.Unlock()
			fmt.Println("productID:j:", d.j)
			d.productIDChan <- productID
			if d.j == 1136 {
				done()
			}
		}

		// }
	}
}
