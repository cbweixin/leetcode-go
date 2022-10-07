<p>你有一个整数数组&nbsp;<code>nums</code>。你只能将一个元素&nbsp;<code>nums[i]</code> 替换为&nbsp;<code>nums[i] * nums[i]</code>。</p>

<p>返回替换后的最大子数组和。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,-1,-4,-3]
<strong>输出：</strong>17
<strong>解释：</strong>你可以把-4替换为16(-4*(-4))，使nums = [2,-1,<strong>16</strong>,-3]. 现在，最大子数组和为 2 + -1 + 16 = 17.</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,-1,1,1,-1,-1,1]
<strong>输出：</strong>4
<strong>解释：</strong>你可以把第一个-1替换为1，使 nums = [1,<strong>1</strong>,1,1,-1,-1,1]. 现在，最大子数组和为 1 + 1 + 1 + 1 = 4.</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>-10<sup>4</sup>&nbsp;&lt;= nums[i] &lt;= 10<sup>4</sup></code></li> 
</ul>

<div><div>Related Topics</div><div><li>数组</li><li>动态规划</li></div></div><br><div><li>👍 22</li><li>👎 0</li></div>