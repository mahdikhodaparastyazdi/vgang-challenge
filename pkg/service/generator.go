package service

import (
	"context"
	"fmt"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// generatorService is a service that generates product ids.
type generatorService struct {
	productURlChan  chan [2]int
	config          *pkg.Config
	token           *model.GetTokenResponse
	vgangRepository model.VgangRepository
}

// NewGeneratorService creates a new generatorService
func NewGeneratorService(config *pkg.Config, vgangRepository model.VgangRepository, productURlChan chan [2]int, token *model.GetTokenResponse) (model.Runner, error) {
	return &generatorService{
		config:          config,
		productURlChan:  productURlChan,
		token:           token,
		vgangRepository: vgangRepository,
	}, nil
}

// Run runs the service
func (g *generatorService) Run(context.Context) error {

	categoryIDsWithTotalCount, err := g.vgangRepository.GetCategoriesIDs(g.token)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("ALL Categories:", categoryIDsWithTotalCount)
	// defer close(g.productURlChan)
	// for i := 1; i <= g.config.MaxSearchNumbers; i++ {
	// 	g.productURlChan <- i
	// }

	// return nil

	//In Task mentioned only products in one category
	//BTW it's use all categories
	// for _, categoryID := range *categoryIDsWithTotalCount {
	categoryID := 1
	// fmt.Println("categoryID", categoryID)
	count := 1
	offset := 0
	productIDsWithTotalCount, err := g.vgangRepository.GetProductIDs(g.token, categoryID, count, offset)
	if err != nil {
		return err
	}
	// g.productURlChan <- [2]int{24, -1}
	// fmt.Println("productIDsTotalCount: ", productIDsWithTotalCount.TotalCount)
	for offsset := 0; offsset < productIDsWithTotalCount.TotalCount; offsset = offsset + 48 {
		// fmt.Println("[2]int{categoryID, offsset}", [2]int{categoryID, offsset})
		g.productURlChan <- [2]int{categoryID, offsset}
	}
	// }
	// close(g.productURlChan)
	return nil
}
