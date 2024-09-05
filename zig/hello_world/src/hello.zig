const std = @import("std");

const testing = std.testing;

pub fn hello() []const u8 {
    return "Hello, World!";
}

test "hello" {
    try testing.expectEqual("Hello, World!", hello());
}
