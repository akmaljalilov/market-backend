package utils

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

// NormalizeUsername converts input to lowercase, removes special chars, and replaces spaces with underscores
func NormalizeUsername(name string) string {
	// 1. trim spaces
	name = strings.TrimSpace(name)

	// 2. lowercase
	name = strings.ToLower(name)

	// 3. replace multiple spaces with a single underscore
	space := regexp.MustCompile(`\s+`)
	name = space.ReplaceAllString(name, "_")

	// 4. remove non-alphanumeric chars except underscore
	clean := regexp.MustCompile(`[^a-z0-9_]`)
	name = clean.ReplaceAllString(name, "")

	return name
}

func NumericToInt64(n pgtype.Numeric) (int64, error) {
	if n.Int == nil {
		return 0, fmt.Errorf("numeric value has no integer")
	}

	// Numeric = Int × 10^Exp
	if n.Exp > 0 {
		return 0, fmt.Errorf("numeric value has fractional precision")
	}

	divisor := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(-n.Exp)),
		nil,
	)

	result := new(big.Int).Quo(n.Int, divisor)

	if !result.IsInt64() {
		return 0, fmt.Errorf("numeric value overflows int64")
	}

	return result.Int64(), nil
}
