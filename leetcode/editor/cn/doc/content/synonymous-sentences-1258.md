<p>给你一个近义词表&nbsp;<code>synonyms</code> 和一个句子&nbsp;<code>text</code>&nbsp;，&nbsp;<code>synonyms</code> 表中是一些近义词对 ，你可以将句子&nbsp;<code>text</code> 中每个单词用它的近义词来替换。</p>

<p>请你找出所有用近义词替换后的句子，按&nbsp;<strong>字典序排序</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：
</strong>synonyms = [[&quot;happy&quot;,&quot;joy&quot;],[&quot;sad&quot;,&quot;sorrow&quot;],[&quot;joy&quot;,&quot;cheerful&quot;]],
text = &quot;I am happy today but was sad yesterday&quot;
<strong>输出：
</strong>[&quot;I am cheerful today but was sad yesterday&quot;,
&quot;I am cheerful today but was sorrow yesterday&quot;,
&quot;I am happy today but was sad yesterday&quot;,
&quot;I am happy today but was sorrow yesterday&quot;,
&quot;I am joy today but was sad yesterday&quot;,
&quot;I am joy today but was sorrow yesterday&quot;]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;=&nbsp;synonyms.length &lt;= 10</code></li>
	<li><code>synonyms[i].length == 2</code></li>
	<li><code>synonyms[0] != synonyms[1]</code></li>
	<li>所有单词仅包含英文字母，且长度最多为&nbsp;<code>10</code> 。</li>
	<li><code>text</code>&nbsp;最多包含&nbsp;<code>10</code> 个单词，且单词间用单个空格分隔开。</li>
</ul>
<div><div>Related Topics</div><div><li>并查集</li><li>数组</li><li>哈希表</li><li>字符串</li><li>回溯</li></div></div><br><div><li>👍 28</li><li>👎 0</li></div>