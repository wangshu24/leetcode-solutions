type ProblemContext struct {
    Amount []int
    Bob    int
    Tree   [][]int
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
    n := len(edges) + 1
    tree := make([][]int, n)

    // Construct adjacency list
    for _, e := range edges {
        tree[e[0]] = append(tree[e[0]], e[1])
        tree[e[1]] = append(tree[e[1]], e[0])
    }

    ctx := &ProblemContext{Amount: amount, Bob: bob, Tree: tree}
    aliceProfit, _ := dpRec(0, -1, 0, ctx)
    return aliceProfit
}

// dpRec computes Alice's max profit and Bob's distance from `node`
func dpRec(node int, parent int, currentDepth int, ctx *ProblemContext) (aliceProfit int, bobDistance int) {
    // Case: Leaf node (not root)
    if len(ctx.Tree[node]) == 1 && currentDepth != 0 {
        if node != ctx.Bob {
            aliceProfit = ctx.Amount[node]
            bobDistance = len(ctx.Tree) // Effectively ∞
        }
        return
    }

    bobDistance = len(ctx.Tree) // Set as ∞ initially
    aliceProfit = math.MinInt
    if node == ctx.Bob {
        bobDistance = 0
    }

    // Traverse child nodes
    for _, neighbor := range ctx.Tree[node] {
        if neighbor != parent {
            nProfit, nBobDistance := dpRec(neighbor, node, currentDepth+1, ctx)
            bobDistance = min(bobDistance, nBobDistance+1)
            aliceProfit = max(aliceProfit, nProfit)
        }
    }

    // Adjust Alice's profit based on Bob's position
    if currentDepth < bobDistance {
        aliceProfit += ctx.Amount[node] // Alice takes full amount
    } else if currentDepth == bobDistance {
        aliceProfit += ctx.Amount[node] / 2 // Alice splits with Bob
    }

    return
}