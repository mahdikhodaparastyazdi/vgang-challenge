package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// dataRendererService is a service that gather productID data  and renders the data to stdout
type dataRendererService struct {
	config        *pkg.Config
	done          chan bool
	productIDChan chan int
	ProductIDs    sync.Map
	writer        io.Writer
	Lock          sync.Mutex
}

// NewDataRendererService creates a new dataRendererService
func NewDataRendererService(config *pkg.Config, productIDChan chan int, writer io.Writer, done chan bool) (model.Runner, error) {
	return &dataRendererService{
		config:        config,
		productIDChan: productIDChan,
		writer:        writer,
		done:          done,
		ProductIDs:    sync.Map{},
	}, nil
}

// Run runs the service
func (d *dataRendererService) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	for i := 0; i < d.config.DataGetterWorkerSize; i++ {
		wg.Add(1)
		go d.worker(ctx, &wg)
	}
	log.Println("Data renderer Service Started with ", d.config.DataRendererWorkerSize, " workers")
	wg.Wait()
	d.GenerateProductData()
	go func() {
		d.done <- true
	}()
	return nil
}

// worker is a worker function of the service. It will run until the  channel is closed.
func (d *dataRendererService) worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for productID := range d.productIDChan {

		d.ProductIDs.Store(productID, productID)

		// d.productIDs = append(d.productIDs, productID)
		// d.Lock.Lock()
		// d.i++
		// d.Lock.Unlock()
		// fmt.Println("d.sadadadasdaproductIDs::", d.i)
		// fmt.Println("ASDDDDDDDDDAasdasdaSDASDASDADADA::", len(d.productIDs))
		// if i%20 == 0 {
		// 	// fmt.Println("d.sadadadasdaproductIDs::", i)
		// }
		// if i == 1110 {
		// 	// fmt.Println(len(d.productIDs))
		// }

	}
}

// generateProductersData generates the data for request
func (d *dataRendererService) GenerateProductData() {
	d.ProductIDs.Range(func(key, value interface{}) bool {
		productID := value.(int)
		_, err := d.writer.Write([]byte(fmt.Sprintf("ProDuctID: %d\n", productID)))
		if err != nil {
			log.Panicln("there is an error in writing data", err)
			return false
		}
		return true
	})
	// fmt.Println("adadsadaADasdASDASDA", len(d.productIDs))
	// for _, productID := range d.productIDs {
	// 	_, err := d.writer.Write([]byte(fmt.Sprintf("ProDuctID: %d\n", productID)))
	// 	if err != nil {
	// 		log.Panicln("there is an error in writing data", err)
	// 	}
	// }
}
