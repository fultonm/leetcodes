use rand::{self, Rng};
use std::cmp::min;

fn main() {
    println!("Hello, world!");
    let mut rng = rand::thread_rng();
    let mut costs1 = Vec::new();
    for _ in 1..81 {
        let a = rng.gen_range(1000..7000);
        let b = rng.gen_range(1000..7000);
        costs1.push(vec![a, b]);
    }
    // let costs2 = vec![vec![10, 20], vec![30, 200], vec![400, 50], vec![30, 20]];

    let min_cost = Solution::two_city_sched_cost(costs1);
    println!(
        "minimum cost to distribute these candidates: ${0}",
        min_cost
    );
}

struct Solution {}

impl Solution {
    pub fn two_city_sched_cost(costs: Vec<Vec<i32>>) -> i32 {
        let mut memos: Vec<Vec<i32>> = vec![vec![-1; costs.len()]; costs.len()];
        let n = costs.len();
        let cost = cs(n, n / 2, n / 2, &costs, &mut memos);
        println!("{0}", cost);
        cost
    }
}

fn cs(n: usize, a: usize, b: usize, costs: &Vec<Vec<i32>>, memos: &mut Vec<Vec<i32>>) -> i32 {
    if memos[a][b] > -1 {
        return memos[a][b];
    } else if n == 0 {
        return 0;
    } else if a > 0 && b > 0 {
        memos[a][b] = min(
            costs[n - 1][0] + cs(n - 1, a - 1, b, costs, memos),
            costs[n - 1][1] + cs(n - 1, a, b - 1, costs, memos),
        );
        return memos[a][b];
    } else if a == 0 {
        memos[a][b] = costs[n - 1][1] + cs(n - 1, 0, b - 1, costs, memos);
        return memos[a][b];
    } else if b == 0 {
        memos[a][b] = costs[n - 1][0] + cs(n - 1, a - 1, 0, costs, memos);
        return memos[a][b];
    }
    0
}
