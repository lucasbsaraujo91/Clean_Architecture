package usecase

import (
	"Clean_Architecture/internal/entity"
	"strconv"
)

type FindAllOrderInputDTO struct {
	Page  string `json:"page"`
	Limit int32  `json:"limit"`
	Sort  string `json:"sort"`
}

type FindAllOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type FindAllOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrderUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *FindAllOrderUseCase {
	return &FindAllOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *FindAllOrderUseCase) FindAll(input FindAllOrderInputDTO) ([]FindAllOrderOutputDTO, error) {
	// Convertendo o valor de página de string para inteiro
	page, err := strconv.Atoi(input.Page)
	if err != nil {
		return nil, err
	}

	// Chamando o repositório para buscar as ordens
	orders, err := c.OrderRepository.FindAll(page, int(input.Limit), input.Sort)
	if err != nil {
		return nil, err
	}

	// Verificando se nenhuma ordem foi encontrada
	if len(orders) == 0 {
		return []FindAllOrderOutputDTO{}, nil
	}

	// Mapeando todas as ordens para o DTO
	var dtos []FindAllOrderOutputDTO
	for _, order := range orders {
		dto := FindAllOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
