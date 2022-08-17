import{_ as s,o as n,c as a,d as l}from"./app.b911d425.js";const A=JSON.parse('{"title":"\u5FEB\u901F\u5F00\u59CB","description":"","frontmatter":{},"headers":[{"level":2,"title":"Step.1: \u521B\u5EFA\u4E00\u4E2A\u65B0\u7684\u9879\u76EE","slug":"step-1-\u521B\u5EFA\u4E00\u4E2A\u65B0\u7684\u9879\u76EE"},{"level":2,"title":"Step.2: \u5B89\u88C5infinite","slug":"step-2-\u5B89\u88C5infinite"},{"level":2,"title":"Step.3: \u4E00\u4E2A\u7B80\u5355\u7684 confirm \u4F7F\u7528 demo","slug":"step-3-\u4E00\u4E2A\u7B80\u5355\u7684-confirm-\u4F7F\u7528-demo"}],"relativePath":"zh/guide/getting-started.md"}'),e={name:"zh/guide/getting-started.md"},p=l(`<h1 id="\u5FEB\u901F\u5F00\u59CB" tabindex="-1">\u5FEB\u901F\u5F00\u59CB <a href="https://github.com/fzdwx/infinite/releases"><img style="display:inline;" src="https://img.shields.io/github/v/release/fzdwx/infinite.svg" alt="release"></a> <a class="header-anchor" href="#\u5FEB\u901F\u5F00\u59CB" aria-hidden="true">#</a></h1><div class="tip custom-block"><p class="custom-block-title">TIP</p><p><code>infinite</code> \u4F9D\u8D56\u4E8E go 1.18.</p></div><h2 id="step-1-\u521B\u5EFA\u4E00\u4E2A\u65B0\u7684\u9879\u76EE" tabindex="-1">Step.1: \u521B\u5EFA\u4E00\u4E2A\u65B0\u7684\u9879\u76EE <a class="header-anchor" href="#step-1-\u521B\u5EFA\u4E00\u4E2A\u65B0\u7684\u9879\u76EE" aria-hidden="true">#</a></h2><p>\u521B\u5EFA\u5E76\u8FDB\u5165\u65B0\u76EE\u5F55:</p><div class="language-shell"><button class="copy"></button><span class="lang">shell</span><pre><code><span class="line"><span style="color:#A6ACCD;">mkdir infinite-demo </span><span style="color:#89DDFF;">&amp;&amp;</span><span style="color:#A6ACCD;"> </span><span style="color:#82AAFF;">cd</span><span style="color:#A6ACCD;"> infinite-demo</span></span>
<span class="line"></span></code></pre></div><p>\u4F7F\u7528 <code>go mod</code> \u521D\u59CB\u5316\u9879\u76EE:</p><div class="language-shell"><button class="copy"></button><span class="lang">shell</span><pre><code><span class="line"><span style="color:#A6ACCD;">go mod init infinite-demo</span></span>
<span class="line"></span></code></pre></div><h2 id="step-2-\u5B89\u88C5infinite" tabindex="-1">Step.2: \u5B89\u88C5<code>infinite</code> <a class="header-anchor" href="#step-2-\u5B89\u88C5infinite" aria-hidden="true">#</a></h2><p>\u6DFB\u52A0<code>infinite</code>\u4F5C\u4E3A\u9879\u76EE\u7684\u4F9D\u8D56:</p><div class="language-shell"><button class="copy"></button><span class="lang">shell</span><pre><code><span class="line"><span style="color:#A6ACCD;">go get github.com/fzdwx/infinite</span></span>
<span class="line"></span></code></pre></div><h2 id="step-3-\u4E00\u4E2A\u7B80\u5355\u7684-confirm-\u4F7F\u7528-demo" tabindex="-1">Step.3: \u4E00\u4E2A\u7B80\u5355\u7684 <code>confirm</code> \u4F7F\u7528 demo <a class="header-anchor" href="#step-3-\u4E00\u4E2A\u7B80\u5355\u7684-confirm-\u4F7F\u7528-demo" aria-hidden="true">#</a></h2><p>\u65B0\u5EFA\u4E00\u4E2A<code>main.go</code>\u6587\u4EF6\uFF0C\u5E76\u5C06\u4E0B\u9762\u4EE3\u7801\u590D\u5236\u5230\u6587\u4EF6\u4E2D:</p><div class="language-go"><button class="copy"></button><span class="lang">go</span><pre><code><span class="line"><span style="color:#89DDFF;">package</span><span style="color:#A6ACCD;"> </span><span style="color:#FFCB6B;">main</span></span>
<span class="line"></span>
<span class="line"><span style="color:#89DDFF;">import</span><span style="color:#A6ACCD;"> </span><span style="color:#89DDFF;">(</span></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;">&quot;</span><span style="color:#FFCB6B;">fmt</span><span style="color:#89DDFF;">&quot;</span></span>
<span class="line"><span style="color:#A6ACCD;">	inf </span><span style="color:#89DDFF;">&quot;</span><span style="color:#FFCB6B;">github.com/fzdwx/infinite</span><span style="color:#89DDFF;">&quot;</span></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;">&quot;</span><span style="color:#FFCB6B;">github.com/fzdwx/infinite/components/selection/confirm</span><span style="color:#89DDFF;">&quot;</span></span>
<span class="line"><span style="color:#89DDFF;">)</span></span>
<span class="line"></span>
<span class="line"><span style="color:#89DDFF;">func</span><span style="color:#A6ACCD;"> </span><span style="color:#82AAFF;">main</span><span style="color:#89DDFF;">()</span><span style="color:#A6ACCD;"> </span><span style="color:#89DDFF;">{</span></span>
<span class="line"></span>
<span class="line"><span style="color:#A6ACCD;">	val</span><span style="color:#89DDFF;">,</span><span style="color:#A6ACCD;"> _ </span><span style="color:#89DDFF;">:=</span><span style="color:#A6ACCD;"> inf</span><span style="color:#89DDFF;">.</span><span style="color:#82AAFF;">NewConfirmWithSelection</span><span style="color:#89DDFF;">(</span></span>
<span class="line"><span style="color:#A6ACCD;">		confirm</span><span style="color:#89DDFF;">.</span><span style="color:#82AAFF;">WithDefaultYes</span><span style="color:#89DDFF;">(),</span></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;">).</span><span style="color:#82AAFF;">Display</span><span style="color:#89DDFF;">()</span></span>
<span class="line"></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;font-style:italic;">if</span><span style="color:#A6ACCD;"> val </span><span style="color:#89DDFF;">{</span></span>
<span class="line"><span style="color:#A6ACCD;">		fmt</span><span style="color:#89DDFF;">.</span><span style="color:#82AAFF;">Println</span><span style="color:#89DDFF;">(</span><span style="color:#89DDFF;">&quot;</span><span style="color:#C3E88D;">yes, you are.</span><span style="color:#89DDFF;">&quot;</span><span style="color:#89DDFF;">)</span></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;">}</span><span style="color:#A6ACCD;"> </span><span style="color:#89DDFF;font-style:italic;">else</span><span style="color:#A6ACCD;"> </span><span style="color:#89DDFF;">{</span></span>
<span class="line"><span style="color:#A6ACCD;">		fmt</span><span style="color:#89DDFF;">.</span><span style="color:#82AAFF;">Println</span><span style="color:#89DDFF;">(</span><span style="color:#89DDFF;">&quot;</span><span style="color:#C3E88D;">no,you are not.</span><span style="color:#89DDFF;">&quot;</span><span style="color:#89DDFF;">)</span></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#89DDFF;">}</span></span>
<span class="line"><span style="color:#89DDFF;">}</span></span>
<span class="line"></span></code></pre></div><p>\u8FD0\u884C\u8FD9\u4E2A\u9879\u76EE:</p><div class="language-shell"><button class="copy"></button><span class="lang">shell</span><pre><code><span class="line"><span style="color:#A6ACCD;">go run </span><span style="color:#82AAFF;">.</span></span>
<span class="line"></span></code></pre></div>`,15),o=[p];function t(c,i,r,D,F,y){return n(),a("div",null,o)}var C=s(e,[["render",t]]);export{A as __pageData,C as default};
