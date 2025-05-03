Application created to test the connection between computers on the local network.  
One of the computers must act as a server, and the others should access IPLOCAL:3000 to verify that the connection is successfully established.

It includes a small built-in speed test to check if the network speed is satisfactory. The test is performed with the server hosting the site on the local network.

## Getting Started

### Running the Application

To run the application in test mode, use the following command:

```bash
go run main.go
```

### Creating the Test File

The application automatically generates the test file based on the `fileSize` variable defined in the `main` program. This file is placed in the `public` folder.  

For a more accurate local network test, it is recommended to set the `count` parameter to 2000.

To manually create a 100 MB test file, you can use the following command:

```bash
dd if=/dev/zero of=public/test-file.bin bs=1M count=100
```

**Note for Windows users:** The `dd` command is not available by default and must be installed separately. On most other platforms, it is natively available.
