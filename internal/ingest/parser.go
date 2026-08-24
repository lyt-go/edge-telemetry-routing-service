package ingest

import "edgetelemetry/internal/model"

type Parser struct {
	buffer []model.Reading
}

func (p *Parser) Parse(input []model.Reading) []model.Reading {
	p.buffer = append(p.buffer[:0], input...)
	result := make([]model.Reading, len(p.buffer))
	copy(result, p.buffer)
	return result
}
