struct Solution {

}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Solution {

    fn new(nums: Vec<i32>) -> Self {
        
        
    }
    
    /** Resets the array to its original configuration and return it. */
    fn reset(&self) -> Vec<i32> {
        
    }
    
    /** Returns a random shuffling of the array. */
    fn shuffle(&self) -> Vec<i32> {
        
    }
}

fn main() {
    let my_vec = vec![1, 2, 3, 4, 5];

    println!("Running array shuffle with {0}", my_vec);

    let obj = Solution::new(my_vec);
    let ret_2: Vec<i32> = obj.shuffle();
    let ret_1: Vec<i32> = obj.reset();

    println!(ret_1);
    println!(ret_2);

}