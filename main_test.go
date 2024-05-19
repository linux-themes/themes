package main

import (
	"log"
	"os"
	"testing"
)

func Test_Remove(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test: themes remove all"},
		{"Test: themes remove icons"},
		{"Test: themes remove themes"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			main()

			home_path, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			_, err = os.Stat(home_path + "/.icons")
			if err != nil {
				log.Fatal(err)
			}

			_, err = os.Stat(home_path + "/.themes")
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}

func Test_List(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test: themes list all"},
		{"Test: themes list icons"},
		{"Test: themes list themes"},
		{"Test: themes list official"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_Set(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test: themes set"},
		{"Test: themes set icons"},
		{"Test: themes set themes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_Install(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test: themes install"},
		{"Test: themes install url"},
		{"Test: themes install package"},
		{"Test: themes install url and package"},
		{"Test: themes install invalidurl and invalidpackage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
