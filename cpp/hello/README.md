# Create Your First CMake Project

This guide demonstrates how to create, build, and run a simple C++ project using
CMake and Clang.

## Prerequisites

Make sure the following tools are installed:

- CMake
- Clang/Clang++
- A terminal

Verify the installation:

```sh
cmake --version
clang++ --version
```

---

# Step 1. Create the Project Directory

Create a new directory for the project:

```sh
mkdir hello
cd hello
```

Current project structure:

```text
hello/
```

---

# Step 2. Create `main.cpp`

Create a file named `main.cpp`:

```cpp
#include <iostream>

int main() {
    std::cout << "Hello, CMake!" << std::endl;
    return 0;
}
```

Project structure:

```text
hello/
└── main.cpp
```

---

# Step 3. Create `CMakeLists.txt`

Create a file named `CMakeLists.txt` in the project root.

```cmake
cmake_minimum_required(VERSION 4.0)

project(Hello)

add_executable(hello main.cpp)
```

### Explanation

- `cmake_minimum_required(VERSION 4.0)` specifies the minimum required CMake
  version.
- `project(Hello)` defines the project name.
- `add_executable(hello main.cpp)` tells CMake to build an executable named
  `hello` from `main.cpp`.

Project structure:

```text
hello/
├── CMakeLists.txt
└── main.cpp
```

---

# Step 4. Configure the Project

Run the following command to configure the project and generate build files:

```sh
cmake -S . -B build
```

CMake will generate the appropriate build system for your platform.

---

# Step 5. Build the Project

Compile the project:

```sh
cmake --build build
```

---

# Step 6. Run the Program

On Linux or macOS:

```sh
build/hello
```

Expected output:

```text
Hello, CMake!
```

---

# Final Project Structure

After building, your project should look similar to this:

```text
hello/
├── build/
│   ├── CMakeCache.txt
│   ├── CMakeFiles/
│   ├── Makefile (or build.ninja)
│   └── hello
├── CMakeLists.txt
└── main.cpp
```
