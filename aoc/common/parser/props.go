package parser

import "errors"

func (p *ParsedData) Number(index int) (number, error) {
	if index >= len(p.parsed) {
		return 0, errors.New("invalid index for parsed data")
	}

	num, ok := p.parsed[index].(number)
	if !ok {
		return 0, errors.New("no number at the specified index")
	}
	return num, nil
}

func (p *ParsedData) Numbers(index int) ([]number, error) {
	if index >= len(p.parsed) {
		return nil, errors.New("invalid index for parsed data")
	}

	anys, ok := p.parsed[index].([]any)
	if !ok {
		return nil, errors.New("no collection at the specified index")
	}

	nums := make([]number, len(anys))
	for i, v := range anys {
		nums[i], ok = v.(number)
		if !ok {
			return nil, errors.New("invalid collection type at the specified index")
		}
	}
	return nums, nil
}

func (p *ParsedData) Bytes(index int) ([]byte, error) {
	if index >= len(p.parsed) {
		return nil, errors.New("invalid index for parsed data")
	}

	nums, ok := p.parsed[index].([]byte)
	if !ok {
		return nil, errors.New("no collection at the specified index")
	}

	return nums, nil
}
