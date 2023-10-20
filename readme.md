![image](https://github.com/Mateusnasciment/Go.hexagonal/assets/106212780/48feba04-2e8f-4668-bfeb-b41fef6ac37d)


To run your application with the provided Docker Compose file, you can follow these steps:

Install Docker and Docker Compose: Make sure you have Docker and Docker Compose installed on your system. If not, you can download and install them from the official websites: Docker and Docker Compose.

Create a directory for your project: Create a directory for your project, and place your Go application source code and the Docker Compose YAML file in that directory.

Update Docker Compose file: Open the Docker Compose YAML file (e.g., docker-compose.yml) and modify it according to your project's specific requirements, such as environment variables and service names.

Navigate to your project directory: Open your terminal or command prompt and navigate to the directory where your project files are located.

Build your application: Run the following command to build your Go application and create a Docker image for it:


docker-compose build
Start your application: Once the build is complete, start your application and all associated services with the following command:


docker-compose up
This will start the services defined in the Docker Compose file, including your Go application, PostgreSQL, Kafka clusters, and Confluent Control Center.

Access your application: Your Go application should be accessible at http://localhost:8080, and Confluent Control Center for Kafka monitoring should be accessible at http://localhost:9021.

Monitor your application: You can use Confluent Control Center to monitor your Kafka clusters and topics by accessing it in your web browser.

Manage the running services: To stop all the services and containers, you can use Ctrl+C in the terminal where you started docker-compose. To stop and remove all containers and networks, run the following command in your project directory:


docker-compose down
This should allow you to run your application with the provided Docker Compose setup in an English environment. Make sure to adapt the environment variables and configurations according to your specific project requirements.
