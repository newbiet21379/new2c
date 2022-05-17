pipeline {
    // install golang 1.14 on Jenkins node
    agent { docker { image 'golang' } }
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
                sh 'docker build . -t beatable2310/new2c:latest'
            }
        }
        stage('deliver') {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', passwordVariable: 'pass', usernameVariable: 'user')]) {
                sh "docker login --username=${user} --password=${pass}"
                sh 'echo "Docker login successful"'
                sh 'docker push beatable2310/new2c:latest'
                }
            }
        }
    }
}