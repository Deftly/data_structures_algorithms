# [379. Design Phone Directory](https://leetcode.com/problems/design-phone-directory/)
Design a phone directory that initially has `maxNumbers` empty slots that can store numbers. The directory should store numbers, check if a certain slot is empty or not, and empty a given slot.

Implement the `PhoneDirectory` class:

- `PhoneDirectory(int maxNumbers)` Initializes the phone directory with the number of available slots `maxNumbers`.
- `int get()` Provides a number that is not assigned to anyone. Returns `-1` if no number is available.
- `bool check(int number)` Returns `true` if the slot `number` is available and `false` otherwise.
- `void release(int number)` Recycles or releases the slot `number`.

#### Example 1:
```shell 
Input
["PhoneDirectory", "get", "get", "check", "get", "check", "release", "check"]
[[3], [], [], [2], [], [2], [2], [2]]
Output
[null, 0, 1, true, 2, false, null, true]
```

<br>

#### Constraints:
- <code>1 <= maxNumbers <= 10<sup>4</sup></code>
- `0 <= number < maxNumbers`
- At most <code>2 * 10<sup>4</sup></code> calls will be made to `get`, `check`, and `release`.
