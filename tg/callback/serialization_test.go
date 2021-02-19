package callback

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type NameTest string

type CountryTest struct {
	Name       NameTest
	Capital    string
	Population uint
	Area       float64
}

type FooTest struct {
	ID            int
	RegionName    string
	Countries     []CountryTest
	PopulationSum uint64
}

type BarTest struct {
	ID            uint // It differs from the FooTest's type
	RegionName    string
	Countries     []CountryTest
	PopulationSum uint64
}

func TestSerialization(t *testing.T) {
	addCallback(FooTest{})
	addCallback(BarTest{})

	sample := FooTest{
		ID:         123,
		RegionName: "CIS",
		Countries: []CountryTest{
			{
				Name:       "Russian Federation",
				Capital:    "Moscow",
				Population: 3000,
				Area:       312.12,
			},
			{
				Name:       "Belarus",
				Capital:    "Minsk",
				Population: 1000,
				Area:       56.123,
			},
		},
		PopulationSum: 4000,
	}

	encoded := Marshal(sample)
	log.WithField("encoded", encoded).Info("an encoded callback")

	value, ok := Unmarshal(encoded).(FooTest)
	assert.True(t, ok)
	assert.Equal(t, sample, value)

	_, ok = Unmarshal(encoded).(BarTest)
	assert.False(t, ok)
}
