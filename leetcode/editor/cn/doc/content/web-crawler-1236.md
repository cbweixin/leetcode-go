<p>给定一个链接&nbsp;<code>startUrl</code> 和一个接口&nbsp;<code>HtmlParser</code>&nbsp;，请你实现一个网络爬虫，以实现爬取同&nbsp;<code>startUrl</code>&nbsp;拥有相同&nbsp;<strong>域名标签&nbsp;</strong>的全部链接。该爬虫得到的全部链接可以&nbsp;<strong>任何顺序&nbsp;</strong>返回结果。</p>

<p>你的网络爬虫应当按照如下模式工作：</p>

<ul>
	<li>自链接&nbsp;<code>startUrl</code>&nbsp;开始爬取</li>
	<li>调用&nbsp;<code>HtmlParser.getUrls(url)</code>&nbsp;来获得链接<code>url</code>页面中的全部链接</li>
	<li>同一个链接最多只爬取一次</li>
	<li>只输出&nbsp;<strong>域名&nbsp;</strong>与<strong>&nbsp;</strong><code>startUrl</code>&nbsp;<strong>相同&nbsp;</strong>的链接集合</li>
</ul>

<p><img alt="" src="https://assets.leetcode.com/uploads/2019/08/13/urlhostname.png" style="height: 164px; width: 600px;"></p>

<p>如上所示的一个链接，其域名为&nbsp;<code>example.org</code>。简单起见，你可以假设所有的链接都采用&nbsp;<strong>http协议&nbsp;</strong>并没有指定&nbsp;<strong>端口</strong>。例如，链接&nbsp;<code>http://leetcode.com/problems</code>&nbsp;和&nbsp;<code>http://leetcode.com/contest</code>&nbsp;是同一个域名下的，而链接<code>http://example.org/test</code>&nbsp;和&nbsp;<code>http://example.com/abc</code> 是不在同一域名下的。</p>

<p><code>HtmlParser</code> 接口定义如下：&nbsp;</p>

<pre>interface HtmlParser {
  // 返回给定 url 对应的页面中的全部 url 。
  public List&lt;String&gt; getUrls(String url);
}</pre>

<p>下面是两个实例，用以解释该问题的设计功能，对于自定义测试，你可以使用三个变量&nbsp;&nbsp;<code>urls</code>,&nbsp;<code>edges</code>&nbsp;和&nbsp;<code>startUrl</code>。注意在代码实现中，你只可以访问&nbsp;<code>startUrl</code>&nbsp;，而&nbsp;<code>urls</code>&nbsp;和&nbsp;<code>edges</code>&nbsp;不可以在你的代码中被直接访问。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2019/10/23/sample_2_1497.png" style="height: 300px; width: 610px;"></p>

<pre><strong>输入：
</strong>urls = [
&nbsp; &quot;http://news.yahoo.com&quot;,
&nbsp; &quot;http://news.yahoo.com/news&quot;,
&nbsp; &quot;http://news.yahoo.com/news/topics/&quot;,
&nbsp; &quot;http://news.google.com&quot;,
&nbsp; &quot;http://news.yahoo.com/us&quot;
]
edges = [[2,0],[2,1],[3,2],[3,1],[0,4]]
startUrl = &quot;http://news.yahoo.com/news/topics/&quot;
<strong>输出：</strong>[
&nbsp; &quot;http://news.yahoo.com&quot;,
&nbsp; &quot;http://news.yahoo.com/news&quot;,
&nbsp; &quot;http://news.yahoo.com/news/topics/&quot;,
&nbsp; &quot;http://news.yahoo.com/us&quot;
]
</pre>

<p><strong>示例 2：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2019/10/23/sample_3_1497.png" style="height: 270px; width: 540px;"></strong></p>

<pre><strong>输入：</strong>
urls = [
&nbsp; &quot;http://news.yahoo.com&quot;,
&nbsp; &quot;http://news.yahoo.com/news&quot;,
&nbsp; &quot;http://news.yahoo.com/news/topics/&quot;,
&nbsp; &quot;http://news.google.com&quot;
]
edges = [[0,2],[2,1],[3,2],[3,1],[3,0]]
startUrl = &quot;http://news.google.com&quot;
<strong>输入：</strong>[&quot;http://news.google.com&quot;]
<strong>解释：</strong>startUrl 链接到所有其他不共享相同主机名的页面。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= urls.length &lt;= 1000</code></li>
	<li><code>1 &lt;= urls[i].length &lt;= 300</code></li>
	<li><code>startUrl</code>&nbsp;为&nbsp;<code>urls</code>&nbsp;中的一个。</li>
	<li>域名标签的长为1到63个字符（包括点），只能包含从&lsquo;a&rsquo;到&lsquo;z&rsquo;的ASCII字母、&lsquo;0&rsquo;到&lsquo;9&rsquo;的数字以及连字符即减号（&lsquo;-&rsquo;）。</li>
	<li>域名标签不会以连字符即减号（&lsquo;-&rsquo;）开头或结尾。</li>
	<li>关于域名有效性的约束可参考:&nbsp;&nbsp;<a href="https://en.wikipedia.org/wiki/Hostname#Restrictions_on_valid_hostnames">https://en.wikipedia.org/wiki/Hostname#Restrictions_on_valid_hostnames</a></li>
	<li>你可以假定url库中不包含重复项。</li>
</ul>
<div><div>Related Topics</div><div><li>深度优先搜索</li><li>广度优先搜索</li><li>字符串</li><li>交互</li></div></div><br><div><li>👍 18</li><li>👎 0</li></div>