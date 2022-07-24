<p>给定两个稀疏向量，计算它们的点积（数量积）。</p>

<p>实现类&nbsp;<code>SparseVector</code>：</p>

<ul> 
 <li><code>SparseVector(nums)</code>&nbsp;以向量&nbsp;<code>nums</code>&nbsp;初始化对象。</li> 
 <li><code>dotProduct(vec)</code>&nbsp;计算此向量与&nbsp;<code>vec</code>&nbsp;的点积。</li> 
</ul>

<p><strong>稀疏向量</strong> 是指绝大多数分量为 0 的向量。你需要 <strong>高效</strong> 地存储这个向量，并计算两个稀疏向量的点积。</p>

<p><strong>进阶：</strong>当其中只有一个向量是稀疏向量时，你该如何解决此问题？</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
<strong>输出：</strong>8
<strong>解释：</strong>v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
v1.dotProduct(v2) = 1*0 + 0*3 + 0*0 + 2*4 + 3*0 = 8
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
<strong>输出：</strong>0
<strong>解释：</strong>v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
v1.dotProduct(v2) = 0*0 + 1*0 + 0*0 + 0*0 + 0*2 = 0
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [0,1,0,0,2,0,0], nums2 = [1,0,0,0,3,0,4]
<strong>输出：</strong>6
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>n == nums1.length == nums2.length</code></li> 
 <li><code>1 &lt;= n &lt;= 10^5</code></li> 
 <li><code>0 &lt;= nums1[i], nums2[i]&nbsp;&lt;= 100</code></li> 
</ul>

<div><div>Related Topics</div><div><li>设计</li><li>数组</li><li>哈希表</li><li>双指针</li></div></div><br><div><li>👍 23</li><li>👎 0</li></div>