<p>给你一个整数数组&nbsp;<code>arr</code>， 对于元素 <code>x</code> ，只有当 <code>x + 1</code> 也在数组&nbsp;<code>arr</code> 里时，才能记为 <code>1</code> 个数。</p>

<p>如果数组&nbsp;<code>arr</code> 里有重复的数，每个重复的数单独计算。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>arr = [1,2,3]
<strong>输出：</strong>2
<strong>解释：</strong>1 和 2 被计算次数因为 2 和 3 在数组 arr 里。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>arr = [1,1,3,3,5,5,7,7]
<strong>输出：</strong>0
<strong>解释：</strong>所有的数都不算, 因为数组里没有 2、4、6、8。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= arr.length &lt;= 1000</code></li> 
 <li><code>0 &lt;= arr[i] &lt;= 1000</code></li> 
</ul>

<div><div>Related Topics</div><div><li>数组</li><li>哈希表</li></div></div><br><div><li>👍 11</li><li>👎 0</li></div>