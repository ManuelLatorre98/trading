package dolars

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"manulatorre98/trading/graph/model"
	"manulatorre98/trading/repository/dolarRepository"
	"net/http"
)

type DolarScraped struct {
	Compra string `json:"compra"`
	Venta  string `json:"venta"`
}

func scrapPage(url string) {
	dolarMap := make(map[string]DolarScraped)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	validTitles := []string{
		"Dólar blue",
		"Dólar Oficial",
		"Dólar MEP/Bolsa",
		"Contado con liqui",
		"Dólar cripto",
		"Dólar Tarjeta",
	}
	// Buscar los elementos que contienen los valores del dólar
	doc.Find(".tile").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		compra := s.Find(".compra .val").Text()
		venta := s.Find(".venta .val").Text()
		// Filtrar solo los títulos específicos
		if isValidTitle(title, validTitles) {
			dolarMap[title] = DolarScraped{
				Compra: compra,
				Venta:  venta,
			}
		}
	})

	// Print the structured data
	fmt.Printf("Dolar data: %+v\n", dolarMap)
}

func isValidTitle(title string, validTitles []string) bool {
	for _, validTitle := range validTitles {
		if title == validTitle {
			return true
		}
	}
	return false
}
func CreateDolarResolver(repo dolarRepository.DolarRepository) (*model.Dolar, error) {
	scrapPage("https://www.dolarhoy.com/")
	/*	dolar := &model.DolarInput{15, 65435, 3453}
		newDolar, err := repo.CreateDolar(dolar)
		if err != nil {
			return nil, fmt.Errorf(customErrors.ErrInternalServer)
		}*/
	return nil, nil
}
