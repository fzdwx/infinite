import{_ as e}from"./plugin-vue_export-helper.21dcd24c.js";import{o,c as i,a as n,b as t,d as p,e as s,r as c}from"./app.21eeb38c.js";const u={},l=p('<div align="center"><h1>infinite</h1><span>\u{1F9EC} \u7528\u4E8E\u5F00\u53D1\u4EA4\u4E92\u5F0F CLI(tui,terminal) \u7A0B\u5E8F\u7684\u7EC4\u4EF6\u5E93.</span><br><a href="https://goreportcard.com/report/github.com/fzdwx/infinite"><img src="https://goreportcard.com/badge/github.com/fzdwx/infinite" alt="go report card"></a><a href="https://github.com/fzdwx/infinite/releases"><img src="https://img.shields.io/github/v/release/fzdwx/infinite.svg?style=flat-square" alt="release"></a></div><img src="https://user-images.githubusercontent.com/65269574/183641765-e8de7441-3c4e-4008-b2a9-b2ba556ddd72.gif" alt="demo">',2),r=s("\u4E2D\u6587 | "),k={href:"https://fzdwx.github.io/infinite/en/",target:"_blank",rel:"noopener noreferrer"},d=s("English"),m=p(`<h2 id="\u7279\u6027" tabindex="-1"><a class="header-anchor" href="#\u7279\u6027" aria-hidden="true">#</a> \u7279\u6027</h2><ul><li>\u63D0\u4F9B\u4E00\u7CFB\u5217\u5F00\u7BB1\u5373\u7528\u7684\u7EC4\u4EF6 <ul><li>autocomplete</li><li>progress bar / progress-bar group</li><li>multi/single select</li><li>spinner</li><li>confirm</li><li>input</li></ul></li><li>\u652F\u6301 window/linux (\u6211\u73B0\u5728\u53EA\u6709\u8FD9\u4E24\u79CD\u64CD\u4F5C\u7CFB\u7EDF)</li><li>\u53EF\u5B9A\u5236,\u4F60\u53EF\u4EE5\u66FF\u6362\u7EC4\u4EF6\u4E2D\u7684\u67D0\u4E9B\u9009\u9879\u6216\u65B9\u6CD5\u4E3A\u4F60\u81EA\u5DF1\u7684\u5B9E\u73B0</li><li>\u53EF\u7EC4\u5408,\u4F60\u53EF\u4EE5\u5C06\u4E00\u4E2A\u6216\u591A\u4E2A\u57FA\u7840\u7EC4\u4EF6\u8054\u5408\u5728\u4E00\u8D77\u4F7F\u7528 <ul><li><code>autocomplete</code> \u7531<code>input</code> \u548C <code>selection</code> \u7EC4\u6210</li><li><code>selection</code> \u901A\u8FC7\u5D4C\u5165<code>input</code> \u6765\u5B9E\u73B0\u8FC7\u6EE4\u529F\u80FD.</li><li>...</li></ul></li></ul><h2 id="\u6700\u4F73\u5B9E\u8DF5" tabindex="-1"><a class="header-anchor" href="#\u6700\u4F73\u5B9E\u8DF5" aria-hidden="true">#</a> \u6700\u4F73\u5B9E\u8DF5</h2><ol><li>\u901A\u8FC7\u6D88\u606F\u6765\u66F4\u65B0\u72B6\u6001,\u4E5F\u5C31\u662F\u901A\u8FC7<code>program.Send(msg)</code>\u6765\u53D1\u9001\u6D88\u606F,<code>Update</code>\u76D1\u542C\u5E76\u8FDB\u884C\u72B6\u6001\u66F4\u65B0,\u6700\u540E\u901A\u8FC7<code>View</code>\u6765\u53CD\u9988\u7ED3\u679C.</li><li>...</li></ol><h2 id="\u5B89\u88C5" tabindex="-1"><a class="header-anchor" href="#\u5B89\u88C5" aria-hidden="true">#</a> \u5B89\u88C5</h2><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>go get github.com/fzdwx/infinite
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="\u4F7F\u7528\u6848\u4F8B" tabindex="-1"><a class="header-anchor" href="#\u4F7F\u7528\u6848\u4F8B" aria-hidden="true">#</a> \u4F7F\u7528\u6848\u4F8B</h2><h3 id="combined-demo" tabindex="-1"><a class="header-anchor" href="#combined-demo" aria-hidden="true">#</a> Combined demo</h3><p>\u4E00\u4E2A <code>progress</code> \u548C <code>spinner</code> \u7EC4\u5408\u4F7F\u7528\u7684demo</p><p><img src="https://user-images.githubusercontent.com/65269574/184496950-dbc246e7-5199-4e85-8167-1292b6eeb574.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;errors&quot;</span>
	tea <span class="token string">&quot;github.com/charmbracelet/bubbletea&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/pkg/strx&quot;</span>
	<span class="token string">&quot;time&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	total <span class="token operator">:=</span> <span class="token number">10</span>
	spinner <span class="token operator">:=</span> components<span class="token punctuation">.</span><span class="token function">NewSpinner</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	spinner<span class="token punctuation">.</span>Prompt <span class="token operator">=</span> strx<span class="token punctuation">.</span>Space <span class="token operator">+</span> spinner<span class="token punctuation">.</span>Prompt
	progress <span class="token operator">:=</span> components<span class="token punctuation">.</span><span class="token function">NewProgress</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">WithTotal</span><span class="token punctuation">(</span><span class="token function">int64</span><span class="token punctuation">(</span>total<span class="token punctuation">)</span><span class="token punctuation">)</span>

	<span class="token function">NewComponent</span><span class="token punctuation">(</span>spinner<span class="token punctuation">,</span> progress<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
			progress<span class="token punctuation">.</span><span class="token function">IncrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
			<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>

		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
			progress<span class="token punctuation">.</span><span class="token function">DecrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
			<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>

		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
			progress<span class="token punctuation">.</span><span class="token function">IncrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
			<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">type</span> Component <span class="token keyword">struct</span> <span class="token punctuation">{</span>
	spinner  <span class="token operator">*</span>components<span class="token punctuation">.</span>Spinner
	progress <span class="token operator">*</span>components<span class="token punctuation">.</span>Progress
	<span class="token operator">*</span>components<span class="token punctuation">.</span>StartUp
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">NewComponent</span><span class="token punctuation">(</span>spinner <span class="token operator">*</span>components<span class="token punctuation">.</span>Spinner<span class="token punctuation">,</span> progress <span class="token operator">*</span>components<span class="token punctuation">.</span>Progress<span class="token punctuation">)</span> <span class="token operator">*</span>Component <span class="token punctuation">{</span>
	<span class="token keyword">return</span> <span class="token operator">&amp;</span>Component<span class="token punctuation">{</span>spinner<span class="token punctuation">:</span> spinner<span class="token punctuation">,</span> progress<span class="token punctuation">:</span> progress<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token function">Display</span><span class="token punctuation">(</span>runner <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	c<span class="token punctuation">.</span>StartUp <span class="token operator">=</span> components<span class="token punctuation">.</span><span class="token function">NewStartUp</span><span class="token punctuation">(</span>c<span class="token punctuation">)</span>
	<span class="token keyword">if</span> runner <span class="token operator">==</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">&quot;runner is null&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token function">runner</span><span class="token punctuation">(</span>c<span class="token punctuation">)</span>
		c<span class="token punctuation">.</span>progress<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		c<span class="token punctuation">.</span><span class="token function">Quit</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token keyword">return</span> c<span class="token punctuation">.</span><span class="token function">Start</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token function">Init</span><span class="token punctuation">(</span><span class="token punctuation">)</span> tea<span class="token punctuation">.</span>Cmd <span class="token punctuation">{</span>

	<span class="token keyword">return</span> tea<span class="token punctuation">.</span><span class="token function">Batch</span><span class="token punctuation">(</span>c<span class="token punctuation">.</span>spinner<span class="token punctuation">.</span><span class="token function">Init</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> c<span class="token punctuation">.</span>progress<span class="token punctuation">.</span><span class="token function">Init</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token function">Update</span><span class="token punctuation">(</span>msg tea<span class="token punctuation">.</span>Msg<span class="token punctuation">)</span> <span class="token punctuation">(</span>tea<span class="token punctuation">.</span>Model<span class="token punctuation">,</span> tea<span class="token punctuation">.</span>Cmd<span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">switch</span> msg <span class="token operator">:=</span> msg<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> tea<span class="token punctuation">.</span>KeyMsg<span class="token punctuation">:</span>
		<span class="token keyword">switch</span> msg<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token string">&quot;ctrl+c&quot;</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> c<span class="token punctuation">,</span> tea<span class="token punctuation">.</span>Quit
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>
	<span class="token boolean">_</span><span class="token punctuation">,</span> c1 <span class="token operator">:=</span> c<span class="token punctuation">.</span>spinner<span class="token punctuation">.</span><span class="token function">Update</span><span class="token punctuation">(</span>msg<span class="token punctuation">)</span>
	<span class="token boolean">_</span><span class="token punctuation">,</span> c2 <span class="token operator">:=</span> c<span class="token punctuation">.</span>progress<span class="token punctuation">.</span><span class="token function">Update</span><span class="token punctuation">(</span>msg<span class="token punctuation">)</span>

	<span class="token keyword">return</span> c<span class="token punctuation">,</span> tea<span class="token punctuation">.</span><span class="token function">Batch</span><span class="token punctuation">(</span>c1<span class="token punctuation">,</span> c2<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token function">View</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span> <span class="token punctuation">{</span>
	<span class="token keyword">return</span> strx<span class="token punctuation">.</span><span class="token function">NewFluent</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span>c<span class="token punctuation">.</span>spinner<span class="token punctuation">.</span><span class="token function">View</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Space</span><span class="token punctuation">(</span><span class="token number">4</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span>c<span class="token punctuation">.</span>progress<span class="token punctuation">.</span><span class="token function">View</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>c <span class="token operator">*</span>Component<span class="token punctuation">)</span> <span class="token function">SetProgram</span><span class="token punctuation">(</span>program <span class="token operator">*</span>tea<span class="token punctuation">.</span>Program<span class="token punctuation">)</span> <span class="token punctuation">{</span>
	c<span class="token punctuation">.</span>spinner<span class="token punctuation">.</span><span class="token function">SetProgram</span><span class="token punctuation">(</span>program<span class="token punctuation">)</span>
	c<span class="token punctuation">.</span>progress<span class="token punctuation">.</span><span class="token function">SetProgram</span><span class="token punctuation">(</span>program<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Millisecond <span class="token operator">*</span> <span class="token number">100</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="progress-group" tabindex="-1"><a class="header-anchor" href="#progress-group" aria-hidden="true">#</a> Progress group</h3><p><img src="https://user-images.githubusercontent.com/65269574/183296585-b0a56827-d9d9-4258-ad32-266ada01b1ed.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components/progress&quot;</span>
	<span class="token string">&quot;time&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	cnt <span class="token operator">:=</span> <span class="token number">10</span>

	group <span class="token operator">:=</span> progress<span class="token punctuation">.</span><span class="token function">NewGroupWithCount</span><span class="token punctuation">(</span><span class="token number">10</span><span class="token punctuation">)</span><span class="token punctuation">.</span>
		<span class="token function">AppendRunner</span><span class="token punctuation">(</span><span class="token keyword">func</span><span class="token punctuation">(</span>progress <span class="token operator">*</span>components<span class="token punctuation">.</span>Progress<span class="token punctuation">)</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			total <span class="token operator">:=</span> cnt
			cnt <span class="token operator">+=</span> <span class="token number">1</span>
			progress<span class="token punctuation">.</span><span class="token function">WithTotal</span><span class="token punctuation">(</span><span class="token function">int64</span><span class="token punctuation">(</span>total<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">.</span>
				<span class="token function">WithDefaultGradient</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

			<span class="token keyword">return</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
					progress<span class="token punctuation">.</span><span class="token function">IncrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
					<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>

				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
					progress<span class="token punctuation">.</span><span class="token function">DecrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
					<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>

				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> total<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
					progress<span class="token punctuation">.</span><span class="token function">IncrOne</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
					<span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	group<span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">sleep</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Millisecond <span class="token operator">*</span> <span class="token number">100</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="multiple-select" tabindex="-1"><a class="header-anchor" href="#multiple-select" aria-hidden="true">#</a> Multiple select</h3><p><img src="https://user-images.githubusercontent.com/65269574/183274216-d2a7af91-0581-4d13-b8c2-00b9aad5ef3a.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	inf <span class="token string">&quot;github.com/fzdwx/infinite&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/color&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components/selection/multiselect&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/style&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	input <span class="token operator">:=</span> components<span class="token punctuation">.</span><span class="token function">NewInput</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	input<span class="token punctuation">.</span>Prompt <span class="token operator">=</span> <span class="token string">&quot;Filtering: &quot;</span>
	input<span class="token punctuation">.</span>PromptStyle <span class="token operator">=</span> style<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Bold</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Italic</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Fg</span><span class="token punctuation">(</span>color<span class="token punctuation">.</span>LightBlue<span class="token punctuation">)</span>

	<span class="token boolean">_</span><span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">=</span> inf<span class="token punctuation">.</span><span class="token function">NewMultiSelect</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
		<span class="token string">&quot;Buy carrots&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy celery&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy kohlrabi&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy computer&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy something&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy car&quot;</span><span class="token punctuation">,</span>
		<span class="token string">&quot;Buy subway&quot;</span><span class="token punctuation">,</span>
	<span class="token punctuation">}</span><span class="token punctuation">,</span>
		multiselect<span class="token punctuation">.</span><span class="token function">WithFilterInput</span><span class="token punctuation">(</span>input<span class="token punctuation">)</span><span class="token punctuation">,</span>
	<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token string">&quot;select your items!&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="spinner" tabindex="-1"><a class="header-anchor" href="#spinner" aria-hidden="true">#</a> Spinner</h3><p><img src="https://user-images.githubusercontent.com/65269574/183074665-42d7d902-a56c-420c-a740-3aacc7dc922c.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	inf <span class="token string">&quot;github.com/fzdwx/infinite&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components/spinner&quot;</span>
	<span class="token string">&quot;time&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token boolean">_</span> <span class="token operator">=</span> inf<span class="token punctuation">.</span><span class="token function">NewSpinner</span><span class="token punctuation">(</span>
		spinner<span class="token punctuation">.</span><span class="token function">WithShape</span><span class="token punctuation">(</span>components<span class="token punctuation">.</span>Dot<span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token comment">//spinner.WithDisableOutputResult(),</span>
	<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token keyword">func</span><span class="token punctuation">(</span>spinner <span class="token operator">*</span>spinner<span class="token punctuation">.</span>Spinner<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
			time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Millisecond <span class="token operator">*</span> <span class="token number">100</span><span class="token punctuation">)</span>
			spinner<span class="token punctuation">.</span><span class="token function">Refreshf</span><span class="token punctuation">(</span><span class="token string">&quot;hello world %d&quot;</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>

		spinner<span class="token punctuation">.</span><span class="token function">Finish</span><span class="token punctuation">(</span><span class="token string">&quot;finish&quot;</span><span class="token punctuation">)</span>

		spinner<span class="token punctuation">.</span><span class="token function">Refresh</span><span class="token punctuation">(</span><span class="token string">&quot;is finish?&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Millisecond <span class="token operator">*</span> <span class="token number">100</span> <span class="token operator">*</span> <span class="token number">15</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="input-text" tabindex="-1"><a class="header-anchor" href="#input-text" aria-hidden="true">#</a> Input text</h3><p><img src="https://user-images.githubusercontent.com/65269574/183075959-031a068d-6f88-40a0-8b5e-f3d5bba481af.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;fmt&quot;</span>
	inf <span class="token string">&quot;github.com/fzdwx/infinite&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components/input/text&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/theme&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	i <span class="token operator">:=</span> inf<span class="token punctuation">.</span><span class="token function">NewText</span><span class="token punctuation">(</span>
		text<span class="token punctuation">.</span><span class="token function">WithPrompt</span><span class="token punctuation">(</span><span class="token string">&quot;what&#39;s your name? &quot;</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		text<span class="token punctuation">.</span><span class="token function">WithPromptStyle</span><span class="token punctuation">(</span>theme<span class="token punctuation">.</span>DefaultTheme<span class="token punctuation">.</span>PromptStyle<span class="token punctuation">)</span><span class="token punctuation">,</span>
		text<span class="token punctuation">.</span><span class="token function">WithPlaceholder</span><span class="token punctuation">(</span><span class="token string">&quot; fzdwx (maybe)&quot;</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
	<span class="token punctuation">)</span>

	<span class="token boolean">_</span> <span class="token operator">=</span> i<span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">&quot;you input: %s\\n&quot;</span><span class="token punctuation">,</span> i<span class="token punctuation">.</span><span class="token function">Value</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="confirm-with-input" tabindex="-1"><a class="header-anchor" href="#confirm-with-input" aria-hidden="true">#</a> Confirm with Input</h3><p><img src="https://user-images.githubusercontent.com/65269574/183076452-5fa73013-42de-47df-97b4-7be743d074c1.gif" alt="demo"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;fmt&quot;</span>
	inf <span class="token string">&quot;github.com/fzdwx/infinite&quot;</span>
	<span class="token string">&quot;github.com/fzdwx/infinite/components/input/confirm&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	c <span class="token operator">:=</span> inf<span class="token punctuation">.</span><span class="token function">NewConfirm</span><span class="token punctuation">(</span>
		confirm<span class="token punctuation">.</span><span class="token function">WithDefaultYes</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
		confirm<span class="token punctuation">.</span><span class="token function">WithDisplayHelp</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span>
	<span class="token punctuation">)</span>

	c<span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token keyword">if</span> c<span class="token punctuation">.</span><span class="token function">Value</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;yes, you are.&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;no,you are not.&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details><hr><h3 id="confirm-with-selection" tabindex="-1"><a class="header-anchor" href="#confirm-with-selection" aria-hidden="true">#</a> Confirm With Selection</h3><p><img src="https://user-images.githubusercontent.com/65269574/184532991-ef3f5290-ae32-4294-906e-c097c3cf8ca1.gif" alt="Image"></p><details><summary>\u4EE3\u7801</summary><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;fmt&quot;</span>
	inf <span class="token string">&quot;github.com/fzdwx/infinite&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	val<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> inf<span class="token punctuation">.</span><span class="token function">NewConfirmWithSelection</span><span class="token punctuation">(</span>
		<span class="token comment">//confirm.WithDisOutResult(),</span>
	<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Display</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>val<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></details>`,35),v={href:"https://github.com/fzdwx/infinite/tree/main/_examples",target:"_blank",rel:"noopener noreferrer"},b=s("\u6240\u6709\u793A\u4F8B"),f=n("h2",{id:"\u4F9D\u8D56",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#\u4F9D\u8D56","aria-hidden":"true"},"#"),s(" \u4F9D\u8D56")],-1),g=n("ul",null,[n("li",null,"https://github.com/charmbracelet/bubbletea"),n("li",null,"https://github.com/charmbracelet/bubbles"),n("li",null,"https://github.com/charmbracelet/lipgloss"),n("li",null,"https://github.com/muesli/termenv"),n("li",null,"https://github.com/sahilm/fuzzy"),n("li",null,"...")],-1),h={href:"https://github.com/fzdwx/infinite/network/dependencies",target:"_blank",rel:"noopener noreferrer"},w=s("\u6240\u6709\u4F9D\u8D56"),y=n("h2",{id:"\u5F00\u6E90\u534F\u8BAE",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#\u5F00\u6E90\u534F\u8BAE","aria-hidden":"true"},"#"),s(" \u5F00\u6E90\u534F\u8BAE")],-1),q=n("p",null,"MIT",-1);function x(_,S){const a=c("ExternalLinkIcon");return o(),i("div",null,[l,n("p",null,[r,n("a",k,[d,t(a)])]),m,n("p",null,[n("a",v,[b,t(a)])]),f,g,n("p",null,[n("a",h,[w,t(a)])]),y,q])}var P=e(u,[["render",x],["__file","index.html.vue"]]);export{P as default};
