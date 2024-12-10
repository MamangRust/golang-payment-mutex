package randomvcc

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func RandomCardNumber() (int64, error) {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	randomVirtualCreditCard := fmt.Sprintf("%.16f", rand.Float64())[2:18]
	visaCreditCard, err := strconv.ParseInt("4"+randomVirtualCreditCard, 10, 64)
	if err != nil {
		return 0, err
	}

	pattern := regexp.MustCompile(`\d{16}`)
	ccNumber := pattern.FindString(strconv.FormatInt(visaCreditCard, 10))

	result, err := strconv.ParseInt(ccNumber, 10, 64)

	if err != nil {
		return 0, err
	}

	return result, nil

}
