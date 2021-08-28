fn main() {
    let tree_array = vec![Some(3), Some(1), Some(4), Some(3), None, Some(1), Some(5)];
    //let tree_array = vec![Some(3), Some(3), None, Some(4), Some(2)];
    //let tree_array = vec![Some(1)];
    //let tree_array = vec![Some(9), None, Some(3), Some(6)];
    let root = create_tree(&tree_array, 0);
    let good_nodes = Solution::good_nodes(root);
    println!("Good nodes: {0}", good_nodes);
}

fn create_tree(input: &Vec<Option<i32>>, k: i32) -> Option<Rc<RefCell<TreeNode>>> {
    let node_value;
    match input[k as usize] {
        Some(val) => {
            node_value = val;
        }
        None => {
            return None;
        }
    }
    let mut parent_node = TreeNode::new(node_value);
    if k + k + 1 < input.len() as i32 {
        parent_node.left = create_tree(input, k + k + 1);
    }
    if k + k + 2 < input.len() as i32 {
        parent_node.right = create_tree(input, k + k + 2);
    }
    Some(Rc::new(RefCell::new(parent_node)))
}

//Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(r) = &root {
            let val = r.borrow().val;
            Solution::dfs(&root, val)
        } else {
            0
        }
    }

    fn dfs(node: &Option<Rc<RefCell<TreeNode>>>, max: i32) -> i32 {
        if let Some(n) = node {
            let val = n.borrow().val;
            let m = max.max(val);
            (if m == val { 1 } else { 0 })
                + Solution::dfs(&n.borrow().left, m)
                + Solution::dfs(&n.borrow().right, m)
        } else {
            0
        }
    }
}
