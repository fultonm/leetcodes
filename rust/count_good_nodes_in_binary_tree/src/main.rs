fn main() {
    //let tree_array = vec![Some(3), Some(1), Some(4), Some(3), None, Some(1), Some(5)];
    //let tree_array = vec![Some(3), Some(3), None, Some(4), Some(2)];
    //let tree_array = vec![Some(1)];
    let tree_array = vec![Some(9), None, Some(3), Some(6)];
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

use std::cell::{Ref, RefCell};
use std::rc::Rc;
impl Solution {
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        match root.as_ref() {
            Some(root_node) => {
                let root_node = (**root_node).borrow();
                let (nodes, bad_nodes) = count_bad_nodes(root_node);
                return nodes + 1 - bad_nodes;
            }
            None => {
                return 0;
            }
        }
    }
}

pub fn count_bad_nodes(parent_node: Ref<TreeNode>) -> (i32, i32) {
    let mut child_nodes = 0;
    let mut child_bad_nodes = 0;
    let (mut left_nodes, mut left_bad_nodes) = (0, 0);
    match parent_node.left.as_ref() {
        Some(left) => {
            let left_node = (**left).borrow();
            child_nodes += 1;
            if left_node.val < parent_node.val {
                child_bad_nodes += 1;
            }
            let tmp = count_bad_nodes(left_node);
            left_nodes = tmp.0;
            left_bad_nodes = tmp.1;
        }
        None => (),
    }
    let (mut right_nodes, mut right_bad_nodes) = (0, 0);
    match parent_node.right.as_ref() {
        Some(right) => {
            let right_node = (**right).borrow();
            child_nodes += 1;
            if right_node.val < parent_node.val {
                child_bad_nodes += 1;
            }
            let tmp = count_bad_nodes(right_node);
            right_nodes = tmp.0;
            right_bad_nodes = tmp.1;
        }
        None => (),
    }
    (
        child_nodes + right_nodes + left_nodes,
        child_bad_nodes + right_bad_nodes + left_bad_nodes,
    )
}
