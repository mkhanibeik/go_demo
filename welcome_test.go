package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return math.Abs(v1-v2) <= 0.001
}

func TestSqrt(t *testing.T) {
	val, err := sqrt(2)

	if err != nil {
		t.Fatalf("error in calcularion - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

func TestMoreSqrt(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			val, err := sqrt(tc.value)

			if err != nil {
				t.Fatalf("error in calcularion - %s", err)
			}

			if !almostEqual(val, tc.expected) {
				t.Fatalf("bad value - %f != %f", val, tc.expected)
			}
		})
	}
}

func TestFromCSV(t *testing.T) {
	file, err := os.Open("sqrTest.csv")
	if err != nil {
		t.Fatalf("csv file not found! - %s", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[0])
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[0])
		}

		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			val, err := sqrt(val)

			if err != nil {
				t.Fatalf("error in calcularion - %s", err)
			}

			if !almostEqual(val, expected) {
				t.Fatalf("bad value - %f != %f", val, expected)
			}
		})
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
