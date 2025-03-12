pub struct Solution;

impl Solution {
    pub fn sum_of_the_digits_of_harshad_number(x: i32) -> i32 {
        match x {
            1..=100 => {
                let sum = i32::from(x.to_string().bytes().map(|b| b - b'0').sum::<u8>());
                if x % sum == 0 {
                    return sum;
                } else {
                    return -1;
                }
            }
            _ => panic!("x must be in range from 1 to 100"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn it_finds_hashrad() {
        assert_eq!(9, Solution::sum_of_the_digits_of_harshad_number(18))
    }

    #[test]
    fn it_not_finds_hashrad() {
        assert_eq!(-1, Solution::sum_of_the_digits_of_harshad_number(23))
    }
}
