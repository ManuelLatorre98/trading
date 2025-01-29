package dolars

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"manulatorre98/trading/graph/model"
	"manulatorre98/trading/repository/dolarRepository"
	"net/http"
	"strconv"
	"strings"
)

func scrapPage(url string) {
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

	var dolar model.DolarInput

	// Find the elements that contain the dollar values
	doc.Find(".tile").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h5").Text()
		value := s.Find(".val").Text()

		// Check if value is empty
		if value == "" {
			log.Printf("Empty value for: %s", title)
			return
		}

		// Clean the value and parse to float64
		cleanedValue := strings.ReplaceAll(strings.ReplaceAll(value, "$", ""), ",", "")
		floatValue, err := strconv.ParseFloat(cleanedValue, 32)
		if err != nil {
			log.Printf("Error parsing value for %s: %s", title, err)
			return
		}

		// Assign the parsed value to the appropriate field
		switch title {
		case "Dólar Blue":
			dolar.PriceBlue = float32(floatValue)
		case "Dólar Oficial":
			dolar.PriceOfficial = float32(floatValue)
		case "Dólar MEP":
			dolar.PriceMep = float32(floatValue)
		}
	})

	// Print the structured data
	fmt.Printf("Dolar data: %+v\n", dolar)
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
