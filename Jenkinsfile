pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go '1.13.8'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
//                 sh 'make unit-tests'
            }
        }
        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
//                 sh 'make functional-tests'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'go build -o /new2c'
            }
        }
        stage('deliver') {
            agent any
            steps {
                sh '/new2c'
            }
        }
    }
}