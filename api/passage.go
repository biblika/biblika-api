package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/biblika/biblika-api/model"
	"github.com/biblika/biblika-api/proto/passage"
)

type PassageService struct {
	passage.UnimplementedPassageServiceServer
}

var (
	Title = "title"
	Text  = "text"
)

func (s *PassageService) Passage(ctx context.Context, req *passage.PassageRequest) (*passage.PassageResponse, error) {
	params := fmt.Sprintf("%v+%d&ver=%s", req.Book, req.Chapter, req.Version)
	resp, err := http.Get("https://alkitab.sabda.org/api/passage.php?passage=" + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var xmlResponse model.Passage
	err = xml.Unmarshal(content, &xmlResponse)
	if err != nil {
		return nil, err
	}

	var verse []*passage.Verse

	for _, v := range xmlResponse.Book.Chapter.Verses.Verse {
		if len(v.Title) > 0 {
			verse = append(verse, &passage.Verse{
				Verse:   0,
				Type:    Title,
				Content: v.Title,
			})
		}
		verse = append(verse, &passage.Verse{
			Verse:   int32(v.Number),
			Type:    Text,
			Content: v.Text,
		})
	}

	response := passage.PassageResponse{
		Verse: verse,
	}

	return &response, nil
}
