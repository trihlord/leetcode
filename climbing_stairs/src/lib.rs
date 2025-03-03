pub struct Solution;

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        match n {
            1..=45 => {
                let mut left = 0;
                let mut right = 1;
                for _ in 1..=n {
                    right = left + right;
                    left = right - left;
                }
                right
            }
            _ => panic!("n must be in range from 1 to 45"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn climbs_one_stair() {
        assert_eq!(1, Solution::climb_stairs(1));
    }

    #[test]
    fn climbs_two_stairs() {
        assert_eq!(2, Solution::climb_stairs(2));
    }

    #[test]
    fn climbs_three_stairs() {
        assert_eq!(3, Solution::climb_stairs(3));
    }

    #[test]
    fn climbs_four_stairs() {
        assert_eq!(5, Solution::climb_stairs(4));
    }

    #[test]
    fn climbs_five_stairs() {
        assert_eq!(8, Solution::climb_stairs(5));
    }

    #[test]
    fn climbs_forty_five_stairs() {
        assert_eq!(1836311903, Solution::climb_stairs(45));
    }
}
