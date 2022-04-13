package scraper

import (
	"data-pirates-challenge-go-version/model"
	"data-pirates-challenge-go-version/services/correiosapi"
	"data-pirates-challenge-go-version/services/datawriter"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"log"
)

type (
	ScraperSvc interface {
		StartScraping() error
		GetUfListFromDoc() ([]string, error)
		GetZipCodeRangeFromDoc(string, interface{}, interface{}) ([][]string, bool, error)
	}

	scraperImp struct {
		api    correiosapi.CorreiosApiSvc
		writer datawriter.DataWriterSvc
	}
)

func NewScraperSvc(api correiosapi.CorreiosApiSvc, writer datawriter.DataWriterSvc) ScraperSvc {
	return &scraperImp{
		api:    api,
		writer: writer,
	}
}

func (s *scraperImp) StartScraping() error {

	ufList, err := s.GetUfListFromDoc()
	if err != nil {
		return err
	}

	for _, uf := range ufList {
		log.Println(fmt.Sprintf("scraping data for uf: %s", uf))

		zipCodeLocationList := make([]model.ZipCoreRange, 0)

		// Get data from first page
		ufZipCodeRange, hasNextPage, err := s.GetZipCodeRangeFromDoc(uf, nil, nil)
		if err != nil {
			return err
		}
		for _, info := range ufZipCodeRange {
			locationInfo := model.ZipCoreRange{
				Location: info[0],
				Range:    info[1],
			}
			zipCodeLocationList = append(zipCodeLocationList, locationInfo)
		}

		startPageParam := 51
		endPageParam := 100

		// Goes through next pages
		for hasNextPage {
			remainingZipCodeRangeList, nextPage, err := s.GetZipCodeRangeFromDoc(uf, startPageParam, endPageParam)
			if err != nil {
				return err
			}
			for _, info := range remainingZipCodeRangeList {
				locationInfo := model.ZipCoreRange{
					Location: info[0],
					Range:    info[1],
				}
				zipCodeLocationList = append(zipCodeLocationList, locationInfo)
			}

			// iterate pages
			startPageParam += 50
			endPageParam += 50
			hasNextPage = nextPage
		}

		// Remove duplicates
		uniqueValues := RemoveDuplicates(zipCodeLocationList)

		// Generate unique ids
		uniqueValuesWithID := make([]model.ZipCoreRange, 0)
		for _, locationData := range uniqueValues {
			uniqueValuesWithID = append(uniqueValuesWithID, model.ZipCoreRange{
				ID:       uuid.New().String(),
				Location: locationData.Location,
				Range:    locationData.Range,
			})
		}

		log.Println(fmt.Sprintf("%d locations founded to uf: %s",
			len(uniqueValuesWithID), uf))

		// Save file
		if err := s.writer.SaveJsonlFile(uf, uniqueValuesWithID); err != nil {
			return err
		}
	}
	return nil
}

func (s *scraperImp) GetZipCodeRangeFromDoc(uf string, startPage interface{}, endPage interface{}) ([][]string, bool, error) {
	res, err := s.api.GetData(uf, startPage, endPage)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, false, err
	}
	data := make([][]string, 0)
	doc.Find("table").Find("tr").Each(func(i int, selection *goquery.Selection) {
		tempTr := make([]string, 0)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			tempTr = append(tempTr, toUtf8([]byte(selection.Text())))
		})

		// Check if slice contain all information about location
		if len(tempTr) == 4 {
			data = append(data, tempTr)
		}
	})

	// CheckForNestPage
	hasNextPage := false
	doc.Find("form").Each(func(i int, selection *goquery.Selection) {
		name, _ := selection.Attr("name")
		if name == "Proxima" {
			hasNextPage = true
		}
	})
	return data, hasNextPage, err
}

func (s *scraperImp) GetUfListFromDoc() ([]string, error) {
	res, err := s.api.GetHomePage()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// searching for UFs
	ufList := make([]string, 0)

	doc.Find("select").Each(func(i int, selectionList *goquery.Selection) {
		className, _ := selectionList.Attr("class")
		if className == "f1col" {
			selectionList.Find("option").Each(func(i int, selectionOptions *goquery.Selection) {
				if value, ok := selectionOptions.Attr("value"); ok && value != "" {
					ufList = append(ufList, value)
				}
			})
		}
	})
	return ufList, nil
}

func toUtf8(ISO88591Buf []byte) string {
	buf := make([]rune, len(ISO88591Buf))
	for i, b := range ISO88591Buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

func RemoveDuplicates(data []model.ZipCoreRange) []model.ZipCoreRange {
	m := make(map[model.ZipCoreRange]struct{})
	dataSlice2 := make([]model.ZipCoreRange, 0)
	for _, d := range data {
		if _, ok := m[d]; !ok {
			dataSlice2 = append(dataSlice2, d)
			m[d] = struct{}{}
		}
	}
	return dataSlice2
}
