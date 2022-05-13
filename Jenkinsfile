pipeline {
    // install golang 1.14 on Jenkins node
    tools {
            go 'go1.14'
    }

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "/usr/bin/go"
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
                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'Newtome19', usernameVariable: 'beatable2310')]) {
                sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                sh 'docker push beatable2310/new2c:latest'
                }
            }
        }
    }
}