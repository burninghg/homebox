package v1

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/repo"
)

func TestBuildLocationLabelDescription(t *testing.T) {
	parent := repo.LocationSummary{ID: uuid.New(), Name: "Garage"}

	tests := []struct {
		name     string
		location repo.LocationOut
		want     string
	}{
		{
			name: "no parent or description",
			location: repo.LocationOut{
				LocationSummary: repo.LocationSummary{Name: "Shed"},
			},
			want: "Homebox Location",
		},
		{
			name: "with parent",
			location: repo.LocationOut{
				Parent:          &parent,
				LocationSummary: repo.LocationSummary{Name: "Workbench"},
			},
			want: "Homebox Location\nParent: Garage",
		},
		{
			name: "with description",
			location: repo.LocationOut{
				LocationSummary: repo.LocationSummary{
					Name:        "Attic",
					Description: "Seasonal storage",
				},
			},
			want: "Homebox Location\nSeasonal storage",
		},
		{
			name: "with parent and description",
			location: repo.LocationOut{
				Parent: &parent,
				LocationSummary: repo.LocationSummary{
					Name:        "Attic",
					Description: "Seasonal storage",
				},
			},
			want: "Homebox Location\nParent: Garage\nSeasonal storage",
		},
		{
			name: "trims whitespace",
			location: repo.LocationOut{
				Parent: &repo.LocationSummary{Name: "   Basement   "},
				LocationSummary: repo.LocationSummary{
					Name:        "Closet",
					Description: "  Supplies  ",
				},
			},
			want: "Homebox Location\nParent: Basement\nSupplies",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildLocationLabelDescription(tt.location); got != tt.want {
				t.Fatalf("buildLocationLabelDescription() = %q, want %q", got, tt.want)
			}
		})
	}
}
