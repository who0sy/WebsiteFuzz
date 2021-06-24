package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"golang.org/x/net/html"
)

// DOMTreeNode HTML DOM Tree Node
type DOMTreeNode struct {
	Name    string
	ID, PID int
	Attrs   map[string]string
}

const (
	dsimension       = 5000
	initialWeight    = 1.0
	attenuationRatio = 0.6
)

func hashAbs(s string) int64 {
	h := md5.New()
	r := h.Sum([]byte(s))
	n := new(big.Int)
	n, _ = n.SetString(fmt.Sprintf("%x", r), 16)
	return int64(math.Abs(float64(n.Int64())))
}

func createNode(tree []*DOMTreeNode, pid, id int, name string, attrs map[string]string) []*DOMTreeNode {
	w := new(DOMTreeNode)
	w.PID = pid
	w.ID = id
	w.Name = name
	w.Attrs = attrs
	tree = append(tree, w)
	return tree
}

func getFeature(node *DOMTreeNode) string {
	res := node.Name
	for k, v := range node.Attrs {
		res += "|" + k + ":" + v
	}
	return res
}

func calculateWeight(tree []*DOMTreeNode, node *DOMTreeNode, featureHash int64) float64 {
	brotherNodeCount := 0
	depth := getNodeDepth(tree, node.ID, 0)
	for _, brotherNode := range tree {
		if brotherNode.PID == node.PID {
			if hashAbs(getFeature(brotherNode)) == featureHash {
				brotherNodeCount++
			}
		}
	}
	nodeWeight := initialWeight * math.Pow(attenuationRatio, float64(depth))
	if brotherNodeCount > 0 {
		nodeWeight = math.Pow(nodeWeight, float64(brotherNodeCount))
	}
	return nodeWeight
}

func getNodeDepth(tree []*DOMTreeNode, id, depth int) int {
	for _, node := range tree {
		if node.ID == 0 {
			break
		} else if node.PID == id {
			depth += getNodeDepth(tree, id, depth)
			break
		}
	}
	return depth
}

func getEigenVector(tree []*DOMTreeNode) map[int]float64 {
	domEigenVector := make(map[int]float64, 5000)
	for i := 0; i < dsimension; i++ {
		domEigenVector[i] = 0.0
	}
	for _, node := range tree {
		feature := getFeature(node)
		featureHash := hashAbs(feature)
		nodeWeight := calculateWeight(tree, node, featureHash)
		domEigenVector[int(featureHash%dsimension)] += nodeWeight
	}
	return domEigenVector
}

func calculateSimilarity(domEigenVector1, domEigenVector2 map[int]float64) float64 {
	var a, b float64
	a = 0.0
	b = 0.0
	for i := 0; i < dsimension; i++ {
		a += math.Abs(domEigenVector1[i] - domEigenVector2[i])
		if domEigenVector1[i] != 0 && domEigenVector2[i] != 0 {
			b += domEigenVector1[i] + domEigenVector2[i]
		}
	}
	similarity := math.Abs(a) / b
	return similarity
}

func parserHTML(content string) []*DOMTreeNode {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	var tree = []*DOMTreeNode{}
	var id = 0
	var walk func(*html.Node, int)
	walk = func(n *html.Node, pid int) {
		nid := 0
		if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
			attr := make(map[string]string)
			for _, a := range n.Attr {
				attr[a.Key] = a.Val
			}
			nid = id
			tree = createNode(tree, pid, nid, n.Data, attr)
			id++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c, nid)
		}
	}
	walk(doc, id)

	return tree
}

// GetSimilar - Get html similarity
func GetSimilar(doc1, doc2 string) (bool, float64) {
	tree1 := parserHTML(doc1)
	tree2 := parserHTML(doc2)

	domEigenVector1 := getEigenVector(tree1)
	domEigenVector2 := getEigenVector(tree2)

	value := calculateSimilarity(domEigenVector1, domEigenVector2)
	if value > 0.2 {
		return false, value
	}
	return true, value
}
