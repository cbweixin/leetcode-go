<p>在数组&nbsp;<code>A</code>&nbsp;和数组 <code>B</code>&nbsp;中，对于第一个满足 <code>A[i] != B[i]</code>&nbsp;的索引&nbsp;<code>i</code>&nbsp;，当 <code>A[i] &gt; B[i]</code>&nbsp;时，数组 <code>A</code> 大于数组 <code>B</code>。</p>

<p>例如，对于索引从 <code>0</code> 开始的数组：</p>

<ul> 
 <li><code>[1,3,2,4] &gt; [1,2,2,4]</code>&nbsp;，因为在索引&nbsp;<code>1</code>&nbsp;上，&nbsp;<code>3 &gt; 2</code>。</li> 
 <li><code>[1,4,4,4] &lt; [2,1,1,1]</code>&nbsp;，因为在索引 <code>0</code> 上，&nbsp;<code>1 &lt; 2</code>。</li> 
</ul>

<p>一个数组的子数组是原数组上的一个连续子序列。</p>

<p>给定一个包含<strong>不同</strong>整数的整数类型数组&nbsp;<code>nums</code>&nbsp;，返回&nbsp;<code>nums</code>&nbsp;中长度为 <code>k</code> 的最大子数组。</p>

<p>&nbsp;</p>

<p><b>示例 1:</b></p>

<pre><strong>输入:</strong> nums = [1,4,5,2,3], k = 3
<strong>输出:</strong> [5,2,3]
<strong>解释:</strong> 长度为 3 的子数组有： [1,4,5]、 [4,5,2] 和 [5,2,3]。
在这些数组中， [5,2,3] 是最大的。</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>输入:</strong> nums = [1,4,5,2,3], k = 4
<strong>输出:</strong> [4,5,2,3]
<strong>解释:</strong> 长度为 4 的子数组有： [1,4,5,2] 和 [4,5,2,3]。
在这些数组中， [4,5,2,3] 是最大的。</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> nums = [1,4,5,2,3], k = 1
<strong>输出:</strong> [5]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= k &lt;= nums.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li> 
 <li><code>nums</code>&nbsp;中的所有整数都是<strong>不同</strong>的。</li> 
</ul>

<p>&nbsp;</p> 
<b>进阶：</b>如果允许&nbsp;
<code>nums</code>&nbsp;中存在相同元素，你该如何解决该问题？

<div><div>Related Topics</div><div><li>贪心</li><li>数组</li></div></div><br><div><li>👍 14</li><li>👎 0</li></div>