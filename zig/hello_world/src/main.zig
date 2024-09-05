const std = @import("std");
const greeting = @import("hello.zig");

pub fn main() void {
    std.debug.print(greeting.hello(), .{});
}
