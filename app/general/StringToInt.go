package general

import "strconv"

func StringToInt(s string) (int, error) {
	idInt, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}
