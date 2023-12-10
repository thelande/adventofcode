#
# Create base files for a new day.
#
[CmdletBinding()]
param (
    [Parameter(Mandatory = $true)]
    [int]
    $Year,

    # The day number to add
    [Parameter(Mandatory = $true)]
    [int]
    $Day
)

$dayDir="$Year/days/day$Day"
$dayFile="$dayDir/day$Day.go"
$dayTestFile="$dayDir/day${Day}_test.go"

if ([System.IO.File]::Exists($dayFile)) {
    Write-Error "$dayFile already exists."
    Exit 1
}

if (![System.IO.File]::Exists($dayDir)) {
    New-Item -Path $dayDir -ItemType Directory
}

@"
package day$Day

import (
    "github.com/go-kit/log"
    util "github.com/thelande/adventofcode/common"
)

type Day$Day struct{}

func (d Day$Day) Part1(filename string, logger log.Logger) int64 {
    var value int64

    util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
        return nil
    })

    return value
}

func (d Day$Day) Part2(filename string, logger log.Logger) int64 {
    var value int64

    util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
        return nil
    })

    return value
}
"@ > $dayFile

@"
package day$DAY

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay${DAY}_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day$DAY{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day$DAY.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay${DAY}_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day$DAY{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day$DAY.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
"@ > $dayTestFile
