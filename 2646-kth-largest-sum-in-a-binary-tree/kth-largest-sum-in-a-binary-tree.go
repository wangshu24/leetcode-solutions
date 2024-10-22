/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	// Slice to store the sum of each level
	var levelSums []int64

	// If the tree is empty, return -1
	if root == nil {
		return -1
	}

	// Queue for level-order traversal
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		// Variable to store the sum of the current level
		var levelSum int64

		// Number of nodes at the current level
		pending := queue.Len()
		for i := 0; i < pending; i++ {
			// Process each node at the current level
			node := queue.Remove(queue.Front()).(*TreeNode)

			// Add the node's value to the level sum
			levelSum += int64(node.Val)

			// Add child nodes to the queue for the next level
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		// Add the sum of the current level to the slice
		levelSums = append(levelSums, levelSum)
	}

	// If there are fewer levels than k, return -1
	if len(levelSums) < k {
		return -1
	}

	// Sort the levelSums slice in descending order
	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})

	// Return the kth largest level sum
	return levelSums[k-1]
}