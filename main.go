package main

/*
func UnpuckString(s string) (string, error) {
	// проверили, что строка не пустая
	if len(s) == 0 {
		return "", nil
	}

	// провеили, что в конце нет '\'
	if s[len(s)-1] == '\\' {
		return "", errors.New("\\ at the end")
	}

	// проверили, что первый символ - не цифра
	if s[0] >= '0' && s[0] <= '9' {
		return "", errors.New("invalid string")
	}

	// результирующий массив
	var res []byte

	for i := 0; i < len(s); i++ {
		// запоминаем текущий символ
		cur := s[i]

		// если текущий - '\'
		if cur == '\\' {
			// если след - цифра
			if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				// добавляем эту цифру (экранирована)
				res = append(res, s[i+1])
				// переходим к след. символу
				i++
				// если дальше тоже цифра, а мы только что вписывали экранированную цифру
				if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
					// клонируем её count-1 раз, так как 1 раз мы её уже добавили
					count := int(s[i+1] - '0')
					if count == 0 {
						return "", errors.New("invalid count")
					}
					for j := 0; j < count-1; j++ {
						res = append(res, s[i])
					}
				}
				continue
			}
		}

		// берем следующий символ, если это цифра
		if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
			count := int(s[i+1] - '0')
			if count == 0 {
				return "", errors.New("invalid count")
			}
			// клонируем символ, стоящий перед цифрой
			for j := 0; j < count; j++ {
				res = append(res, cur)
			}
			// пропускаем цифру
			i++
		} else {
			// иначе, если следующий символ - тоже НЕ цифра, то просто добавляем
			res = append(res, cur)
		}
	}

	return string(res), nil
}
*/

func main() {

}
