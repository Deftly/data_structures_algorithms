# [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)
Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

#### Example 1:
```shell 
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

#### Example 2:
```shell 
Input: strs = [""]
Output: [[""]]
```

#### Example 3:
```shell 
Input: strs = ["a"]
Output: [["a"]]
```

<br>

#### Constraints:
- <code>1 <= strs.length <= 10<sup>4</sup></code>
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.
