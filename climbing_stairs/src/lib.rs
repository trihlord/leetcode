/// # 70. Climbing Stairs
///
/// You are climbing a staircase. It takes n steps to reach the top.
///
/// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
///
/// ## Example 1
///
/// **Input:** n = 2
/// 
/// **Output:** 2
/// 
/// **Explanation:** There are two ways to climb to the top.
/// 
/// 1. 1 step + 1 step
/// 2. 2 steps
///
/// ## Example 2
///
/// **Input:** n = 3
/// 
/// **Output:** 3
/// 
/// **Explanation:** There are three ways to climb to the top.
///
/// 1. 1 step + 1 step + 1 step
/// 2. 1 step + 2 steps
/// 3. 2 steps + 1 step
///
/// ## Constraints:
///
/// - `1 <= n <= 45`
pub fn climb_stairs(n: i32) -> i32 {
    n
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_1_works() {
        assert_eq!(climb_stairs(2), 2);
    }

    #[test]
    fn example_2_works() {
        assert_eq!(climb_stairs(3), 3);
    }

    /// ## Case 3
    ///
    /// **Input:** n = 4
    /// 
    /// **Output:** 5
    /// 
    /// **Explanation:** There are five ways to climb to the top.
    /// 
    /// 1. 1 step + 1 step + 1 step + 1 step
    /// 2. 1 step + 1 step + 2 steps
    /// 3. 1 step + 2 steps + 1 step
    /// 4. 2 steps + 1 step + 1 step
    /// 5. 2 steps + 2 steps
    #[test]
    fn case_3_works() {
        assert_eq!(climb_stairs(4), 5);
    }
}
