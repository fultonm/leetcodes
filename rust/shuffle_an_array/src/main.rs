use rand::Rng;

fn main() {
    let my_vec = vec![1, 2, 3, 4, 5];
    println!("Running array shuffle with {:?}", my_vec);
    let obj = Solution::new(my_vec);
    let ret_2: Vec<i32> = obj.shuffle();
    println!("Shuffled: {:?}", ret_2);
    let ret_1: Vec<i32> = obj.reset();
    println!("Reset: {:?}", ret_1);
}

struct Solution {
    nums: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Solution {
    fn new(nums: Vec<i32>) -> Self {
        Solution { nums }
    }

    /** Resets the array to its original configuration and return it. */
    fn reset(&self) -> Vec<i32> {
        self.nums.to_owned()
    }

    /** Returns a random shuffling of the array. */
    fn shuffle(&self) -> Vec<i32> {
        let mut nums = self.nums.to_owned();
        let mut shuffled: Vec<i32> = Vec::with_capacity(nums.len());
        let mut rng = rand::thread_rng();
        while nums.len() > 0 {
            let random_index: usize = rng.gen_range(0..nums.len());
            shuffled.push(nums[random_index]);
            nums = [&nums[0..random_index], &nums[random_index + 1..nums.len()]].concat();
        }
        shuffled
    }
}
