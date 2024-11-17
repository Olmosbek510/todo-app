# To-Do App Backend

## Overview
This backend for a To-Do application is developed in Go. It facilitates basic CRUD (Create, Read, Update, Delete) operations for managing to-do items and lists. This application supports user authentication through sign-in and sign-up functionalities, utilizing JWT tokens for secure access.

## Key Features
- **User Authentication:** Secure sign-in and sign-up functionalities.
- **CRUD Operations:** Manage to-do items and lists.
- **Secure Access:** Utilizes JWT tokens for authentication and secure access.

## Technologies Used
- **Go:** The application is written in Go, known for its efficiency and scalability in backend development.
- **PostgreSQL:** A robust and scalable relational database.
- **Docker:** For containerization and easy deployment.
- **JWT:** For securing the API endpoints.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
Ensure you have the following installed:
- Docker
- Docker Compose
- Go (latest version recommended)

### Installation

1. **Clone the Repository**
    ```
    git clone [https://github.com/Olmosbek510/todo-app.git]
    cd [todo-app]
    ```

2. **Set Up Environment Variables**
   Create a `.env` file in the project directory and add the following environment variables:
    ```
    PORT=8XXX
    DB_PASSWORD=some_password
    ```

3. **Install Dependencies**
   Load the necessary Go packages:
    ```
    go mod tidy
    ```

4. **Docker Compose**
   Use Docker Compose to set up and start the PostgreSQL database:
    ```
    docker-compose up -d
    ```

### Running the Application
To start the server, run the following command in the project's root directory:

## Usage
Once the server is running, you can perform CRUD operations on to-do items and lists via the exposed API endpoints.

## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Your Name – [@Olmos_Urazboev](https://telegram.org/Olmos_Urazboev) - @Olmos_Urazboev
Project Link: [https://github.com/Olmosbek510/todo-app](https://github.com/Olmosbek510/todo_app)