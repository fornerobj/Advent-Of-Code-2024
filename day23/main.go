package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func isNeighbor(adjs map[string][]string, a, b string) bool {
    for _, n := range adjs[a] {
	if n == b {
	    return true
	}
    }
    return false
}

func sliceToString(slce []string) string {
    str := ""
    for _, node := range slce {
	str += node
    }
    return str
}

func hasT(clique []string) bool {
    for _, node := range clique {
	if string(node[0]) == "t" {
	    return true
	}
    }
    return false
}

func part1(adjs map[string][]string, nodes []string) int{
    count := 0
    cliques := make([][]string, 0)
    seen := make(map[string]bool)

    for _, node := range nodes {
	ns := adjs[node]
	if len(ns) < 2 {
	    continue
	}

	for i := 0; i < len(ns); i++ {
	    for j := i+1; j < len(ns); j++ {
		n1, n2 := ns[i], ns[j]

		if isNeighbor(adjs, n1, n2) {
		    clique := []string{node, n1, n2}
		    sort.Strings(clique)
		    cliqueString := sliceToString(clique)

		    if !seen[cliqueString] {
			seen[cliqueString] = true
			cliques = append(cliques, clique)
			if hasT(clique) {
			    count++
			}
		    }
		    
		}
	    }
	}
    }
    return count
}

//I googled about finding maximum size cliques
// and I found out about the Bron-Kerbosch alg
// This is my attempt to implement the algorithm.

func intersect(a, b []string) []string {
    set := make(map[string]bool)
    for _, node := range b {
	set[node] = true
    }
    var res []string
    for _, node := range a {
	if set[node] {
	    res = append(res, node)
	}
    }
    return res
}

func remove(slice []string, element string) []string {
    for i, node := range slice {
	if node == element {
	    return append(slice[:i], slice[i+1:]...)
	}
    }
    return slice
}

func BK(R, P, X []string, adjs map[string][]string, maxClique *[]string) {
    if len(P) == 0 && len(X) == 0 {
	if len(R) > len(*maxClique) {
	    *maxClique = append([]string{}, R...)
	}
	return
    }

    for _, node := range append([]string{}, P...) {
	neighbors := adjs[node]

	BK(append(R, node), intersect(P, neighbors), intersect(X, neighbors), adjs, maxClique)
	P = remove(P, node)
	X = append(X, node)
    }
}

func part2(adjs map[string][]string, nodes []string) string {
    R := []string{}
    P := nodes
    X := []string{}
    maxClique := []string{}

    BK(R,P,X,adjs,&maxClique)
    sort.Strings(maxClique)

    res := "\n"
    for i, node := range maxClique {
	res += node
	if i != len(maxClique)-1 {
	    res += ","
	}    
    }
    return res
}

func in(neighbors []string, node string) bool{
    for _, name := range neighbors {
	if name == node {
	    return true
	}
    }
    return false
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    adjs := make(map[string][]string)
    for _, line := range lines {
	computers := strings.Split(line, "-")
	if _, exists := adjs[computers[0]]; !exists {
	    adjs[computers[0]] = make([]string, 0)
	}
	if _, exists := adjs[computers[1]]; !exists {
	    adjs[computers[1]] = make([]string, 0)
	}
	if !in(adjs[computers[0]], computers[1]) {
	    adjs[computers[0]] = append(adjs[computers[0]], computers[1])
	}
	if !in(adjs[computers[1]], computers[0]) {
	    adjs[computers[1]] = append(adjs[computers[1]], computers[0])
	}

    }

    nodes := make([]string, len(adjs))
    i:=0
    for a := range adjs {
	nodes[i] = a
	i++
    }

    fmt.Println("part 1:", part1(adjs, nodes))
    fmt.Println("part 2:", part2(adjs, nodes))

}
