package ingest

import "edgetelemetry/internal/model"

type Parser struct {
	buffer []model.Reading
}

func (p *Parser) Parse(input []model.Reading) []model.Reading {
	p.buffer = append(p.buffer[:0], input...)
	return p.buffer
}
