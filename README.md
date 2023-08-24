# A Golang API for Cybersecurity Job Opportunities 🛡️💼

This project is a modern Job Opportunity API for CyberSecurity (or others) built using Golang. The API is developed by Go-Gin as the router, GoORM for database communication, SQLite (for now) as the database, and Swagger for documentation and API testing. The project follows a modern package structure to keep the code base organized and maintainable.

---

## Features 🌟

- Introduction to Golang and building modern APIs
- Setting up the development environment for creating the API
- Using Go-Gin as a router for route management
- Implementing SQLite (for now) as database API
- Use of GoORM to communicate with the database
- Swagger integration for API documentation and testing
- Creation of a modern package structure for project organization
- Implementation of a complete job opportunities API with endpoints for searching, creating, editing and deleting opportunities (CRUD)
  
## Installation 🚚

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/username/repo-name.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands 🛠️

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `secjobs`.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `secjobs` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Docker and Docker Compose 🐳

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

For more information on Docker and Docker Compose, refer to the official documentation:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Used Tools 📝🛠️

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage 📝

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License 📜

This project is licensed under the GNU GENERAL PUBLIC LICENSE - see the LICENSE.md file for details.

## Credits 👏

This project was created by [fowkeio](https://github.com/fowkeio).
