<p><strong><a href="https://en.wikipedia.org/wiki/Binary_expression_tree" target="_blank">二叉表达式树</a></strong>是一种表达算术表达式的二叉树。二叉表达式树中的每一个节点都有零个或两个子节点。&nbsp;叶节点（有 0 个子节点的节点）表示操作数，非叶节点（有 2 个子节点的节点）表示运算符。在本题中，我们只考虑 <code>'+'</code> 运算符（即加法）。</p>

<p>给定两棵二叉表达式树的根节点&nbsp;<code>root1</code>&nbsp;和&nbsp;<code>root2</code>&nbsp;。<em>如果两棵二叉表达式树等价</em>，返回&nbsp;<code>true</code>&nbsp;，否则返回&nbsp;<code>false</code>&nbsp;。</p>

<p>当两棵二叉搜索树中的变量取任意值，<strong>分别求得的值都相等</strong>时，我们称这两棵二叉表达式树是等价的。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<b>输入：</b> root1 = [x], root2 = [x]
<b>输出：</b> true
</pre>

<p><strong>示例 2:</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2020/10/04/tree1.png" style="width: 211px; height: 131px;" /></strong></p>

<pre>
<b>输入：</b>root1 = [+,a,+,null,null,b,c], root2 = [+,+,a,b,c]
<b>输出：</b>true
<span><code><span style=""><b>解释：</b></span>a + (b + c) == (b + c) + a</code></span></pre>

<p><strong>示例 3:</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2020/10/04/tree2.png" style="width: 211px; height: 131px;" /></strong></p>

<pre>
<b>输入：</b> root1 = [+,a,+,null,null,b,c], root2 = [+,+,a,b,d]
<b>输出：</b> false
<b>解释：</b> <span><code>a + (b + c) != (b + d) + a</code></span>
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li>两棵树中的节点个数相等，且节点个数为范围&nbsp;<code>[1, 4999]</code>&nbsp;内的奇数。</li> 
 <li><code>Node.val</code>&nbsp;是&nbsp;<code>'+'</code>&nbsp;或小写英文字母。</li> 
 <li>给定的树<strong>保证</strong>是有效的二叉表达式树。</li> 
</ul>

<p>&nbsp;</p>

<p><b>进阶：</b>当你的答案需同时支持&nbsp;<code>'-'</code>&nbsp;运算符（减法）时，你该如何修改你的答案</p>

<div><div>Related Topics</div><div><li>树</li><li>深度优先搜索</li><li>二叉树</li></div></div><br><div><li>👍 8</li><li>👎 0</li></div>