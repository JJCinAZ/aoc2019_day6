package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	test1Buf = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	test2Buf = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
)

func Test_getDepth(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{strings.NewReader(test1Buf)}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildMap(tt.args.r)
			if got := getDepth(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllDepths(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{strings.NewReader(test1Buf)}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildMap(tt.args.r)
			if got := getAllDepths(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllDepths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDistToChild(t *testing.T) {
	root := buildMap(strings.NewReader(test2Buf))
	you := findNode(root, "YOU")
	san := findNode(root, "SAN")
	fmt.Println(findDistToChild(root, you))
	fmt.Println(findDistToChild(root, san))
}

func Test_Part1(t *testing.T) {
	if file, err := os.Open("data.txt"); err != nil {
		t.Error(err)
	} else {
		root := buildMap(file)
		fmt.Println("PART1: ", getAllDepths(root))
	}
}

func Test_Part2(t *testing.T) {
	if file, err := os.Open("data.txt"); err != nil {
		t.Error(err)
	} else {
		root := buildMap(file)
		you := findNode(root, "YOU")
		san := findNode(root, "SAN")
		fmt.Println("PART2: ", findDistToNode(you.Parent, san.Parent))
	}
}
