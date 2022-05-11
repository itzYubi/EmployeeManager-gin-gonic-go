<h1 align="center">EmployeeManager made with gin-gonic-go</h1>

<h3>This a gin-gonic API based employee management application made with Golang. This is for educational purposes only. It provides a simplistic way of reading and writing to a CSV(Comma Seperated Values) file which acts as a database.</h3>


<h3 align="left">Languages and Tools:</h3>
<a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.w3.org/html/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/html5/html5-original-wordmark.svg" alt="html5" width="40" height="40"/> </a> <a href="https://postman.com" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a>


--------------------------------------------------------------------------------------------------------------------------------------------------------------------
<h3>Make sure you have Go installed in your system and install gin-gonic framework by running the command : </h3>
$ go get -u github.com/gin-gonic/gin

--------------------------------------------------------------------------------------------------------------------------------------------------------------------
<h3>Here is how to use it:</h3>
<h3> Run the command: </h3> go run src/main.go <h3> int he terminal and it should start running on Port 8080. If it shows some error make sure to kill the process on port 8080.</h3>

<h3>Download postman on your system and type the following url in it :</h3> http://localhost:8080/empdata/ 
<h3>Use the GET request to get the employee data from the CSV. (as in the image) </h3>

![GET](https://user-images.githubusercontent.com/63493407/167923694-f6a06554-9942-4d06-bdd6-6870df6e3a5a.png)

<h3>Use the POST request to add an employee data to the CSV. Add the new Employee in the body (as in the image) </h3>
![POST](https://user-images.githubusercontent.com/63493407/167921166-25df6879-2e9f-4c67-88b5-2d6f6f80dc8a.png)


--------------------------------------------------------------------------------------------------------------------------------------------------------------------

<h3 align="center"> Improvements </h3>
<h3> Feel free to improve this code. Try to make the name field of the imploy unique. Associate an employee id as a primary key and fetch employee details from it. Delete an employee from the csv. Add a database instead of a CSV! MongoDB is a good pick! </h3>


--------------------------------------------------------------------------------------------------------------------------------------------------------------------

Feel free to reach me incase of any doubts! Drop an email at utkarsh.b1358@gmail.com

--------------------------------------------------------------------------------------------------------------------------------------------------------------------
