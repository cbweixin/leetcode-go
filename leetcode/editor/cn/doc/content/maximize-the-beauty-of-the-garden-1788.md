<p>有一个花园，有&nbsp;<code>n</code>&nbsp;朵花，这些花都有一个用整数表示的美观度。这些花被种在一条线上。给定一个长度为 <code>n</code> 的整数类型数组&nbsp;<code>flowers</code>&nbsp;，每一个&nbsp;<code>flowers[i]</code>&nbsp;表示第&nbsp;<code>i</code>&nbsp;朵花的美观度。</p>

<p>一个花园满足下列条件时，该花园是<strong>有效</strong>的。</p>

<ul> 
 <li>花园中至少包含两朵花。</li> 
 <li>第一朵花和最后一朵花的美观度相同。</li> 
</ul>

<p>作为一个被钦定的园丁，你可以从花园中<strong>去除</strong>任意朵花（也可以不去除任意一朵）。你想要通过一种方法移除某些花朵，使得剩下的花园变得<strong>有效</strong>。花园的美观度是其中所有剩余的花朵美观度之和。</p>

<p>返回你去除了任意朵花（也可以不去除任意一朵）之后形成的<strong>有效</strong>花园中最大可能的美观度。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<pre><strong>输入:</strong> flowers = [1,2,3,1,2]
<strong>输出:</strong> 8
<strong>解释:</strong> 你可以修整为有效花园 [2,3,1,2] 来达到总美观度 2 + 3 + 1 + 2 = 8。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入:</strong> flowers = [100,1,1,-3,1]
<strong>输出:</strong> 3
<strong>解释:</strong> 你可以修整为有效花园 [1,1,1] 来达到总美观度 1 + 1 + 1 = 3。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入:</strong> flowers = [-1,-2,0,-1]
<strong>输出:</strong> -2
<strong>解释:</strong> 你可以修整为有效花园 [-1,-1] 来达到总美观度 -1 + -1 = -2。
</pre>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul> 
 <li><code>2 &lt;= flowers.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>-10<sup>4</sup> &lt;= flowers[i] &lt;= 10<sup>4</sup></code></li> 
 <li>去除一些花朵（可能没有）后，是有可能形成一个有效花园的。</li> 
</ul>

<div><div>Related Topics</div><div><li>贪心</li><li>数组</li><li>前缀和</li></div></div><br><div><li>👍 4</li><li>👎 0</li></div>