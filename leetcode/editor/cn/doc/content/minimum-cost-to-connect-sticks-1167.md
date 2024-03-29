<p>你有一些长度为正整数的棍子。这些长度以数组<meta charset="UTF-8" />&nbsp;<code>sticks</code>&nbsp;的形式给出，<meta charset="UTF-8" />&nbsp;<code>sticks[i]</code>&nbsp;是 <code>第i个</code> 木棍的长度。</p>

<p>你可以通过支付 <code>x + y</code> 的成本将任意两个长度为 <code>x</code> 和 <code>y</code> 的棍子连接成一个棍子。你必须连接所有的棍子，直到剩下一个棍子。</p>

<p>返回以这种方式将所有给定的棍子连接成一个棍子的 <em>最小成本</em> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>sticks = [2,4,3]
<strong>输出：</strong>14
<strong>解释：</strong>从 sticks = [2,4,3] 开始。
1. 连接 2 和 3 ，费用为 2 + 3 = 5 。现在 sticks = [5,4]
2. 连接 5 和 4 ，费用为 5 + 4 = 9 。现在 sticks = [9]
所有木棍已经连成一根，总费用 5 + 9 = 14
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>sticks = [1,8,3,5]
<strong>输出：</strong>30
<strong>解释：</strong>从 sticks = [1,8,3,5] 开始。
1. 连接 1 和 3 ，费用为 1 + 3 = 4 。现在 sticks = [4,8,5]
2. 连接 4 和 5 ，费用为 4 + 5 = 9 。现在 sticks = [9,8]
3. 连接 9 和 8 ，费用为 9 + 8 = 17 。现在 sticks = [17]
所有木棍已经连成一根，总费用 4 + 9 + 17 = 30
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>sticks = [5]
<strong>输出：</strong>0
<strong>解释：</strong>只有一根木棍，不必再连接。总费用 0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= sticks.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= sticks[i] &lt;= 10<sup>4</sup></code></li>
</ul>
<div><div>Related Topics</div><div><li>贪心</li><li>数组</li><li>堆（优先队列）</li></div></div><br><div><li>👍 60</li><li>👎 0</li></div>