use std::collections::HashMap;

fn main() {
    // println!("Hello, world!");
    // let my_vec = vec![2, 2, 2, 3, 1, 3, 2, 2, 5, 2]; //26
    // let my_vec = vec![1, 3, 2, 2, 2, 3, 4, 3, 1]; //23
    let my_vec = vec![
        1, 2, 2, 1, 1, 1, 2, 1, 1, 2, 1, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 1, 2, 2, 2, 2, 1, 2, 1,
        1, 2, 2, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2, 1, 2, 2, 1, 1, 1, 2, 2, 1, 2, 1, 2, 2, 1, 2, 1, 1,
        1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1,
        1, 1, 1, 1, 1, 2, 1, 2, 2, 1,
    ];
    let max_score = Solution::remove_boxes(my_vec);
    println!("max score is {0}", max_score);
}

struct Solution {}

impl Solution {
    pub fn remove_boxes(boxes: Vec<i32>) -> i32 {
        let l = 0usize;
        let r = boxes.len() - 1;
        let k = 0usize;
        let mut memos = HashMap::<(usize, usize, usize), usize>::new();
        let max = dp(l, r, k, &boxes, &mut memos);
        max as i32
    }
}

fn dp(
    pl: usize,
    pr: usize,
    pk: usize,
    boxes: &Vec<i32>,
    memos: &mut HashMap<(usize, usize, usize), usize>,
) -> usize {
    if pl > pr {
        return 0;
    }
    if let Some(memo) = memos.get(&(pl, pr, pk)) {
        //println!("used memo for l:{0}, r:{1}, k:{2}, = {3}", l, r, k, *memo);
        return *memo;
    }
    let mut l = pl;
    let r = pr;
    let mut k = pk;
    while l + 1 <= r && boxes[l] == boxes[l + 1] {
        l += 1;
        k += 1;
    }
    let mut o1 = (k + 1) * (k + 1) + dp(l + 1, r, 0, boxes, memos);
    for m in l + 1..r + 1 {
        if boxes[l] == boxes[m] {
            let o2 = dp(m, r, k + 1, boxes, memos) + dp(l + 1, m - 1, 0, boxes, memos);
            if o2 > o1 {
                o1 = o2;
            }
        }
    }
    (*memos).insert((pl, pr, pk), o1);
    return o1;
}
