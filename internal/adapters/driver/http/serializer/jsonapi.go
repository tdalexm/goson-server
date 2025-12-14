package serializer

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tdalexm/goson-server/internal/domain"
)

type JSONAPISerializer struct {
	baseURL string
}

func NewJSONAPISerializer(baseUrl string) *JSONAPISerializer {
	return &JSONAPISerializer{baseURL: baseUrl}
}

func (s *JSONAPISerializer) SerializeCollection(
	collectionType string,
	collection []domain.Record,
	total, page, limit int,
	queryParams url.Values,
) JSONResponse {
	var data []ResourceObject

	for _, record := range collection {
		data = append(data, s.transformRecord(collectionType, record))
	}

	meta := &Meta{
		Total: total,
		Page:  page,
		Limit: limit,
	}

	links := s.createPaginationLinks(collectionType, meta, queryParams)

	return JSONResponse{
		Data:  data,
		Meta:  meta,
		Links: links,
	}
}

func (s *JSONAPISerializer) SerializeResource(collectionType string, record domain.Record) ResourceObject {
	return s.transformRecord(collectionType, record)
}

func (s *JSONAPISerializer) createPaginationLinks(collectionType string, meta *Meta, queryParams url.Values) *Links {
	totalPages := (meta.Total + meta.Limit - 1) / meta.Limit

	if totalPages <= 0 {
		return &Links{}
	}

	buildURL := func(p int) string {
		if len(queryParams) == 0 && p <= 0 {
			return fmt.Sprintf("%s/%s", s.baseURL, collectionType)
		}

		paramsCopy := queryParams
		paramsCopy.Set("page", strconv.Itoa(p))
		var responseUrl *url.URL
		responseUrl, _ = url.Parse(s.baseURL)
		responseUrl.Path += fmt.Sprintf("/%s", collectionType)
		responseUrl.RawQuery = paramsCopy.Encode()

		return responseUrl.String()
	}

	links := &Links{
		Self:  buildURL(meta.Page),
		First: buildURL(1),
		Last:  buildURL(totalPages),
	}

	if meta.Page > 1 {
		links.Prev = buildURL(meta.Page - 1)
	}

	if meta.Page < totalPages {
		links.Next = buildURL(meta.Page + 1)
	}

	return links
}

func (s *JSONAPISerializer) transformRecord(collectionType string, record domain.Record) ResourceObject {
	id := fmt.Sprintf("%v", record["id"])

	attributes := make(map[string]any)

	for key, value := range record {
		if key == "id" {
			continue
		}
		attributes[key] = value
	}

	return ResourceObject{
		Type:       collectionType,
		ID:         id,
		Attributes: attributes,
	}
}
