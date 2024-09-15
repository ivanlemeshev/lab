const std = @import("std");

pub fn squareOfSum(number: usize) usize {
    var sum: usize = 0;
    var n: usize = 1;
    while (n <= number) {
        sum += n;
        n += 1;
    }
    return sum * sum;
}

test "squareOfSum" {
    try std.testing.expectEqual(@as(usize, 0), squareOfSum(0));
    try std.testing.expectEqual(@as(usize, 1), squareOfSum(1));
    try std.testing.expectEqual(@as(usize, 9), squareOfSum(2));
    try std.testing.expectEqual(@as(usize, 3025), squareOfSum(10));
}

pub fn sumOfSquares(number: usize) usize {
    var sum: usize = 0;
    var n: usize = 1;
    while (n <= number) {
        sum += n * n;
        n += 1;
    }
    return sum;
}

test "sumOfSquares" {
    try std.testing.expectEqual(@as(usize, 0), sumOfSquares(0));
    try std.testing.expectEqual(@as(usize, 1), sumOfSquares(1));
    try std.testing.expectEqual(@as(usize, 5), sumOfSquares(2));
    try std.testing.expectEqual(@as(usize, 385), sumOfSquares(10));
}

pub fn differenceOfSquares(number: usize) usize {
    return squareOfSum(number) - sumOfSquares(number);
}

test "differenceOfSquares" {
    try std.testing.expectEqual(@as(usize, 0), differenceOfSquares(0));
    try std.testing.expectEqual(@as(usize, 0), differenceOfSquares(1));
    try std.testing.expectEqual(@as(usize, 4), differenceOfSquares(2));
    try std.testing.expectEqual(@as(usize, 2640), differenceOfSquares(10));
}
