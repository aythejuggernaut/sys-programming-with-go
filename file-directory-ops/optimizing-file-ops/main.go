package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	filePath := "example.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()

	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("Failed to mmap file: %v\n", err)
		return
	}

	defer syscall.Munmap(data)

	fmt.Printf("Initial content: %s\n", string(data))

	newContent := []byte("Hello, mmap")
	copy(data, newContent)

	fmt.Printf("Content updated successfully!")

	if err := file.Sync(); err != nil {
		fmt.Printf("Failed to sync file: %v\n", err)
	}
}

/*
data, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.
PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED). There are two main areas
of this code to pay attention to:

• syscall.Mmap is used to map the file into memory. It takes the following arguments:
• int(file.Fd()): This extracts the file descriptor (an integer representing the opened file) from the file object. The file.Fd() method returns the file descriptor.
• 0: This represents the offset within the file where the mapping should begin. In this case, it starts at the beginning of the file (offset 0).
• int(fileSize): The length of the mapping, specified as an integer representing the size of the file (fileSize). This determines how much of the file will be mapped into memory.
• syscall.PROT_READ|syscall.PROT_WRITE: This sets the protection modes for the mapped memory. PROT_READ allows read access, and PROT_WRITE allows write access.
• syscall.MAP_SHARED: This specifies that the mapped memory is shared among multiple processes. Changes made to the memory will be reflected in the file, and vice versa.
• defer syscall.Munmap(data):
    • Assuming the memory mapping operation was successful (i.e., no error occurred), this defer statement schedules the syscall.Munmap function to be called when the surrounding function returns.
    • syscall.Munmap is used to unmap the memory region previously mapped with syscall.Mmap. It ensures that the mapped memory is released properly when it is no longer needed.
*/

/*
Out-of-memory safety
It's important to note that using a file-backed mapping is the appropriate choice for mmap, as opposed to an anonymous mapping. If you intend to make modifications to the mapped memory and have those changes written back to the file, then a shared mapping is necessary. With a file-backed, shared mapping, concerns about the Out-of-Memory (OOM) killer are alleviated, if your process operates in a 64-bit environment. Even in a non-64-bit environment, the issue would be related to addressing space limitations rather than RAM constraints, so the OOM killer would not be a concern; instead, your mmap operation would simply fail gracefully.
*/

/*
The mmap system call is a powerful tool in systems programming that allows you to map a file or device directly into the virtual address space of a process Here's a detailed breakdown of why it's used and how it works in the context of your Go code:

Why Use mmap?

1. Performance: I/O operations (reading from and writing to files) typically involve system calls, which are relatively expensive because they require context switching between the user space (where your application code runs) and the kernel space (where the operating system manages hardware and file systems). mmap avoids many of these system calls by allowing the operating system to handle file I/O through the virtual memory system. Once the file is mapped, accessing it becomes like accessing an array in memory, which is much faster.

2. Simplified Programming Model: With mmap, you can treat a file as if it were a byte array in memory. You can read, write, and modify its contents using simple pointer arithmetic or slice operations, just like you would with any other in-memory data structure. This eliminates the need for explicit read() and write() system calls and the management of read/write buffers.

3. Efficient Memory Usage: The operating system can load only the necessary portions of the file into physical memory (RAM) when they are accessed, using a technique called page faulting. This means you don't need to load the entire file into memory at once, which is particularly beneficial for large files.

4. Shared Memory: When using MAP_SHARED, multiple processes can map the same file into their address spaces. This allows them to share data directly through the mapped memory, providing a very efficient mechanism for inter-process communication.

How it Works

1. Memory Mapping: When you call mmap(), the operating system creates a mapping between a range of virtual addresses in your process's address space and a region in the file (or other memory-mapped object). This mapping is not an immediate copy of the file's data into memory; rather, it's a configuration that tells the OS how to translate virtual addresses to physical addresses when the process tries to access them.

2. Page Faults: When your program first tries to access a memory location within the mapped region, the CPU detects that this virtual address doesn't have a corresponding physical address mapping. This triggers a page fault, which is an interrupt that transfers control to the operating system kernel. The kernel then determines which part of the file corresponds to the accessed address, loads that part (typically a "page" of data, usually 4KB on modern systems) from the disk into a physical memory frame, and updates the process's page table to map the virtual address to this physical frame.

3. Accessing Data: After the page fault is handled, the process can access the data as if it were regular memory. The OS ensures that reads and writes to this memory are reflected in the underlying file.

4. Caching and Buffering: The operating system maintains a buffer cache in memory to store recently accessed file data. When you read from a mapped file, the OS might serve the data from the buffer cache if it's already there (which is faster than reading from disk). Similarly, when you write to the mapped memory, the data is initially written to the buffer cache and is eventually flushed to the disk by the OS (typically by a background process). The file.Sync() call in your code explicitly flushes any dirty (modified) data from the cache to the disk.

5. Unmapping: When you're done with the mapped memory, you should unmap it using syscall.Munmap(). This tells the OS that the process no longer needs the mapping, and the kernel can reclaim the resources associated with it. It also ensures that any pending writes in the buffer cache are flushed to disk.

In summary, mmap is used to bridge the gap between file I/O and memory access, enabling you to work with files in a more direct and efficient manner by leveraging the operating system's virtual memory management capabilities. It's particularly useful for applications that need to perform frequent or large file operations, or where performance is critical.
*/

/*
Memory-mapped I/O (mmap) is a technique that maps a file or a device directly into the virtual address space of a process. This allows the operating system to treat the file as if it were a large array in memory, enabling efficient random access and reducing the number of system calls required for I/O operations. Here's a breakdown of the key concepts:

Virtual Memory: Modern operating systems use virtual memory to provide each process with its own private address space. This virtual address space is independent of the physical memory (RAM) installed in the system. When a process accesses a virtual address, the hardware (specifically, the Memory Management Unit or MMU) translates it to a physical address in RAM using page tables maintained by the operating system.

Memory Mapping: Memory mapping involves creating a correspondence between a region of a file and a range of virtual addresses in a process's address space. This mapping is not an immediate copy of the file's contents into memory; rather, it's a configuration that tells the operating system how to handle accesses to those virtual addresses. The actual loading of file data into physical memory happens on demand through a mechanism called paging.

Paging and Page Faults: Physical memory is divided into fixed-size blocks called pages (typically 4KB on most systems). When a process tries to access a virtual address that is part of a memory-mapped file but is not currently in physical memory, a "page fault" occurs. This is an interrupt that transfers control to the operating system kernel. The kernel then:

Identifies the corresponding portion of the file that needs to be accessed.

Allocates a physical memory page (if one isn't already available).

Loads the required data from the file into that physical memory page.

Updates the process's page tables to map the virtual address to this physical page.

From the process's perspective, the page fault is handled transparently, and the data becomes accessible as if it were already in memory.

Page Fault Handling (on-demand loading): The operating system's page fault handler is responsible for managing the loading and unloading of pages from physical memory. When a page fault occurs:

The kernel determines which page of the file is being accessed and where it is located on disk.

It looks for a free page in physical memory. If none are available, it might evict an existing page (usually one that hasn't been used recently) to make room. This process of evicting pages is called swapping or paging out.

Once a page is available in physical memory, the kernel reads the corresponding data from the file into that page.

Finally, the kernel updates the process's page table to map the virtual address to the newly populated physical page, allowing the process to access the data. This on-demand loading ensures that only the necessary parts of the file are loaded into memory, which is particularly beneficial for large files.

Shared Memory: When a file is mapped with the MAP_SHARED flag (as in your code), the mapping is shared among multiple processes. This means that if one process modifies the data in the mapped memory, the changes are visible to other processes that have mapped the same file. This provides a very efficient mechanism for inter-process communication (IPC) because the processes can share data directly through memory without needing to copy it between their address spaces.

Memory Coherency and the Buffer Cache: The operating system maintains a buffer cache (also known as the page cache) in memory to store recently accessed file data. When multiple processes access the same file, the buffer cache helps to avoid redundant disk I/O operations. If a page from the file is already in the buffer cache (perhaps because another process recently accessed it), the page fault handler can satisfy the page fault by simply mapping that existing page into the new process's address space, without needing to read from the disk again.

Synchronization and Persistence: When data is modified in memory-mapped memory, it is initially written to the buffer cache. These "dirty" pages are eventually flushed to the underlying file by the operating system's page replacement algorithm or through explicit synchronization calls. The file.Sync() method in your code explicitly flushes any dirty pages associated with the file to disk, ensuring that the changes are made persistent and visible to other processes or systems that might access the file.
*/
