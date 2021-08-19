fn main() {
    println!("Hello, world!");
    let mut m1 = vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]];
    let mut m2 = vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]];
    print_matrix(&m2);
    Solution::set_zeroes(&mut m2);
    print_matrix(&m2);
}

struct Solution {}

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let mut zcols = Vec::<usize>::new();
        let mut zrows = Vec::<usize>::new();
        let m = matrix.len();
        let n = matrix[0].len();
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    if !zcols.contains(&i) {
                        zcols.push(i);
                    }
                    if !zrows.contains(&j) {
                        zrows.push(j);
                    }
                }
            }
        }
        zcols.iter().for_each(|c| {
            for j in 0..n {
                matrix[*c][j] = 0;
            }
        });
        zrows.iter().for_each(|r| {
            for i in 0..m {
                matrix[i][*r] = 0;
            }
        });
    }
}

fn print_matrix(matrix: &Vec<Vec<i32>>) {
    println!();
    for i in 0..matrix.len() {
        println!("{:?}", matrix[i]);
    }
}
