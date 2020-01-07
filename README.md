# dailyQuestions
Initial project was written in Python, then rewritten in Golang for educational purposes and as a smaller personal goal.

Lambda with Python was pretty straight forward.

Lambda using Golang caused some issues with deployment, writing it wasn't too difficult as it is quite a simple project, but ensuring that the Lambda is executable on AWS and deployed properly required the use of a Python script to give the files the appropriate permissions to be executed on Linux. Also required the use of the [AWS Lambda build command-line tool for windows](https://github.com/aws/aws-lambda-go/tree/master/cmd/build-lambda-zip) for building a deployment package from a Windows machine, this is apparently much less of a hassle when writing Go for Lambda of a Linux machine already.
