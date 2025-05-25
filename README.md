# Local Network Test

Application created to test the connection between computers on the local network.  
One of the computers must act as a server, and the others should access IPLOCAL:8080 to verify that the connection is successfully established.

It includes a small built-in speed test to check if the network speed is satisfactory. The test is performed with the server hosting the site on the local network.

## Getting Started

### Running the Application

To run the application in test mode, use the following command:

```bash
go run main.go
```

### Creating the Test File

The application automatically generates the test file based on the `fileSize` variable defined in the `main` program. This file is placed in the `public` folder.  

For a more accurate local network test, the default value of the `fileSize` parameter is set to 2000 (for a 2GB file).

## Building the Application

### Build using Makefile
Use `make all` to compile the project for the main platforms.

To compile for a specific platform only, use `make {platform}`.  
Available options: `linux`, `windows`, `mac`, `mac-arm`.

To remove generated binaries, run `make clean`.

---

Made with ❤️ by Bi-Ga Tech
