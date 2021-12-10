---
title: Coding interview question patterns
author: Kelvin Wangonya
pinned: true
last_updated: 2021-11-18
tags:
  - cheatsheet
  - reference
---

<table>
  <tr>
    <th>Problem</th>
    <th>Category</th>
    <th>Notes</th>
  </tr>
  <tr>
    <td class="easy"><a href="https://leetcode.com/problems/two-sum/">Two Sum</a></td>
    <td>Arrays, Hash Table</td>
    <td>
      <ul>
        <li>create hash map <code>M</code> to keep track of value:index</li>
        <li>for each <code>num in nums</code>:</li>
        <li>get <code>diff = target - nums</code></li>
        <li>if <code>diff</code> exists as a key in <code>M</code>, return value of <code>M.diff</code> and index of current <code>num</code></li>
        <li>else, add <code>num:num index</code> to <code>M</code>
      </ul>
    </td>
  </tr>
  <tr>
    <td class="easy"><a href="https://leetcode.com/problems/best-time-to-buy-and-sell-stock/">Best Time to Buy and Sell Stock</a></td>
    <td>Arrays, DP</td>
    <td>
      <ul>
        <li><a href="https://wangonya.com/blog/kadanes-algorithm/">Kadane's Algorithm</a></li>
      </ul>
    </td>
  </tr>
  <tr>
    <td class="easy"><a href="https://leetcode.com/problems/contains-duplicate/">Contains Duplicate</a></td>
    <td>Arrays, Hash Table, Sorting</td>
    <td>
      <ul>
        <li>convert list to set and compare length of set vs original list</li>
        <li>list to set conversion is <a href="https://stackoverflow.com/a/34642209/9312256"><code>O(n)</code></a> and len check is <a href="https://stackoverflow.com/a/1115329/9312256"><code>O(1)</code></a> so overall complexity is <code>O(n)</code></li>
        <hr />
        Alternative solution:
        <li>Loop over list, adding each element to hash table</li>
        <li>On each iteration, check if element exists in hash table</li>
      </ul>
    </td>
  </tr>
</table>

<!-- row template
<tr>
  <td class="easy"><a href=""></a></td>
  <td>Arrays, Hash Table</td>
  <td>
    <ul>
      <li></li>
    </ul>
  </td>
</tr>
-->
