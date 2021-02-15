//Alice and Bob have an undirected graph of n nodes and 3 types of edges: 
//
// 
// Type 1: Can be traversed by Alice only. 
// Type 2: Can be traversed by Bob only. 
// Type 3: Can by traversed by both Alice and Bob. 
// 
//
// Given an array edges where edges[i] = [typei, ui, vi] represents a bidirectio
//nal edge of type typei between nodes ui and vi, find the maximum number of edges
// you can remove so that after removing the edges, the graph can still be fully t
//raversed by both Alice and Bob. The graph is fully traversed by Alice and Bob if
// starting from any node, they can reach all other nodes. 
//
// Return the maximum number of edges you can remove, or return -1 if it's impos
//sible for the graph to be fully traversed by Alice and Bob. 
//
// 
// Example 1: 
//
// 
//
// 
//Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
//Output: 2
//Explanation: If we remove the 2 edges [1,1,2] and [1,1,3]. The graph will stil
//l be fully traversable by Alice and Bob. Removing any additional edge will not m
//ake it so. So the maximum number of edges we can remove is 2.
// 
//
// Example 2: 
//
// 
//
// 
//Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
//Output: 0
//Explanation: Notice that removing any edge will not make the graph fully trave
//rsable by Alice and Bob.
// 
//
// Example 3: 
//
// 
//
// 
//Input: n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
//Output: -1
//Explanation: In the current graph, Alice cannot reach node 4 from the other no
//des. Likewise, Bob cannot reach 1. Therefore it's impossible to make the graph f
//ully traversable. 
//
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 10^5 
// 1 <= edges.length <= min(10^5, 3 * n * (n-1) / 2) 
// edges[i].length == 3 
// 1 <= edges[i][0] <= 3 
// 1 <= edges[i][1] < edges[i][2] <= n 
// All tuples (typei, ui, vi) are distinct. 
// 
// Related Topics Union Find 
// 👍 288 👎 3

// 2021-02-09 16:01:27

//leetcode submit region begin(Prohibit modification and deletion)
func maxNumEdgesToRemove(n int, edges [][]int) int {

}
//leetcode submit region end(Prohibit modification and deletion)
