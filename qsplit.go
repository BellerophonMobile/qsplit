package qsplit

import (
	"unicode"
)

func Split(str string) ([]string,error) {

	fields := make([]string, 0)

	runes := make([]rune, 0)
	var inquote, escaped bool

loop:
	for _,r := range(str) {

		switch inquote {
		case true:
			switch escaped {
			case true:
				switch r {
				case '"':
					runes = append(runes, r)
				case 'n':
					runes = append(runes, '\n')
				case '\\':
					runes = append(runes, '\\')
				default:
					return nil, SplitError{Type: InvalidEscapedCharacter, Data: string(r)}
				}
				escaped = false

			case false:				
				if r == '"' {
					inquote = false
					if len(runes) > 0 {
						fields = append(fields, string(runes))
						runes = make([]rune, 0)				
					}
				} else if r == '\\' {
					escaped = true
				} else {
					runes = append(runes, r)
				}
			}
			
		case false:
			if unicode.IsSpace(r) {
				if len(runes) > 0 {
					fields = append(fields, string(runes))
					runes = make([]rune, 0)				
				}
			} else if r =='#' {
				break loop
			} else if r == '"' {
				inquote = true
			} else {
				runes = append(runes, r)
			}
		}
		
	}

	if inquote {
		return nil,SplitError{Type: UnterminatedQuote}
	}
	
	if len(runes) > 0 {
		fields = append(fields, string(runes))
	}

	return fields,nil
	
}
