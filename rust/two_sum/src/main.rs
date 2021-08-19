use std::convert::TryInto;

fn main() {
    println!("Hello, world!");
    let my_nums = vec![0, 4, 2, 0];
    let my_target = 0;
    let ans = Solution::two_sum(my_nums, my_target);
    println!("ans is {:?}", ans);
}

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let (a, b) = dp(0usize, nums.len(), target, &nums);
        vec![a.try_into().unwrap(), b.try_into().unwrap()]
    }
}

fn dp(a: usize, b: usize, t: i32, n: &Vec<i32>) -> (usize, usize) {
    if a + 1 >= b {
        return (0, 0);
    }
    for i in a + 1..b {
        if n[a] + n[i] == t {
            return (a, i);
        }
    }
    dp(a + 1, b, t, n)
}
